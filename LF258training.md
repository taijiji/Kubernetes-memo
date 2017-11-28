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
        - Borg and Kubernetes address these principles as well

# Others
Resource
- https://training.linuxfoundation.org/cm/LFS258/

Discussion Board
- https://www.linux.com/forums/lfs258-class-forum

