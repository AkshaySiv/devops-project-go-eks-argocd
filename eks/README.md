<!-- 
This section provides instructions to create an Amazon EKS cluster using `eksctl`. 
The command utilizes a specified cluster name and AWS region to set up the cluster. 
Ensure that `eksctl` is installed and configured with the necessary permissions 
before running the command.
-->
Create the cluster using the config file:


```bash
eksctl create cluster --config=cluster-t3medium.yaml
```
