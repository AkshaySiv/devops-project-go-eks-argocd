# DevOps Project: Go EKS ArgoCD

## Overview
This project demonstrates a DevOps workflow using **Go**, **Amazon EKS (Elastic Kubernetes Service)**, and **ArgoCD** for continuous delivery and deployment.

## Features
- **Go Application**: A sample Go application to deploy.
- **EKS Integration**: Kubernetes cluster on AWS for container orchestration.
- **ArgoCD**: GitOps-based continuous delivery tool for Kubernetes.

## Prerequisites
- AWS CLI configured with appropriate permissions.
- kubectl installed and configured.
- ArgoCD CLI installed.
- Docker installed for containerization.
- Go installed for application development.

## Setup Instructions

### 1. Clone the Repository
```bash
git clone https://github.com/your-username/devops-project-go-eks-argocd.git
cd devops-project-go-eks-argocd
```

### 3. Steps to run the project locally



### 3. Create EKS Cluster
- Create the EKS cluster using `eksctl`:
```bash
eksctl create cluster -f eks/cluster-t3medium.yaml
```
![alt text](images/eksctl.png)

- Verify the cluster is running:
```bash
kubectl get nodes
```


### 3. Install NGINX Ingress Controller
- Apply the NGINX Ingress Controller manifest:
```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.12.2/deploy/static/provider/aws/deploy.yaml
```
- Verify the ingress controller pods are running:
```bash
kubectl get pods -n ingress-nginx
```
- Check the service to get the external IP:
```bash
kubectl get svc -n ingress-nginx
```
- Use the external IP to access your application.





### 4. Set Up ArgoCD
- Install ArgoCD in your cluster:
```bash
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
- Access the ArgoCD UI and configure the application.





## Directory Structure
```
devops-project-go-eks-argocd/
├── app/                # Go application source code
├── k8s/                # Kubernetes manifests
├── Dockerfile          # Dockerfile for the Go app
├── README.md           # Project documentation
```

## Contributing
Contributions are welcome! Please fork the repository and submit a pull request