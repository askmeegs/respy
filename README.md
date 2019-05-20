# respy

A command-line tool to show a percentage distribution  

(Created to demonstrate [Istio traffic splitting](https://istio.io/docs/concepts/traffic-management/#splitting-traffic-between-versions) across versions.)

## how-to 

```
kubectl apply -f manifests/deployment.yaml 

RESPY_POD=$(kubectl get pod  | grep respy | awk '{ print $1 }') 

 kubectl exec -it $RESPY_POD  -c respy ./respy 
```

**should see:** 


```
$ kubectl exec -it $RESPY_POD  -c respy ./respy 

☎️   1000 requests to http://numbersapi.com/42...
+--------------------------------+--------------------+
|            RESPONSE            | % OF 1000 REQUESTS |
+--------------------------------+--------------------+
| 42 is the angle in degrees for | 10.1%              |
| which a rainbow appears or the |                    |
| critical angle.                |                    |
| 42 is the answer to the        | 9.7%               |
| Ultimate Question of Life, the |                    |
| Universe, and Everything.      |                    |
| 42 is the number of US gallons | 9.9%               |
| in a barrel of oil.            |                    |
| 42 is the number of gallons    | 9.8%               |
| that one barrel of petroleum   |                    |
| holds.                         |                    |
| 42 is the number of kilometers | 10.6%              |
| in a marathon.                 |                    |
| 42 is the number of laws of    | 10.7%              |
| cricket.                       |                    |
| 42 is the number of museums    | 10.0%              |
| in Amsterdam (Netherlands has  |                    |
| the highest concentration of   |                    |
| museums in the world).         |                    |
| 42 is the number of spots      | 10.6%              |
| (or pips, circular patches or  |                    |
| pits) on a pair of standard    |                    |
| six-sided dice.                |                    |
| 42 is the result given by the  | 9.5%               |
| web search engines Google,     |                    |
| Wolfram Alpha and Bing when    |                    |
| the query "the answer to life  |                    |
| the universe and everything"   |                    |
| is entered as a search.        |                    |
| 42 is the sum of the codes of  | 9.1%               |
| the letters in the words "BIG  |                    |
| BANG" using the encoding A=1,  |                    |
| B=2, C=3, etc.                 |                    |
+--------------------------------+--------------------+
```


## customize 

```
$ respy --help

respy shows the percentage distrubtion for HTTP response text.

Usage:
  respy [flags]

Flags:
      --c int      # concurrent requests (default 100)
  -h, --help       help for respy
      --n int      number of total requests (default 1000)
      --u string   a valid url, HTTP(S) (default "http://numbersapi.com/42")
`` 


