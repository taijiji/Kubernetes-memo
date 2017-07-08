# Keywords of kubernetes

# Node
- host server
- 1 node = 1 Virtual Machine

# Cluster
- aggregation of nodes

# Pod
- aggregation of containers = 1 application
    - reverse proxy 
    - web application
- minimum unit to deploy by k8s
- pod runs on node
- containers of pod are co-sheduled
- containers of pod are co-located = deploy on 1 node

# Container
- 1 container = 1 process of linux machine

# Selector
- using "Label" to control target group
    - label example : role = web

# Service
- 1 service = multiple pod = multiple application
- "Service Discovery" using DNS record. 
    - ex: web.default.svc.cluster.local
- service are defined name(=DNS record) and selectoer(=label)
    - name: web
    - role: web
- types of Service
    - ClusterIP     : DNS name or Virtual IP address on cluster
    - NodePort      : publish static port number using ClusterIP
    - LoadBalancer  : make load balancer and link to service using NodePort
    - ExternalName  : To use exernal services, make alias for internal cluster using DNS CNAME.
    - Headless      : To use "StatefulSet" function, make DNS SRV record of pod to DNS name resolution.

# Ingress
- HTTP(Layer 7) load balancer for access to Service
- function
    - forward by URL pattern
    - SSL termination

# Volume
- external strage area
    - variable data shuouldn't be used as contianer 
        - container should be immutable area.
    - Can mount to contaier
- for cloud provider's storage
    - for GCE, for Azure, for AWS
- for onpremise hardware environment
    - iscsi
    - nfs
    - gitRepo   : Git repository
    - cinder    : OpenStack block storage

## emptyDir Vulume
- volume on node
- its be remove when pod be remove
    - cannot perpetuate data

## hotPath Volume
- volume on node. asign path to node directory.
- its be not remove when pod be remove
- its be remove when Node be remove
    - cannot perpetuate data
    - Do not use on production environment

## PersistentVolume (PV)
- volume to perpetuate data safety
    - use iscsi/nfs
    - use cloud provider's storage

## PersistentVolumeClaim (PVC)
    - asign volume size to use PV
    - 1 VPC on PV
    - Can perpetuate data to "complete PVC" and "mount to container".

## StorageClass
- dynamic storage volume provisioning on clound provider's environment
- automatically mapping to PVC for auto-scale Pod

# ConfigMap / Secret
- setting file to configure container/pod
- key and value format
- useful to change the parameter on setting files
    - ex: nginx.conf. 
    - No need to re-create container image.
- ConfigMap
    - plain text
- Secret
    - base 64 encoded. Not plain text.But unsecure.
- use cases
    1. mount as Volume
        - ex: nginx.conf
    1. register as engironment variables on Pod
        - convert variables to use on application program.

# Role-Based Access Contorl (RBAC)
- Access control management
- after k8s 1.6 (beta)

# ReplicaSet
- controller to support Horizontal Auto-scale of Pod
    - ex: replication the application Pod
- template
    - Pod definition file.
- replicas
    - define Number of Pod.
    - ReplicaSet keeps the number of Pod by increasing/decreacing Pods.
- New Pod are asigned new ID (random charactor)
- Cannot expect which Node be used when new Pod create.
    - But, it doesn't use less resource Nodes. It's assured.

# Deployment
- generation management of ReplicaSet template 
- When new pod created, careate new and old one.then it replaced them
- Can rolling update . Non-stopped .
- Can retrieve old status by 1 command

# Horizontal Pod Autoscaling (hpa
- autoscaling contorol by CPU load average or other metric
- Control number of pod by changing deployment or ReplicaSet

# DaemonSet
- set 1pod on every each Nodes
- Ex:fluent, Prometheus

# StatefulSet
- it set specific ID to Pod
- Ex: mysql-0, mysql-1, mysql-2
- Static link name to Pod.
- start/stop/restart in order to ID

# Job
- Assurance system to correctly start and stop
- Parallelism
    - Number of Pod running in parallel
- Completions
    - Job completion condition number to complete Pod

# Cronjob
- Job with schedule field to use crontab to run job priodically.


 
