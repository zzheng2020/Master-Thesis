# K8s Operator

Custom Resource Definition (CRD) is used to define own resource type.

## Controller

Each resource type has its own controller.

When controller is running, it tries to get the current state of the resource, and compare it with the desired state. If they don’t match, controller will adjust them and change current state to desired state as soon as possible. This process is called ***reconciliation***. 

Reconcile aims to move the current state of the cluster closer to the desired state.

<img src="pic4md/operator-controller.png" style="zoom:60%;" />
