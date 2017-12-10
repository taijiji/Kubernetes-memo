# LF258 Cource
https://training.linuxfoundation.org/linux-courses/system-administration-training/kubernetes-fundamentals

# mypage
https://training.linuxfoundation.org/cas?destination=portal

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


# Q&A
## Chapter 1
- Q: On which of the following in K8s based on?
    - A: Borg
- Q: What language is K8s written in?
    - A: Go
- Q: whre is the cluster state stored?
    - A: etcd
## Chapter 2
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


