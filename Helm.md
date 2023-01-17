# Helm

A simple way to explain how Helm works: Helm installs charts into the Kubernetes cluster, creating a new release each time it is installed. You can find new charts in Helm's chart repositories.

```shell
# check which chats are published
$ helm list

# delete a specific chart
# NOTE: This command will delete all the relevent resources including (service, deployment, pod, etc.) and even the history version.
$ helm uninstall [chart_name]
	--keep-history will keep the history version of this chart.

# check the history of chart
$ helm status [chart_name]
```

## How to create a helm charts

1. Use `helm create` to initialise a chart. The structure of the new chart is like the following.

   ```shell
   $ helm create [chart_name]
   
   .
   ├── Chart.yaml
   ├── charts
   ├── templates
   │   ├── NOTES.txt
   │   ├── _helpers.tpl
   │   ├── deployment.yaml
   │   ├── hpa.yaml
   │   ├── ingress.yaml
   │   ├── service.yaml
   │   ├── serviceaccount.yaml
   │   └── tests
   │       └── test-connection.yaml
   └── values.yaml
   ```

2. Use `helm upgrade ...` to install all resources.

   This command also can be used to upgrade the chart when we implement some changes.

   ```shell
   $ helm upgrade --install [helm_chart_name] --values values.yaml .
   ```

3. Use `helm rollback ...` to rollback to a specific version.

   ```shell
   $ helm rollback [helm_chart_name] [version_num]
   ```

4. Use following commands to package/install a chart.

   ```shell
   $ helm package hello-helm
   
   $ helm repo index .
   
   $ helm upgrade --install hello-helm hello-helm-0.1.0.tgz
   ```

   