# LF258 Cource
https://training.linuxfoundation.org/linux-courses/system-administration-training/kubernetes-fundamentals

# mypage
https://training.linuxfoundation.org/cas?destination=portal

# Lab
https://lms.quickstart.com/custom/858487/LFS258%20-%20Labs.pdf

# Chapter1
Intoroduction Linux foundation

# Chapter2
- Basics of Kubernetes
- what is Kubenetes
    - Connecting containers across multiple-hosts
    - scaling
    - deploying app without downtime
    - Service discovery
    - a set of Primitives and a Powerful API
    - 15 years of experience of Google
- the meaning of Kubernetes
    - Greek: Helmsman or Pilot of the ship
    - pilot of a ship of containers
- Containers and Docker have emppered developers with:
    - Ease of building container images
    - Simplicity of shaing images via Docker registries
    - Powerful user experience to manage contaiers
- Challenge
    - managing containers at scale
    - architecting a distributed application based on Microsercvices' principal
- Continuous Integration pipeline
    1. Build your contaier images
    1. test them
    1. verify them
    1. need a cluster of machines acting as your base infrastructure on which to run your containers
    1. need a system to lunch your containers
    1. watch over them when things fail and self-heal
    1. be able to perform rolling updates, rollbacks
    1. need a network setup which permits self-discovery of services in a very ephemeral environment
- Other Solutions
    - Docker Swarm
        - Docker Solution. 
        - It's embedded with the Docker Engine.
    - Apache Mesos
        - DataCenter scheduler. 
        - it builds Multi-level shedurler
        - it can run containers throuch the use of Marathon frameworks
        - it aims to better use a datacenter cluster.
    - Nomad
        - HashiCorp product. the makers of Vagrant and Consul.
        - Managing Containerized application
        - Nomad schedule task defined in Jobs
        - It has Docker driver
    - Rancher
        - container orchestrator agnostic system
        - it provides a sigle pane of glass interface, to managing application.
        - it supports Mesos, Swarm, Kubernetes, Cattle(native system)
    - Kubenetes's distinguishing point is Heritage(History)
- The Borg Heritage
    - Kubernetes is inspired by Borg
        - internal system used by Google
    - Borg is operating over 15 years, world wide.
    - Borg was a Google Secret for a long time.
    - Borg was finally described publicly in 2015. 
    - Borg Documentes: https://research.google.com/pubs/pub43438.html
- Borg Lineage (they inspired by Borg)
    - Borg
        - Mesos
        - Kubernetes
        - Omega
        - Cloud Foundry
        - cgroups
            - LXC
            - Docker
            - OCI
                - rkt/appc
    - cgroups
        - Linux kernel,2007,Contributed by Google. It limits the resource used by collection of processes
        - cgroups and Linux namespaces are at the heart of containres today, including Docker.
    - Mesos
        nspired by discussion with Google when Borg was still a secret.
    - Cloud Foundry Foundation
        - 12 factor application principles. great guidance to build web application
            - it can scale easily
            - it can be deployed in the cloud
            - its build is automated
            - https://12factor.net/
        - Borg and Kubernetes address these principles as well
- Kubernetes Architechture
    - API Server : you can communicate with API
    - kubectl : local client
        - ajax
        - etc
    - Scheduler : requests for running container coming to API an dfind a suitable node to run
    - Controller Manager
    - Nodes
        - Kubelet : it receive requests to run the container and watch over them
        - Service Proxy : it creates and manages neteworking rules
- Kubernetes componentes
    - master : Central Manager
        - it run API server, a Sheduler, a controller, a storage system
        - to keep the sate of the cluster
    - worker nodes
    - Each nodes in the cluster runs tow process, "kubelet" and "Services Proxy"
- Kubernetes characteristics
    - it is made of a manager and a set of nodes
    - it has a scheduler to place container in a cluster
    - it has an API server and a persistence layer with etcd
    - it has a controller to reconcile states
    - it is deployed on VMs or bare-metal machines, in public Clouds, or on-premise
    - it is written in GO
    - it licensed under the Apache Software License v2.0
