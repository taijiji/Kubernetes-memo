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


## Lab4: Installtion using Minukube, kubeadm, manual
https://lms.quickstart.com/custom/858487/Lab4.pdf


# Chapter5 Accessing k8s cluster using API

## The API Endopoint
make sure that Minikube is running

```
$ minikube status

minikube: Running
cluster: Running
kubectl: Correctly Configured: pointing to minikube-vm at 192.168.99.100
```

run a local proxy to connect to the API server

```
$ kubectl proxy 

Starting to serve on 127.0.0.1:8001
```

```
$ curl http://localhost:8001

{
  "paths": [
    "/api",
    "/api/v1",
    "/apis",
    "/apis/",
    "/apis/admissionregistration.k8s.io",
    "/apis/admissionregistration.k8s.io/v1alpha1",
    "/apis/apiextensions.k8s.io",
    "/apis/apiextensions.k8s.io/v1beta1",
    "/apis/apiregistration.k8s.io",
    "/apis/apiregistration.k8s.io/v1beta1",
    "/apis/apps",
    "/apis/apps/v1beta1",
    "/apis/apps/v1beta2",
    "/apis/authentication.k8s.io",
    "/apis/authentication.k8s.io/v1",
    "/apis/authentication.k8s.io/v1beta1",
    "/apis/authorization.k8s.io",
    "/apis/authorization.k8s.io/v1",
    "/apis/authorization.k8s.io/v1beta1",
    "/apis/autoscaling",
    "/apis/autoscaling/v1",
    "/apis/autoscaling/v2beta1",
    "/apis/batch",
    "/apis/batch/v1",
    "/apis/batch/v1beta1",
    "/apis/batch/v2alpha1",
    "/apis/certificates.k8s.io",
    "/apis/certificates.k8s.io/v1beta1",
    "/apis/extensions",
    "/apis/extensions/v1beta1",
    "/apis/networking.k8s.io",
    "/apis/networking.k8s.io/v1",
    "/apis/policy",
    "/apis/policy/v1beta1",
    "/apis/rbac.authorization.k8s.io",
    "/apis/rbac.authorization.k8s.io/v1",
    "/apis/rbac.authorization.k8s.io/v1alpha1",
    "/apis/rbac.authorization.k8s.io/v1beta1",
    "/apis/scheduling.k8s.io",
    "/apis/scheduling.k8s.io/v1alpha1",
    "/apis/settings.k8s.io",
    "/apis/settings.k8s.io/v1alpha1",
    "/apis/storage.k8s.io",
    "/apis/storage.k8s.io/v1",
    "/apis/storage.k8s.io/v1beta1",
    "/healthz",
    "/healthz/autoregister-completion",
    "/healthz/etcd",
    "/healthz/ping",
    "/healthz/poststarthook/apiservice-openapi-controller",
    "/healthz/poststarthook/apiservice-registration-controller",
    "/healthz/poststarthook/apiservice-status-available-controller",
    "/healthz/poststarthook/bootstrap-controller",
    "/healthz/poststarthook/ca-registration",
    "/healthz/poststarthook/generic-apiserver-start-informers",
    "/healthz/poststarthook/kube-apiserver-autoregistration",
    "/healthz/poststarthook/start-apiextensions-controllers",
    "/healthz/poststarthook/start-apiextensions-informers",
    "/healthz/poststarthook/start-kube-aggregator-informers",
    "/healthz/poststarthook/start-kube-apiserver-informers",
    "/logs",
    "/metrics",
    "/swagger-2.0.0.json",
    "/swagger-2.0.0.pb-v1",
    "/swagger-2.0.0.pb-v1.gz",
    "/swagger.json",
    "/swaggerapi",
    "/ui",
    "/ui/",
    "/version"
  ]
}
```

## API Resources

find a pod.
API group a resource ia in specifies the apiVersion

