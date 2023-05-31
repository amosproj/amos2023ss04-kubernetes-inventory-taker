# Proxy information

## Prequesits

Since the proxy Interacts with the kubernetes api of a cluster, you need to have one up and running (like [minikube](https://minikube.sigs.k8s.io/docs/start/)).

## Building

To be able to build the proxy docker image, you have to supply a kubernetes config file with all the required authentication files to be able to access the api.
The config file needs to be stored in `Proxy/.kube/config` as well as edited to reflect the destination of the certificates and keys. (Should be `/.kube/<filename>` )

The following is an excerpt from my modified minikube config:

```yaml
apiVersion: v1
clusters:
  - cluster:
      certificate-authority: /.kube/ca.crt
---
contexts:
  - context:
      cluster: minikube
---
users:
  - name: minikube
    user:
      client-certificate: /.kube/client.crt
      client-key: /.kube/client.key
```

## Running

To run the proxy with access to the host network stack you need to use:

```bash
docker run --network host amos2023ss04-kubernetes-inventory-taker-proxy:latest
```
