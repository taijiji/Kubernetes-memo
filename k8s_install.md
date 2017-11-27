# Installation pattern

# Kubernetes on docker on CoreOS

1. install CoreOS on vagrant
https://coreos.com/os/docs/latest/booting-on-vagrant.html

Clone Vagrant repo

'''
git clone https://github.com/coreos/coreos-vagrant.git

cd coreos-vagrant

vagrant up
'''

'''
vagrant status
Current machine states:

core-01                   running (virtualbox)
'''

```
vagrant ssh

pwd
/home/core
```

2. generate kubernetes on CoreOS
https://github.com/coreos/coreos-kubernetes/blob/master/Documentation/kubernetes-on-generic-platforms.md


download controller-install.sh (or Copy & Paste)
https://github.com/coreos/coreos-kubernetes/blob/master/multi-node/generic/controller-install.sh

run controller-install.sh

'''
'''

install CoreOS on mac
https://github.com/TheNewNormal/coreos-osx

Install CoreOS ISO image
https://coreos.com/os/docs/latest/booting-with-iso.html

