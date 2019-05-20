package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"

	"github.com/olekukonko/tablewriter"
	"net/http"
	"os"
	"sort"
	"sync"
)

var (
	u          string
	concurrent int
	requests   int
	resultLock sync.RWMutex
	client     *http.Client
	m          map[string]int
)

var rootCmd = &cobra.Command{
	Use:   "respy",
	Short: "respy gets responses",
	Long:  `respy shows the percentage distrubtion for HTTP response text.`,
	Run: func(cmd *cobra.Command, args []string) {
		runRespy()
	},
}

func init() {
	rootCmd.PersistentFlags().StringVar(&u, "u", "http://numbersapi.com/42", "a valid url, HTTP(S)")
	viper.BindPFlag("u", rootCmd.PersistentFlags().Lookup("u"))

	rootCmd.PersistentFlags().IntVar(&requests, "n", 1000, "number of total requests")
	viper.BindPFlag("n", rootCmd.PersistentFlags().Lookup("n"))

	rootCmd.PersistentFlags().IntVar(&concurrent, "c", 100, "# concurrent requests")
	viper.BindPFlag("c", rootCmd.PersistentFlags().Lookup("c"))

	client = &http.Client{}
	m = map[string]int{}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func quit(err error) {
	fmt.Printf("⚠️ %v\n", err)
	os.Exit(1)
}

// 1) Make "n" total requests, "c" concurrently (save responses in a map)
// 2) print in a table
func runRespy() {
	fmt.Printf("☎️   %d requests to %s...\n", requests, u)

	// 1
	err := makeRequests()
	if err != nil {
		quit(err)
	}
	// 2
	printResults()
}

// makeRequests -> runWorker -> oneRequest
func makeRequests() error {
	var wg sync.WaitGroup
	wg.Add(concurrent)

	// run <concurrent> workers. each makes n/c requests
	for i := 0; i < concurrent; i++ {
		go func() {
			runWorker(requests / concurrent)
			wg.Done()
		}()
	}
	wg.Wait()
	return nil
}

func runWorker(n int) {
	for j := 0; j < n; j++ {
		res, err := oneRequest()
		if err != nil {
			fmt.Printf("⚠️ %s", err)
			continue
		}
		// log the result
		resultLock.Lock()
		if _, ok := m[res]; !ok {
			m[res] = 1
		} else {
			m[res] = m[res] + 1
		}
		resultLock.Unlock()
	}
}

// helper for makeRequests
func oneRequest() (string, error) {
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("foo", "bar1")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	// resp, err := http.Get(u)
	// if err != nil {
	// 	quit(err)
	// }

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		quit(err)
	}
	return string(body), nil
}

// processes map
// TODO - have an option to print percents not raw totals
func printResults() {
	proc := map[string]float64{}
	for k, v := range m {
		percent := 0.0
		percent = float64(v) / float64(requests) * 100
		proc[k] = percent
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Response", fmt.Sprintf("%% of %d Requests", requests)})

	// alphabetize keys
	alphaKeys := []string{}
	for k := range proc {
		alphaKeys = append(alphaKeys, k)
	}
	sort.Strings(alphaKeys)
	for _, k := range alphaKeys {
		v := proc[k]
		valStr := fmt.Sprintf("%.1f%%", v)
		table.Append([]string{k, valStr})
	}
	table.Render() // Send output
}
