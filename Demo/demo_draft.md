In this video, I will demonstrate how to achieve the near-zero downtime database version upgrade/downgrade.

1. Introduce the environment

   * `kl get all`
   * `kl describe nginx`

   The following environment will be utilized throughout this demonstration:

   - `postgres-master` represents the master node, which currently has an older version of the database. Prior to this demonstration, the master node has been configured to support logical replication.
   - `nginx` serves as a load balancer, enabling users outside the cluster to read and write data to and from the database. Initially, it points to the master node. As our Kubernetes operator functions, it will automatically switch to the follower node.
   - `grafana` provides a visual representation of when synchronisation can be terminated.

2. Start the client, continuously send and read from the database

   This Client programme is used to imitate the the client outside the cluster sending requests to the database through the nginx.

3. Show the master node and database information

   Let’s execute into the master node.

4. `make install` and `make deploy`

5. Check the follower node and data in it

6. Show the Grafana result

----

Background:

In production environments, many PostgreSQL databases typically operate within a cluster. Periodic upgrades are essential to maintain optimal performance and security. The conventional upgrade strategy involves the following steps: suspending the service, pausing all user requests, backing up and migrating data to the updated database version, verifying data consistency between the old and new versions, and finally resuming the service and processing requests. However, this approach necessitates a certain amount of downtime, which falls short of the desired five nines (99.999%) high availability.

In this study, we propose and implement a Kubernetes (k8s) operator that automates database version management while minimizing downtime. This operator streamlines the upgrade process and brings us closer to achieving 99.999% high availability.