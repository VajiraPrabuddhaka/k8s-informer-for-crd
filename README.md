# k8s client-go informer for CRD

Use the following steps to tryout this.

1. Create crd and cr
```shell
kubectl apply -f resources/api_crd.yaml
kubectl apply -f resources/api_cr.yaml
```

2. Run the informer
```shell
go run main.go
```

3. Try Create, Update, delete of Custom Resources and you will see onAdd, onUpdate, onDelete methods will call accordingly.

