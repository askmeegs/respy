# respy

show the percentage of responses

## how-to 

```
kubectl apply -f manifests/deployment.yaml 

RESPY_POD=$(kubectl get pod  | grep respy | awk '{ print $1 }') 

kubectl exec -it $RESPY_POD  -c respy ./respy  -- --u http://weather-backend:5000/version
```

## should see 


```
$ kubectl exec -it $RESPY_POD  -c respy ./respy  -- --u http://weather-backend:5000/version
☎️   1000 requests to http://weather-backend:5000/version...
+--------------------------------+--------------------+
|            RESPONSE            | % OF 1000 REQUESTS |
+--------------------------------+--------------------+
| weather-backend: single        | 46%                |
| upstream connect error or      | 4%                 |
| disconnect/reset before        |                    |
| headers                        |                    |
| weather-backend: multiple      | 50%                |
+--------------------------------+--------------------+
```


## customize 

```
respy shows the percentage distrubtion for HTTP response text. good 4 istio

Usage:
  respy [flags]

Flags:
      --c int      # concurrent requests (default 100)
  -h, --help       help for respy
      --n int      number of total requests (default 1000)
      --u string   a valid url, HTTP(S) (default "http://numbersapi.com/42")
`` 


