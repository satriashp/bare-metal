# Installation

# Ensure kube-proxy set mode to ipvs and has strictARP set to true.
# Edit configmap/kube-proxy
apiVersion: kubeproxy.config.k8s.io/v1alpha1
kind: KubeProxyConfiguration
mode: "ipvs"
ipvs:
  strictARP: true

# Restart kube-proxy daemonset
kubectl rollout restart daemonset kube-proxy -n kube-system

# Download metal-lb manifest
wget -O metallb-native.yaml https://raw.githubusercontent.com/metallb/metallb/v0.15.3/config/manifests/metallb-native.yaml

# Apply configuration
kubectl apply -f ./metallb-native.yaml
## Adjust ip addresses pool;
kubectl apply -f ./config.yaml
