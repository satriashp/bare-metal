# enable sudoers

sudo apt-get install vim -y
sudo vim /etc/sudoers.d/shp
shp ALL=(ALL) NOPASSWD: ALL

# kubeadm init (control-plane)

sudo kubeadm init --config init.yaml

# kubeadm init (worker-node)

sudo kubeadm join --config join.yaml

## Approve kubelet CSR (Certificate Signing Requests)

Test connection to kubernetes cluster, then approve CSR request from kubelet

```bash
# test connection
export KUBECONFIG=/etc/kubernetes/admin.conf
kubectl cluster-info

# Approve kubelet CSR
kubectl get csr --no-headers \
  | awk '$3=="kubernetes.io/kubelet-serving" && $6=="Pending" {print $1}' \
  | xargs kubectl certificate approve
```

## Install CNI

Install a Container Network Interface (CNI) plugin. Example using cilium(ref):

```bash
# via cilium cli
export KUBECONFIG=~/.kube/admin.conf
kubectl cluster-info

cilium install 1.20.1 --set ipam.operator.clusterPoolIPv4PodCIDRList="10.244.0.0/16"
```

helm install cilium oci://quay.io/cilium/charts/cilium 1.20.1 --namespace kube-system --values /home/shp/Learns/bare-metal/kubeadm/cilium-values.yaml
