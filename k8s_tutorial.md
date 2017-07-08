# Tutorial
Official Tutorial page

[Using Minikube to Create a Cluster](https://kubernetes.io/docs/tutorials/kubernetes-basics/cluster-intro/)

# Minikube?
Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a VM on your laptop for users looking to try out Kubernetes or develop with it day-to-day.


# Install for mac

1. install xhyve (Mac OS Hypervisor software)

```
brew install docker-machine-driver-xhyve
```

1. install minikube

```
brew cask install minikube
```

try minikube command

```
% minikube version                                                                  

minikube version: v0.19.1
```

1. Quickstart minikube

founded Error about dupulex IP addoress with VirtualBox

```
% minikube start                                                                      (git)-[master]
Starting local Kubernetes v1.6.4 cluster...
Starting VM...
Downloading Minikube ISO
 89.51 MB / 89.51 MB [==============================================] 100.00% 0s
E0619 08:37:30.289211   90303 start.go:127] Error starting host: Error creating host: Error with pre-create check: "VirtualBox is configured with multiple host-only adapters with the same IP \"192.168.33.1\". Please remove one".

 Retrying.
E0619 08:37:30.290661   90303 start.go:133] Error starting host:  Error creating host: Error with pre-create check: "VirtualBox is configured with multiple host-only adapters with the same IP \"192.168.33.1\". Please remove one"
================================================================================
An error has occurred. Would you like to opt in to sending anonymized crash
information to minikube to help prevent future errors?
To opt out of these messages, run the command:
	minikube config set WantReportErrorPrompt false
================================================================================
```


solution
```
http://qiita.com/tukiyo3/items/24dff2b3c234382a83e2
```

```
% /sbin/ifconfig vboxnet0                                                             (git)-[master]
vboxnet0: flags=8843<UP,BROADCAST,RUNNING,SIMPLEX,MULTICAST> mtu 1500
	ether 0a:00:27:00:00:00
	inet 192.168.33.1 netmask 0xffffff00 broadcast 192.168.33.255
```

