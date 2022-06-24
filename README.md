# in-memory-http-service
This is an in-memory key-value store HTTP API service.

1. Run go file using: go build cmd/main.go && ./main (if package are not sync run go mod download , go mod tidy)
2. call the endpoint using curl localhost:8080/api/v1/


# package structure
```
├── Dockerfile
├── LICENSE
├── README.md
├── app.env
├── cmd
│   └── main.go
├── go.mod
├── go.sum
├── in-memory-http-service
│   ├── Chart.yaml
│   ├── charts
│   ├── templates
│   │   ├── NOTES.txt
│   │   ├── _helpers.tpl
│   │   ├── deployment.yaml
│   │   ├── hpa.yaml
│   │   ├── ingress.yaml
│   │   ├── service.yaml
│   │   ├── serviceaccount.yaml
│   │   └── tests
│   │       └── test-connection.yaml
│   └── values.yaml
├── internal
│   ├── handlers
│   │   ├── handlers.go
│   │   └── handlers_test.go
│   └── startup.go
└── k8s-deployment.yml
```

Install or Upgrade helm using 
   ```shell
    ╰─ helm install in-memory-http-service in-memory-http-service
              OR
    ╰─ helm upgrade --install in-memory-http-service in-memory-http-service 
   
   Expected Output:

   NAME: in-memory-http-service
   LAST DEPLOYED: Sun Jun 12 15:45:53 2022
   NAMESPACE: default
   STATUS: deployed
   REVISION: 1
   NOTES:
   1. Get the application URL by running these commands:
      export NODE_PORT=$(kubectl get --namespace default -o jsonpath="{.spec.ports[0].nodePort}" services in-memory-http-service)
   
      export NODE_IP=$(kubectl get nodes --namespace default -o jsonpath="{.items[0].status.addresses[0].address}")
   
      echo http://$NODE_IP:$NODE_PORT : http://192.168.64.2:30088
   
   
   And last
   for deleting helm, you can use:
   ╰─ helm uninstall in-memory-http-service in-memory-http-service 
   ```

5. Now we can hit the service 
   ```shell
      curl 'http://192.168.64.2:30088/api/v1/'
   ```
   
   
# steps to run the docker file 
1. docker build -t in-memory-http-service ..
2. docker run -d -p 8080:8080 --name in-memory-http-service in-memory-http-service:latest


# k8s resource created
```shell
NAME                     TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)          AGE
service/in-memory-http-service   NodePort    10.103.8.244   <none>        8000:30088/TCP   35s
service/kubernetes       ClusterIP   10.96.0.1      <none>        443/TCP          41h

NAME                                  READY   STATUS    RESTARTS   AGE
pod/config-service-6bbfff8bb9-559gn   1/1     Running   0          35s
pod/config-service-6bbfff8bb9-cm9fd   1/1     Running   0          35s

NAME                             READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/in-memory-http-service    3/3     3            3           35s
```



## Thank You
Ankit Kumar