- Kubernetes Community
    - it was open sourced in June 2014
    - 1500+ contributors
    - K8s Core has over 55,000 commits
    - Meet up:  over 150 cities, 100,000 particioantes in 50 countories
    - slack: 20,000 people
    - Major release: every 3 months
    - Weekly Hangouts take places
- Cloud Native Computing Foundation
    - Google donate Kubernetes to a newly formed foundation within the Linux Foundation in July 2015, when K8s v1.0
    - CNCF members; cisco, Cloud Foundry Foundation, AT&T, Goldman Sachs, etc
- K8s Documents
    - Official Page: https://kubernetes.io/
    - Github: https://github.com/kubernetes/kubernetes
    - Paper: Google Borg(2015): https://research.google.com/pubs/pub43438.html
    - Paper: Borg, Omega, and Kubernetes(2016): https://research.google.com/pubs/pub44843.html
    - Podcast: Borg and Kubernetes with John Wilkes(2016): https://www.gcppodcast.com/post/episode-46-borg-and-k8s-with-john-wilkes/
    - YouTube: Cluster Management at Google(2014):  https://www.youtube.com/watch?v=VQAAkO5B5Hg 
    - CNCF page: https://www.cncf.io/

# Chapter3
- Kubernetes Cluster
    - Master node
    - a set of Worker nodes
- For testing process: Minikube
    - these nodes can be run on the same node(Virtual or Physical).
- Kubernets 6 compornents, it can run as starnderd Linux process or docker container.
    - API server : on master node
    - Scheduler : on master node
    - Controller manager : on master node
    - kubelet: on worker node
    - proxy (aka kube proxy) : on worker node
    - etcd cluster : on master node
- K8s Master nodes runs;
    - API server
        - REST interface to all the K8s resources
        - it's highly configurable
    - scheduler
        - place to the containers on the node in the cluster.  according to various policies, metrics, resource requiremetns
        - it's configurable via command line flags
    - controller manager
        - reconciling the actual state of cluster with the desired state, specified via the API.
            - control loop that performs action based on the observed state of the cluster and the desired state.
        - it's highly configurable
    - the master node can be configured in a multi-master highly available
        - setup: https://github.com/kelseyhightower/kubernetes-the-hard-way
        - indeed, the shcedulers and contoller manages can elect a leader, while the API servers can be fronted by a load-balancer.
- K8s worker nodes runs;
    - kubelet
        - interact with the underlyining docker engine
        - installed on all worker nodes
        - make sure that the containers that need to run are actually running.
    - kube-proxy
        - managing the network connectivity to the containers
    - as well as docker engine