```
$ curl http://localhost:8001/api/v1
{
    (snip)

    {
      "name": "pods",
      "singularName": "",
      "namespaced": true,
      "kind": "Pod",
      "verbs": [
        "create",
        "delete",
        "deletecollection",
        "get",
        "list",
        "patch",
        "proxy",
        "update",
        "watch"
      ],
      "shortNames": [
        "po"
      ],
      "categories": [
        "all"
      ]
    }

    (snip)
  ]
}%
```

Typically a resource manifests(available in JSON or YAML formats) will basic skeleton.
- kind:
- apiVersion:
- metadata:
- spec:

API Conventions
- https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md


## Your First Pod

an expample of a pod manifest in YAML format.

```
apiVersion: v1
kind: Pod
metadata:
    name: busybox
    namespace: default
spec:
    containers:
    - name: busybox
      image: busybox
      command:
        - sleep
        - "3600"
```

create this pod in K8s from YAML format.
 
```
$ kubectl create -f sample_pod.yaml

pod "busybox" created
```

```
kubectl get pods
NAME                     READY     STATUS    RESTARTS   AGE
busybox                  1/1       Running   0          29s
```

## Managing Resource with REST
REST = managed via HTTP calls
- GET
- POST
- PATCH
- DELETE
- etc

TO discovery the API

```
kubectl --v=99 get pods busybox

(snip)

I0104 07:53:27.135814   63287 round_trippers.go:417] curl -k -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.8.1 (darwin/amd64) kubernetes/f38e43b" https://192.168.99.100:8443/api/v1/namespaces/default/pods/busybox
I0104 07:53:27.147042   63287 round_trippers.go:436] GET https://192.168.99.100:8443/api/v1/namespaces/default/pods/busybox 200 OK in 11 milliseconds
```

If you delete this pod, HTTP method changes from GET to DELETE

URI : /api/v1/namespaces/default/pods/busybox

```
kubectl --v=99 delete pods busybox

(snip)

I0104 07:59:04.029377   63336 round_trippers.go:417] curl -k -v -XDELETE  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.8.1 (darwin/amd64) kubernetes/f38e43b" https://192.168.99.100:8443/api/v1/namespaces/default/pods/busybox
I0104 07:59:04.032576   63336 round_trippers.go:436] DELETE https://192.168.99.100:8443/api/v1/namespaces/default/pods/busybox 200 OK in 3 milliseconds
```

## Namespaces

Every Request is namespaced
- ex) GET https://192.168.99.100:84443/api/v1/namespaces/default/pods


```
kubectl get ns

NAME          STATUS    AGE
default       Active    6d
kube-public   Active    6d
kube-system   Active    6d
```

```
kubectl create ns linuxcon

namespace "linuxcon" created
```

```
kubectl get ns/linuxcon

NAME       STATUS    AGE
linuxcon   Active    53s
```

```
kubectl get ns/linuxcon -o yaml

apiVersion: v1
kind: Namespace
metadata:
  creationTimestamp: 2018-01-04T16:07:19Z
  name: linuxcon
  resourceVersion: "190082"
  selfLink: /api/v1/namespaces/linuxcon
  uid: 5773f71a-f169-11e7-a7f8-0800278459df
spec:
  finalizers:
  - kubernetes
status:
  phase: Active
```

```
kubectl get ns/linuxcon -o json

{
    "apiVersion": "v1",
    "kind": "Namespace",
    "metadata": {
        "creationTimestamp": "2018-01-04T16:07:19Z",
        "name": "linuxcon",
        "resourceVersion": "190082",
        "selfLink": "/api/v1/namespaces/linuxcon",
        "uid": "5773f71a-f169-11e7-a7f8-0800278459df"
    },
    "spec": {
        "finalizers": [
            "kubernetes"
        ]
    },
    "status": {
        "phase": "Active"
    }
}
```

```
kubectl delete ns/linuxcon

namespace "linuxcon" delete
```

sample: create a pod in the linuxcon namespace:

```
piVersion: v1
kind: Pod
metadata:
  name: redis
  namespace: linuxcon
...
```

## Check API Resource with kubectl

- kubectl get pods
- kubectl get namespaces
- kubectl get replicasets
- kubectl get deployments
- kubectl get services

