# Tutorial on Google Container Engine
Reference : Web DB Press vol99,2017

# Install gcloud
https://cloud.google.com/sdk/
- Reference: https://cloud.google.com/sdk/docs/


- Download https://cloud.google.com/sdk/docs/quickstart-mac-os-x


'''
mv google-cloud-sdk /Users/taiji/work/
'''

'''
./google-cloud-sdk/install.sh
'''

'''
exec -l $SHELL
'''

'''
gcloud components update
'''

# Global setting

'''
gcloud config set project testk8s-177117
gcloud config set compute/zone us-east1-b
'''

# Account setting

'''
cloud auth login
'''

# Create Kubernetes cluster
3 nodes in 1 cluster, g1-small

'''
gcloud container clusters create testk8s-177117 --cluster-version=1.7.3 --machine-type=g1-small
'''

download auth data to connect 
'''
gcloud container clusters get-credentials testk8s-177117
'''

# Install Kuberctl

'''
gcloud components install kubectl
'''

'''
mba% kubectl version --short

Client Version: v1.7.3
Server Version: v1.7.3
'''

'''
kubectl cluster-info
'''

# Hello world

Hello world in go lang

'''
vi main.go
'''

make Docker file

'''
vi Dockerfile
'''

make cloud builder config
'''
vi cloudbuild.yaml
'''

Build by Google Cloud Builder
'''
gcloud container builds submit --config=cloudbuild.yaml .
'''

# Deploy 1 pod (Not use ReplicaSet)
make manufest file

'''
vi pod yaml
'''

make pod by kubectl

'''
kubectl create -f pod.yaml

or 

kubectl run hello-world --image=us.gcr.io/$PROJECT_ID/wdpress/hello-world --port=8080
'''