- Keeping State with etcd
    - K8s needs a persistency later,  to know the state fo the cluster over time.
    - Traditionally, relational database may become a single point of failure
    - K8s uses etcd.
        - Distributed key-value stores, made to run on multiple nodes
        - It can be run on a single node (it loses its distributed characteristic, and hence, it's tolerance to failure)
        - etcd use a leader election algorithm
        - to provide strong consistency of the stored state among the nodes.
        - etcdctl
- Networking
    - a pod is a group of co-localted containers the share the same IP address.
    - the network needs to assing IP address to pods, and needs to provide traffic routes between all pods on any nodes
    - Networking challenges
        - Coupled container-to-container communications
            - solved by the pod concept
        - pod-to-pod communication
        - external-to-pod communications
            - solved by the services concept
- Pod
    - K8s call the "single IP-per-Pod Model" 
    - Pod represents a group fo co-locationed contaienrs with associated data volumes.
    - The lowest compute unit in K8s is the Pod.
    - all containers in a pod communicate over "localhost"
    - Pod has a single IP address
    - Containers in a pod, Container A and B and "pause conteiner" share the network namespace. 
        - "pause container" is used to get an IPaddress, then all the containers will use its network namespace
_ CNI
    - K8s is standardizing on "the Container Network Interface(CNI)" spectification.
    - kubeadam(K8s cluster bootstrapping tool) use CNI as the default network interface mechanism.
    - CNI doesn't not help you with pod-to-pod communication across nodes
    - https://github.com/containernetworking/cni

```    
$ cat >/etc/cni/net.d/10-mynet.conf <<EOF
{
	"cniVersion": "0.2.0",
	"name": "mynet",
	"type": "bridge",
	"bridge": "cni0",
	"isGateway": true,
	"ipMasq": true,
	"ipam": {
		"type": "host-local",
		"subnet": "10.22.0.0/16",
		"routes": [
			{ "dst": "0.0.0.0/0" }
		]
	}
}
EOF
```

- Pod-to-Pod communication (Networking)
    - the requirement from K8s
         - All pods can communication eith eath other across nodes.
         - All nodes can communicate with all popds.
         - No "Network Address Translation(NAT)"
    - All IPs involved(nodes and pods) are routable without NAT.
    - Software defined overlay Solution
        - Weave
        - Flannel
        - Calico
        - Romana
- kubectl : main command line cliant
- Network Add-Ons
    - example: networking using Weave net.
        - "$ kubectl apply -f "https://cloud.weave.works/k8s/net?k8s-version=$(kubectl version | base64 | tr -d '\n')""
        - https://www.weave.works/docs/net/latest/kubernetes/kube-addon/
- Mesos Architecture
    - No diffrerent from K8s, at high level.
    - Persistence layer: Zookeeper (instead of etcd)
    - head node / worker nodes
- K8s command
    - check status of each module on master node
        - "systemctl list-units | grep kube"
        - "systemctl list-units | grep etcd"
        - "systemctl status kube-apiserver
            - return "active(running)"
        - "systemctl status kube-proxy
    - configuration
        - "vi /etc/systemd/system/kube-apiserver.service"
        - "vi /etc/systemd/system/kube-proxy.service"
    - check the inside-KVS
        - "etcdctl ls"
        - "etcdctl ls /registry"
        - "etcdctl ls /registry/namespaces"
        - "etcdctl ls /registry/pods/kube-system"
    - check status of each module on worker node
        - "docker ps"
        - "systemctl list-units | grep kube"
    - check Images name, status, version, networking
        - "kubectl get nodes"
        - "kubectl get pods --all-namespaces"
    - configure networking plugin (CNI) (Demo)
        - master node
            - "systecmctl status kubelet"
            - "more /etc/systemd/system/kubelet.service.d/10-kubeadm.conf"
            - "cd /etc/kubernetes/manifests/"
                - etcd.yaml
                - kube-apiserver.yaml
                - kube-contoller0manager.yaml
                - kube-scheduler.yaml
            - "kubectl apply -f ..."
            - kubectl get nodes
            - kubectl get pods --all-namespaces
            - ifconfig
        - worker node
            - "systecmctl status kubelet"
            - "more /etc/systemd/system/kubelet.service.d/10-kubeadm.conf"
            - "more /etc/cni/net.d/weave.conf"
            - "docker ps"
                - k8s_weave-npc..


# Chapter4 K8s install and configuratin
## Install kubectl

```
curl -LO https://storage.googleapis.com/kubernetes-release/release/`curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt`/bin/darwin/amd64/kubectl
```

config file: /.kube/config
- see cluster definitions (IP endopons)
- see credentials
- contexts

switch between context

```
kubectl config use-context foobar
```
## Use GKE

Install

```
curl https://sdk.cloud.google.com | bash
exec -l $SHELL
gcloud init
```

create k8s cluster

```
gcloud container clusters create linuxfoundation

Creating cluster linuxfoundation...done.
Created [https://container.googleapis.com/v1/projects/k8s-training-190416/zones/us-west1-a/clusters/linuxfoundation].
kubeconfig entry generated for linuxfoundation.
NAME             LOCATION    MASTER_VERSION  MASTER_IP      MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
linuxfoundation  us-west1-a  1.7.8-gke.0     35.203.149.93  n1-standard-1  1.7.8-gke.0   3          RUNNING
```

check cluster

```
gcloud container clusters list

NAME             LOCATION    MASTER_VERSION  MASTER_IP      MACHINE_TYPE   NODE_VERSION  NUM_NODES  STATUS
linuxfoundation  us-west1-a  1.7.8-gke.0     35.203.149.93  n1-standard-1  1.7.8-gke.0   3          RUNNING
```

check k8s node list

```
kubectl get nodes

NAME                                             STATUS    ROLES     AGE       VERSION
gke-linuxfoundation-default-pool-34a8474a-8szm   Ready     <none>    11m       v1.7.8-gke.0
gke-linuxfoundation-default-pool-34a8474a-lk20   Ready     <none>    11m       v1.7.8-gke.0
gke-linuxfoundation-default-pool-34a8474a-sk55   Ready     <none>    11m       v1.7.8-gke.0
```

delete cluster

```
gcloud container clusters delete linuxfoundation
```

## Using Minikube (local environmnet)

install (MacOS)

```
brew cask install minikube
```

Start k8s.

```
minikube start
```

check k8s nodes

```
kubectl get nodes
```

## Minikube internals
- Minikube is an open source project on github https://github.com/kubernetes/minikube#configuring-kubernetes
- it runs a single "Go" binary called "localkube".
- it start a VirtualBox VM that will contain a single nodes k8s deployment and docekr engine
- Minikube VM also runs Docker, to be able to run containers
    - k8s container on docker engine on Virtualbox

connect minikube(virtual box)

```
minikube ssh
```

```
docker ps

CONTAINER ID        IMAGE                                                 COMMAND                  CREATED              STATUS              PORTS               NAMES
7b1ef9c1f72b        gcr.io/google_containers/kubernetes-dashboard-amd64   "/dashboard --inse..."   21 seconds ago       Up 21 seconds                           k8s_kubernetes-dashboard_kubernetes-dashboard-hvr8b_kube-system_e112aabe-ebf3-11e7-a7f8-0800278459df_0
5934570a8d43        gcr.io/k8s-minikube/storage-provisioner               "/storage-provisioner"   About a minute ago   Up About a minute                       k8s_storage-provisioner_storage-provisioner_kube-system_e0b5c968-ebf3-11e7-a7f8-0800278459df_0
6e895b61950c        fed89e8b4248                                          "/sidecar --v=2 --..."   About a minute ago   Up About a minute                       k8s_sidecar_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
a7c09c35a8a6        459944ce8cc4                                          "/dnsmasq-nanny -v..."   About a minute ago   Up About a minute                       k8s_dnsmasq_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
2aa747d09357        512cd7425a73                                          "/kube-dns --domai..."   About a minute ago   Up About a minute                       k8s_kubedns_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
1195d46ba6c1        gcr.io/google_containers/pause-amd64:3.0              "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
4ab3164f900a        gcr.io/google_containers/pause-amd64:3.0              "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kubernetes-dashboard-hvr8b_kube-system_e112aabe-ebf3-11e7-a7f8-0800278459df_0
09fca79b2aee        gcr.io/google_containers/pause-amd64:3.0              "/pause"                 About a minute ago   Up About a minute                       k8s_POD_storage-provisioner_kube-system_e0b5c968-ebf3-11e7-a7f8-0800278459df_0
6c04d444c578        0a951668696f                                          "/opt/kube-addons.sh"    About a minute ago   Up About a minute                       k8s_kube-addon-manager_kube-addon-manager-minikube_kube-system_7b19c3ba446df5355649563d32723e4f_0
e3a281b3f484        gcr.io/google_containers/pause-amd64:3.0              "/pause"                 About a minute ago   Up About a minute                       k8s_POD_kube-addon-manager-minikube_kube-system_7b19c3ba446df5355649563d32723e4f_0
```

extract NAMES
```
$ docker ps --format "table {{.Names}}"

NAMES
k8s_kubernetes-dashboard_kubernetes-dashboard-hvr8b_kube-system_e112aabe-ebf3-11e7-a7f8-0800278459df_0
k8s_storage-provisioner_storage-provisioner_kube-system_e0b5c968-ebf3-11e7-a7f8-0800278459df_0
k8s_sidecar_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
k8s_dnsmasq_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
k8s_kubedns_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
k8s_POD_kube-dns-86f6f55dd5-tdz5s_kube-system_e12a8a1c-ebf3-11e7-a7f8-0800278459df_0
k8s_POD_kubernetes-dashboard-hvr8b_kube-system_e112aabe-ebf3-11e7-a7f8-0800278459df_0
k8s_POD_storage-provisioner_kube-system_e0b5c968-ebf3-11e7-a7f8-0800278459df_0
k8s_kube-addon-manager_kube-addon-manager-minikube_kube-system_7b19c3ba446df5355649563d32723e4f_0
k8s_POD_kube-addon-manager-minikube_kube-system_7b19c3ba446df5355649563d32723e4f_0
```

```
$ ps -aux | grep localkube

root      3594  9.2 18.0 11287024 370328 ?     Ssl  17:23   0:43 /usr/local/bin/localkube --dns-domain=cluster.local --node-ip=192.168.99.100 --generate-certs=false --logtostderr=true --enable-dns=false
docker    5480  0.0  0.0   9120   472 pts/0    S+   17:31   0:00 grep localkube
```

```
eval $(minikube docker-env)
```

## Minikube usage

usage sample

```
minikube dashboard : open the dashboard
minikube ip : get the IP address of your minikube VM.
```


```
Usage:
  minikube [command]

Available Commands:
  addons           Modify minikube's kubernetes addons
  cache            Add or delete an image from the local cache.
  completion       Outputs minikube shell completion for the given shell (bash or zsh)
  config           Modify minikube config
  dashboard        Opens/displays the kubernetes dashboard URL for your local cluster
  delete           Deletes a local kubernetes cluster
  docker-env       Sets up docker env variables; similar to '$(docker-machine env)'
  get-k8s-versions Gets the list of Kubernetes versions available for minikube when using the localkube bootstrapper
  ip               Retrieves the IP address of the running cluster
  logs             Gets the logs of the running localkube instance, used for debugging minikube, not user code
  mount            Mounts the specified directory into minikube
  profile          Profile sets the current minikube profile
  service          Gets the kubernetes URL(s) for the specified service in your local cluster
  ssh              Log into or run a command on a machine with SSH; similar to 'docker-machine ssh'
  ssh-key          Retrieve the ssh identity key path of the specified cluster
  start            Starts a local kubernetes cluster
  status           Gets the status of a local kubernetes cluster
  stop             Stops a running local kubernetes cluster
  update-check     Print current and latest version number
  update-context   Verify the IP address of the running cluster in kubeconfig.
  version          Print the version of minikube
```


## Lab3: Docker networking

```
docker run -d --name=source busybox sleep 3600
docker run -d --name=same-ip --net=container:source busybox sleep 3600
```

check the ip address of the each container

```
# Container 1
docker exec -ti source ifconfig
# container 2
docker exec -ti same-ip ifconfig
```

## Start your K8s head node

run etcd

```
docker run -d --name=k8s -p 8080:8080 gcr.io/google_containers/etcd:3.1.10 etcd --data-dir /var/lib/data
```

start the API server using "hyperkube"
 (But it doesn't up container. Error?)

```
docker run -d --net=container:k8s gcr.io/google_containers/hyperkube:v1.7.6 /apiserver --etcd-servers=http://127.0.0.1:2379 --service-cluster-ip-range=10.0.0.1/24 --insecure-bind-address=0.0.0.0 --insecure-pord=8080 --admission-control=AlwaysAdmit


docekr ps -a

CONTAINER ID        IMAGE                                       COMMAND                  CREATED             STATUS                       PORTS                                                       NAMES
33f968550713        gcr.io/google_containers/hyperkube:v1.7.6   "/apiserver --etcd..."   2 minutes ago       Exited (1) 2 minutes ago                                                                 youthful_almeida
```

Start the Admission contoroller

```
docker run -d --net=container:k8s gcr.io/google_containers/hyperkube:v1.7.6 /controller-manager --master=127.0.0.1:8080
```

To test that, use etcdctl in the etcd container

```
docker exec -ti k8s /bin/sh
export ETCDCTL_API=3
etcdctl get “/registry/api” --prefix=true
```

You can reach your Kubenetes API server and start exploring the API.
 (But it doesn't return. Error?)

```
curl http://127.0.0.1:8080/api/v1

curl: (52) Empty reply from server
```

## K8s application via Dashborad
Open Dashboad
- http://192.168.99.100:30000/#!/overview?namespace=default

```
minikube dashboard
```

Start App

- Workloads > Deployments > Deploy a Containerized App
- redis

## Install K8s with kubeadmin

1. "kubeadm init" on head nodes"
2. "kubeadm join --token <token> <head node IP address>" on worker nodes
3. "kubectl apply -f http:// ..."
    - use resource manifest of the network

## other installation mechanism
- kubespray
    - ansible playbook for k8s cluster
- kops
    - crate a k8s on AWS via CLI
- kube-aws
    - AWS cloud Formation to Provision K8s on AWS
- kubicorn
    - golong libray
    - provision and manage k8s culuster

Insltall K8s suing step by steo manual command
- https://github.com/kelseyhightower/kubernetes-the-hard-way


## Installation Collection
Single node
- Minikube
- kube-solo

## For deployment configuration
- Single-node deployment
    - all components run on the same server.
- Single head node and multiple workers
    - sigle node etcd instance running on the head node with, scheduler, API, controller manager
- Multiple head nodes in asn HA configuration and multiple workers
    -  API server will be fronted by a load balancer, scheduler and controller-manager will elect a leader
    - congitured via flags
    - etcd setp can be single-node
- HA etcd cluster, with HA head nodes and mutiple workers
    - etcd would run as true cluster
    - cluster provide HA and would run separate nodes than the K8s head nodes.

## Running components via containers using Hyperkube
- kubeadam can run the API server, scheduler, controller-manager as container
- hyperkube: all-in-one binary of them, as container images
    - gcr.io/google_containers/hyperkube:v1.7.6
- installation 
    - running kubelet as a system daemon

```
docker run --rm gcr.io/google_containers/hyperkube:v1.7.6 /hyperkube apiserver --help

docker run --rm gcr.io/google_containers/hyperkube:v1.7.6 /hyperkube scheduler --help

docker run --rm gcr.io/google_containers/hyperkube:v1.7.6 /hyperkube controller-manager --help
```
## Compiliing K8s from source

build natively with Golang

```
cd $GOPATH/src/github.com
git clone https://github.com/kubernetes/kubernetes
cd kubenetes
make
```

make quick-release
ls -al _output/bin

## Minikube Demo

```
which minikube
which kubectl
    - k8s client
which gcloud
```

```
kubectl logs xxx
kubectl run ghost --images-ghost
kubectl get pods
kubectl get deploymentes
kubectl config use-context minikube
kubectl expose deployment/ghost --port=2368 --type=NodePort 
kubectl get svc
kubectl describe svc ghost
minikube service ghost
kubectl config view
kubectl config use-context <contaier_id>
kubectl config use-context minikube
```

## Kubeadm demo
create k8s on ubuntu on DegitalOcean

https://kubernetes.io/docs/setup/independent/install-kubeadm/

https://kubernetes.io/docs/setup/independent/create-cluster-kubeadm/#14-installing-kubeadm-on-your-hosts

kubectl get ndoes
kubectl get ndoes --watch

# Chapter5 Accessing k8s cluster using API




## Lab4: Installtion using Minukube, kubeadm, manual
https://lms.quickstart.com/custom/858487/Lab4.pdf

# Q&A
## Chapter 1
## Chapter 2
- Q: On which of the following in K8s based on?
    - A: Borg
- Q: What language is K8s written in?
    - A: Go
- Q: whre is the cluster state stored?
    - A: etcd
## Chapter 3
## Chapter 4
## Chapter 5
## Chapter 6
## Chapter 7
## Chapter 8
## Chapter 9
## Chapter 10
## Chapter 11
## Chapter 12
## Chapter 13
## Chapter 14
## Chapter 15
## Chapter 16

# Others
Resource
- https://training.linuxfoundation.org/cm/LFS258/
Discussion Board
- https://www.linux.com/forums/lfs258-class-forum


# TODO
- suevery Network namespace

# Keywords
## kubadm
- a tool built to provide as best-practice "fast paths" for creating k8s cluster
- networking

## Hyperkube
- hyperkube is an all-in-one binary for the Kubernetes server components
- https://github.com/kubernetes/kubernetes/tree/master/cluster/images/hyperkube