```
kubectl get pods

NAME                     READY     STATUS    RESTARTS   AGE
redis-7f5f77dc44-qmx56   1/1       Running   0          8d
```

```
kubectl get namespaces
NAME          STATUS    AGE
default       Active    9d
kube-public   Active    9d
kube-system   Active    9d
```

```
kubectl get replicasets
NAME               DESIRED   CURRENT   READY     AGE
redis-7f5f77dc44   1         1         1         8d
```

```
kubectl get deployments
NAME      DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
redis     1         1         1            1           8d
```

```
kubectl get services
NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)    AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP    9d
redis        ClusterIP   10.108.112.135   <none>        8002/TCP   8d
```

detailed manuffest of a specific resource
- kubectl get pods busybox -o yaml

```
kubectl get pods redis-7f5f77dc44-qmx56 -o yaml
apiVersion: v1
kind: Pod
metadata:
  annotations:

(snip)
```

## Swagger specification and Open API

K8s APU use a Swagger specification
- https://swagger.io/specification/
- this evolving toward the OpenAPI

OpenAPI initiative
- It is usefule, auto-generate client code
- https://www.openapis.org/

## Additional Resource Management Methods
access log with endpoins
- GET /api/v1/namespaces/{namespace}/pods/{name}/exec
- GET /api/v1/namespaces/{namespace}/pods/{name}/logs
- GET /api/v1/namespaces/{namespace}/pods/{name}/portforward
- GET /api/v1/watch/{namespace}/pods/{name}

Kubectl: access the logs of a container
- kubectl logs --help
- kubectl exec --help
- kubectl port-forward --help
- kubectl attach --help

```
kubectl logs --help
Print the logs for a container in a pod or specified resource. If the pod has only one container, the container name isoptional.

Aliases:
logs, log

Examples:
  # Return snapshot logs from pod nginx with only one container
  kubectl logs nginx

  # Return snapshot logs for the pods defined by label app=nginx
  kubectl logs -lapp=nginx

  # Return snapshot of previous terminated ruby container logs from pod web-1
  kubectl logs -p -c ruby web-1

  # Begin streaming the logs of the ruby container in pod web-1
  kubectl logs -f -c ruby web-1

  # Display only the most recent 20 lines of output in pod nginx
  kubectl logs --tail=20 nginx

  # Show all logs from pod nginx written in the last hour
  kubectl logs --since=1h nginx

  # Return snapshot logs from first container of a job named hello
  kubectl logs job/hello

  # Return snapshot logs from container nginx-1 of a deployment named nginx
  kubectl logs deployment/nginx -c nginx-1

Options:
  -c, --container='': Print the logs of this container
  -f, --follow=false: Specify if the logs should be streamed.
      --include-extended-apis=true: If true, include definitions of new APIs via calls to the API server. [default true]
      --interactive=false: If true, prompt the user for input when required.
      --limit-bytes=0: Maximum bytes of logs to return. Defaults to no limit.
      --pod-running-timeout=20s: The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running
  -p, --previous=false: If true, print the logs for the previous instance of the container in a pod if it exists.
  -l, --selector='': Selector (label query) to filter on.
      --since=0s: Only return logs newer than a relative duration like 5s, 2m, or 3h. Defaults to all logs. Only one ofsince-time / since may be used.
      --since-time='': Only return logs after a specific date (RFC3339). Defaults to all logs. Only one of since-time /since may be used.
      --tail=-1: Lines of recent log file to display. Defaults to -1 with no selector, showing all log lines otherwise 10, if a selector is provided.
      --timestamps=false: Include timestamps on each line in the log output

Usage:
  kubectl logs [-f] [-p] (POD | TYPE/NAME) [-c CONTAINER] [options]

Use "kubectl options" for a list of global command-line options (applies to all commands).
```

