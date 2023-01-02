# K8s

* Big picture of K8s.

<img src="/Users/zihengzhang/Code/Master-Thesis/pic4md/big_pic_k8s.png" style="zoom:30%;" />

* Master node in K8s.

  <img src="/Users/zihengzhang/Code/Master-Thesis/pic4md/master_in_k8s.png" style="zoom:40%;" />

  * etcd: distributed kv strorage, use Raft.

  * API server: Interact with cluster.

  * Scheduler: scheduling and decision making.

  * Controller Manager: guarantee the eventual consistency of status.

* Worker in k8s.

  <img src="/Users/zihengzhang/Code/Master-Thesis/pic4md/worker_in_k8s.png" style="zoom:40%;" />

  * kubelet: resource controller of worker.
  * container runtime: controller of container.
  * kube-proxy: manage the service.

* Process of k8s.

  <img src="/Users/zihengzhang/Code/Master-Thesis/pic4md/process_k8s.png" style="zoom:40%;" />

* Framework of k8s.

  <img src="/Users/zihengzhang/Code/Master-Thesis/pic4md/framework_k8s.png" style="zoom:40%;" />

* Pod: 

  * Example

    ```yaml
    apiVersion: v1
    kind: Pod
    metadata:
      name: ett
    spec:
      containers:
      - name: ett
        image: zzheng2020/ett-web-framework
    ```

    

  