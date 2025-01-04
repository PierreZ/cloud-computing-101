# Deploy my-app with Kubernetes

```shell
kubectl apply -f deploy.yml

# Access the service
kubectl describe service my-api-service | grep NodePort
kubectl describe nodes | grep InternalIP

kubectl get pods -w
```