```
kubectl exec --help
Execute a command in a container.

Examples:
  # Get output from running 'date' from pod 123456-7890, using the first container by default
  kubectl exec 123456-7890 date

  # Get output from running 'date' in ruby-container from pod 123456-7890
  kubectl exec 123456-7890 -c ruby-container date

  # Switch to raw terminal mode, sends stdin to 'bash' in ruby-container from pod 123456-7890
  # and sends stdout/stderr from 'bash' back to the client
  kubectl exec 123456-7890 -c ruby-container -i -t -- bash -il

  # List contents of /usr from the first container of pod 123456-7890 and sort by modification time.
  # If the command you want to execute in the pod has any flags in common (e.g. -i),
  # you must use two dashes (--) to separate your command'sflags/arguments.
  # Also note, do not surround your command and its flags/arguments with quotes
  # unless that is how you would execute it normally (i.e.,do ls -t /usr, not "ls -t /usr").
  kubectl exec 123456-7890 -i -t -- ls -t /usr

Options:
  -c, --container='': Container name. If omitted, the firstcontainer in the pod will be chosen
  -p, --pod='': Pod name
  -i, --stdin=false: Pass stdin to the container
  -t, --tty=false: Stdin is a TTY

Usage:
  kubectl exec POD [-c CONTAINER] -- COMMAND [args...] [options]

Use "kubectl options" for a list of global command-line options (applies to all commands).
```

```

➜  Kubernetes-memo git:(master) ✗ kubectl port-forward --help
Forward one or more local ports to a pod.

Examples:
  # Listen on ports 5000 and 6000 locally, forwarding data to/from ports 5000 and 6000 in the pod
  kubectl port-forward mypod 5000 6000

  # Listen on port 8888 locally, forwarding to 5000 in the pod
  kubectl port-forward mypod 8888:5000

  # Listen on a random port locally, forwarding to 5000 in the pod
  kubectl port-forward mypod :5000

Options:
  -p, --pod='': Pod name

Usage:
  kubectl port-forward POD [LOCAL_PORT:]REMOTE_PORT [...[LOCAL_PORT_N:]REMOTE_PORT_N] [options]

Use "kubectl options" for a list of global command-line options (applies to all commands).
```

```
kubectl attach --help
Attach to a process that is already running inside an existing container.

Examples:
  # Get output from running pod 123456-7890, using the first container by default
  kubectl attach 123456-7890

  # Get output from ruby-container from pod 123456-7890
  kubectl attach 123456-7890 -c ruby-container

  # Switch to raw terminal mode, sends stdin to 'bash' in ruby-container from pod 123456-7890
  # and sends stdout/stderr from 'bash' back to the client
  kubectl attach 123456-7890 -c ruby-container -i -t

  # Get output from the first pod of a ReplicaSet named nginx
  kubectl attach rs/nginx

Options:
  -c, --container='': Container name. If omitted, the firstcontainer in the pod will be chosen
      --pod-running-timeout=1m0s: The length of time (like 5s, 2m, or 3h, higher than zero) to wait until at least one pod is running
  -i, --stdin=false: Pass stdin to the container
  -t, --tty=false: Stdin is a TTY

Usage:
  kubectl attach (POD | TYPE/NAME) -c CONTAINER [options]

Use "kubectl options" for a list of global command-line options (applies to all commands).
```

## Pod Logs and Exec
check log of the container

```
kubectl get pods

NAME                     READY     STATUS    RESTARTS   AGE
redis-7f5f77dc44-qmx56   1/1       Running   0          8d
```

```
kubectl logs redis-7f5f77dc44-qmx56

1:C 29 Dec 19:34:39.877 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 29 Dec 19:34:39.877 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 29 Dec 19:34:39.877 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
1:M 29 Dec 19:34:39.878 * Running mode=standalone, port=6379.
1:M 29 Dec 19:34:39.879 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:M 29 Dec 19:34:39.879 # Server initialized
1:M 29 Dec 19:34:39.879 * Ready to accept connections
```

```
kubectl exec -ti redis-7f5f77dc44-qmx56 bash

root@redis-7f5f77dc44-qmx56:/data#
root@redis-7f5f77dc44-qmx56:/data# redis-cli
127.0.0.1:6379>
```

## Connecting to a Pod with Port Forwarding (Not exec)

```
kubectl get pods

NAME                     READY     STATUS    RESTARTS   AGE
redis-7f5f77dc44-qmx56   1/1       Running   0          9d
```

```
kubectl port-forward redis-7f5f77dc44-qmx56 6379:6379

Forwarding from 127.0.0.1:6379 -> 6379
```

## API demo

check API 

```
minikube ssh
```

```
curl localhost:8080
curl localhost:8080/version
curl localhost:8080/api/v1
```

```
exit
```

namespace

```
kubectl get ns

NAME          STATUS    AGE
default       Active    10d
kube-public   Active    10d
kube-system   Active    10d
```

```
 kubectl get pods --all-namespaces


NAMESPACE     NAME                          READY     STATUS    RESTARTS   AGE
default       redis-7f5f77dc44-qmx56        1/1       Running   0          9d
kube-system   kube-addon-manager-minikube   1/1       Running   0          10d
kube-system   kube-dns-86f6f55dd5-tdz5s     3/3       Running   0          10d
kube-system   kubernetes-dashboard-hvr8b    1/1       Running   0          10d
kube-system   storage-provisioner           1/1       Running   0          10d
```

create namespace

```
kubectl create ns foobar

namespace "foobar" created
```

```
kubectl get ns

NAME          STATUS    AGE
default       Active    10d
foobar        Active    2m
kube-public   Active    10d
kube-system   Active    10d
```

create namespace from yaml file

```
vi ns.yaml

apiVersion: v1
kind: Namespace
metadata:
  name: lfs258
```

```
kubectl create -f ns.yaml

namespace "lfs258" created
```

```
kubectl get ns

NAME          STATUS    AGE
default       Active    10d
foobar        Active    6m
kube-public   Active    10d
kube-system   Active    10d
lfs258        Active    2m
```

## LAb5
https://lms.quickstart.com/custom/858487/Lab5.pdf


create pod in a namespace

```
vi redis-lfs258.yaml

apiVersion: v1
kind: Pod
metadata:
  namespace: lfs258
  name: redis
spec:
  containers:
    - name: redis
      image: redis
```

```
kubectl create -f redis-lfs258.yaml

pod "redis" created
```

```
kubectl get pods --namespace=lfs258

NAME      READY     STATUS    RESTARTS   AGE
redis     1/1       Running   0          5m
```

```
kubectl get pods --all-namespaces
NAMESPACE     NAME                          READY     STATUS    RESTARTS   AGE
default       redis-7f5f77dc44-qmx56        1/1       Running   0          9d
kube-system   kube-addon-manager-minikube   1/1       Running   0          10d
kube-system   kube-dns-86f6f55dd5-tdz5s     3/3       Running   0          10d
kube-system   kubernetes-dashboard-hvr8b    1/1       Running   0          10d
kube-system   storage-provisioner           1/1       Running   0          10d
lfs258        redis                         1/1       Running   0          5m
```

log

```
kubectl logs redis --namespace=lfs258
1:C 07 Jan 20:04:26.215 # oO0OoO0OoO0Oo Redis is starting oO0OoO0OoO0Oo
1:C 07 Jan 20:04:26.215 # Redis version=4.0.6, bits=64, commit=00000000, modified=0, pid=1, just started
1:C 07 Jan 20:04:26.215 # Warning: no config file specified, using the default config. In order to specify a config file use redis-server /path/to/redis.conf
1:M 07 Jan 20:04:26.216 * Running mode=standalone, port=6379.
1:M 07 Jan 20:04:26.217 # WARNING: The TCP backlog setting of 511 cannot be enforced because /proc/sys/net/core/somaxconn is set to the lower value of 128.
1:M 07 Jan 20:04:26.217 # Server initialized
1:M 07 Jan 20:04:26.217 * Ready to accept connections
```

delete namespace, pods

```
kubectl delete pods redis

pod "redis" deleted
```

```
kubectl delete pods redis --namespace=lfs258

pod "redis" deleted
```

```
kubectl delete ns lfs258

namespace "lfs258" deleted
```


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