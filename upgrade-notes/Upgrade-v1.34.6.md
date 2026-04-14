# Upgrade notes
# Example kubeadm upgrade plan message output

kubeadm upgrade plan --kubeconfig /etc/kubernetes/admin.conf

[preflight] Running pre-flight checks.
[upgrade/config] Reading configuration from the "kubeadm-config" ConfigMap in namespace "kube-system"...
[upgrade/config] Use 'kubeadm init phase upload-config kubeadm --config your-config-file' to re-upload it.
[upgrade] Running cluster health checks
[upgrade] Fetching available versions to upgrade to
[upgrade/versions] Cluster version: 1.34.2
[upgrade/versions] kubeadm version: v1.34.2
I0414 13:51:09.947369  953921 version.go:260] remote version is much newer: v1.35.3; falling back to: stable-1.34
[upgrade/versions] Target version: v1.34.6
[upgrade/versions] Latest version in the v1.34 series: v1.34.6

Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':
COMPONENT   NODE             CURRENT   TARGET
kubelet     control-plane    v1.34.2   v1.34.6
kubelet     worker-node-01   v1.34.2   v1.34.6
kubelet     worker-node-02   v1.34.2   v1.34.6

Upgrade to the latest version in the v1.34 series:

COMPONENT                 NODE            CURRENT   TARGET
kube-apiserver            control-plane   v1.34.2   v1.34.6
kube-controller-manager   control-plane   v1.34.2   v1.34.6
kube-scheduler            control-plane   v1.34.2   v1.34.6
kube-proxy                                1.34.2    v1.34.6
CoreDNS                                   v1.12.1   v1.12.1
etcd                      control-plane   3.6.5-0   3.6.5-0

You can now apply the upgrade by executing the following command:

        kubeadm upgrade apply v1.34.6

Note: Before you can perform this upgrade, you have to update kubeadm to v1.34.6.

_____________________________________________________________________


The table below shows the current state of component configs as understood by this version of kubeadm.
Configs that have a "yes" mark in the "MANUAL UPGRADE REQUIRED" column require manual config upgrade or
resetting to kubeadm defaults before a successful upgrade can be performed. The version to manually
upgrade to is denoted in the "PREFERRED VERSION" column.

API GROUP                 CURRENT VERSION   PREFERRED VERSION   MANUAL UPGRADE REQUIRED
kubeproxy.config.k8s.io   v1alpha1          v1alpha1            no
kubelet.config.k8s.io     v1beta1           v1beta1             no
_____________________________________________________________________


# 1. Upgrade kubeadm binary

curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubeadm
mv kubeadm /usr/local/bin/kubeadm
chmod +x /usr/local/bin/kubeadm

# 2. Run upgrade plan
kubeadm upgrade plan --kubeconfig /etc/kubernetes/admin.conf

[preflight] Running pre-flight checks.
[upgrade/config] Reading configuration from the "kubeadm-config" ConfigMap in namespace "kube-system"...
[upgrade/config] Use 'kubeadm init phase upload-config kubeadm --config your-config-file' to re-upload it.
[upgrade] Running cluster health checks
[upgrade] Fetching available versions to upgrade to
[upgrade/versions] Cluster version: 1.34.2
[upgrade/versions] kubeadm version: v1.34.6
I0414 14:15:38.422526  974739 version.go:260] remote version is much newer: v1.35.3; falling back to: stable-1.34
[upgrade/versions] Target version: v1.34.6
[upgrade/versions] Latest version in the v1.34 series: v1.34.6

Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':
COMPONENT   NODE             CURRENT   TARGET
kubelet     control-plane    v1.34.2   v1.34.6
kubelet     worker-node-01   v1.34.2   v1.34.6
kubelet     worker-node-02   v1.34.2   v1.34.6

Upgrade to the latest version in the v1.34 series:

COMPONENT                 NODE            CURRENT   TARGET
kube-apiserver            control-plane   v1.34.2   v1.34.6
kube-controller-manager   control-plane   v1.34.2   v1.34.6
kube-scheduler            control-plane   v1.34.2   v1.34.6
kube-proxy                                1.34.2    v1.34.6
CoreDNS                                   v1.12.1   v1.12.1
etcd                      control-plane   3.6.5-0   3.6.5-0

You can now apply the upgrade by executing the following command:

        kubeadm upgrade apply v1.34.6

_____________________________________________________________________


The table below shows the current state of component configs as understood by this version of kubeadm.
Configs that have a "yes" mark in the "MANUAL UPGRADE REQUIRED" column require manual config upgrade or
resetting to kubeadm defaults before a successful upgrade can be performed. The version to manually
upgrade to is denoted in the "PREFERRED VERSION" column.

API GROUP                 CURRENT VERSION   PREFERRED VERSION   MANUAL UPGRADE REQUIRED
kubeproxy.config.k8s.io   v1alpha1          v1alpha1            no
kubelet.config.k8s.io     v1beta1           v1beta1             no
_____________________________________________________________________

# 3. Apply upgrade
kubeadm upgrade apply v1.34.6 --kubeconfig /etc/kubernetes/admin.conf

[upgrade] Reading configuration from the "kubeadm-config" ConfigMap in namespace "kube-system"...
[upgrade] Use 'kubeadm init phase upload-config kubeadm --config your-config-file' to re-upload it.
[upgrade/preflight] Running preflight checks
[upgrade] Running cluster health checks
[upgrade/preflight] You have chosen to upgrade the cluster version to "v1.34.6"
[upgrade/versions] Cluster version: v1.34.2
[upgrade/versions] kubeadm version: v1.34.6
[upgrade] Are you sure you want to proceed? [y/N]: y
[upgrade/preflight] Pulling images required for setting up a Kubernetes cluster
[upgrade/preflight] This might take a minute or two, depending on the speed of your internet connection
[upgrade/preflight] You can also perform this action beforehand using 'kubeadm config images pull'
[upgrade/control-plane] Upgrading your static Pod-hosted control plane to version "v1.34.6" (timeout: 5m0s)...
[upgrade/staticpods] Writing new Static Pod manifests to "/etc/kubernetes/tmp/kubeadm-upgraded-manifests4056920248"
[upgrade/staticpods] Preparing for "etcd" upgrade
[upgrade/staticpods] Renewing etcd-server certificate
[upgrade/staticpods] Renewing etcd-peer certificate
[upgrade/staticpods] Renewing etcd-healthcheck-client certificate
[upgrade/staticpods] Restarting the etcd static pod and backing up its manifest to "/etc/kubernetes/tmp/kubeadm-backup-manifests-2026-04-14-14-17-07/etcd.yaml"
[upgrade/staticpods] Waiting for the kubelet to restart the component
[upgrade/staticpods] This can take up to 5m0s
[apiclient] Found 1 Pods for label selector component=etcd
[upgrade/staticpods] Component "etcd" upgraded successfully!
[upgrade/etcd] Waiting for etcd to become available
[upgrade/staticpods] Preparing for "kube-apiserver" upgrade
[upgrade/staticpods] Renewing apiserver certificate
[upgrade/staticpods] Renewing apiserver-kubelet-client certificate
[upgrade/staticpods] Renewing front-proxy-client certificate
[upgrade/staticpods] Renewing apiserver-etcd-client certificate
[upgrade/staticpods] Moving new manifest to "/etc/kubernetes/manifests/kube-apiserver.yaml" and backing up old manifest to "/etc/kubernetes/tmp/kubeadm-backup-manifests-2026-04-14-14-17-07/kube-apiserver.yaml"
[upgrade/staticpods] Waiting for the kubelet to restart the component
[upgrade/staticpods] This can take up to 5m0s
[apiclient] Found 1 Pods for label selector component=kube-apiserver
[upgrade/staticpods] Component "kube-apiserver" upgraded successfully!
[upgrade/staticpods] Preparing for "kube-controller-manager" upgrade
[upgrade/staticpods] Renewing controller-manager.conf certificate
[upgrade/staticpods] Moving new manifest to "/etc/kubernetes/manifests/kube-controller-manager.yaml" and backing up old manifest to "/etc/kubernetes/tmp/kubeadm-backup-manifests-2026-04-14-14-17-07/kube-controller-manager.yaml"
[upgrade/staticpods] Waiting for the kubelet to restart the component
[upgrade/staticpods] This can take up to 5m0s
[apiclient] Found 1 Pods for label selector component=kube-controller-manager
[upgrade/staticpods] Component "kube-controller-manager" upgraded successfully!
[upgrade/staticpods] Preparing for "kube-scheduler" upgrade
[upgrade/staticpods] Renewing scheduler.conf certificate
[upgrade/staticpods] Moving new manifest to "/etc/kubernetes/manifests/kube-scheduler.yaml" and backing up old manifest to "/etc/kubernetes/tmp/kubeadm-backup-manifests-2026-04-14-14-17-07/kube-scheduler.yaml"
[upgrade/staticpods] Waiting for the kubelet to restart the component
[upgrade/staticpods] This can take up to 5m0s
[apiclient] Found 1 Pods for label selector component=kube-scheduler
[upgrade/staticpods] Component "kube-scheduler" upgraded successfully!
[upgrade/control-plane] The control plane instance for this node was successfully upgraded!
[upload-config] Storing the configuration used in ConfigMap "kubeadm-config" in the "kube-system" Namespace
[kubelet] Creating a ConfigMap "kubelet-config" in namespace kube-system with the configuration for the kubelets in the cluster
[upgrade/kubeconfig] The kubeconfig files for this node were successfully upgraded!
W0414 14:18:27.848671  975448 postupgrade.go:116] Using temporary directory /etc/kubernetes/tmp/kubeadm-kubelet-config510466628 for kubelet config. To override it set the environment variable KUBEADM_UPGRADE_DRYRUN_DIR
[upgrade] Backing up kubelet config file to /etc/kubernetes/tmp/kubeadm-kubelet-config510466628/config.yaml
[patches] Applied patch of type "application/strategic-merge-patch+json" to target "kubeletconfiguration"
[kubelet-start] Writing kubelet configuration to file "/var/lib/kubelet/config.yaml"
[upgrade/kubelet-config] The kubelet configuration for this node was successfully upgraded!
[upgrade/bootstrap-token] Configuring bootstrap token and cluster-info RBAC rules
[bootstrap-token] Configured RBAC rules to allow Node Bootstrap tokens to get nodes
[bootstrap-token] Configured RBAC rules to allow Node Bootstrap tokens to post CSRs in order for nodes to get long term certificate credentials
[bootstrap-token] Configured RBAC rules to allow the csrapprover controller automatically approve CSRs from a Node Bootstrap Token
[bootstrap-token] Configured RBAC rules to allow certificate rotation for all node client certificates in the cluster
[addons] Applied essential addon: CoreDNS
[addons] Applied essential addon: kube-proxy

[upgrade] SUCCESS! A control plane node of your cluster was upgraded to "v1.34.6".

[upgrade] Now please proceed with upgrading the rest of the nodes by following the right order.

# 4. Upgrade kubelet

curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubelet
mv kubelet /usr/local/bin/kubelet
chmod +x /usr/local/bin/kubelet
systemctl restart kubelet.service

## kubectl upgrade
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubectl
mv kubectl /usr/local/bin/kubectl
chmod +x /usr/local/bin/kubectl

# 5. Upgrade remaining nodes

# ################### worker-node-01 ###################
## kubeadm
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubeadm
mv kubeadm /usr/local/bin/kubeadm
chmod +x /usr/local/bin/kubeadm

## drain node
kubectl drain worker-node-01 --ignore-daemonsets --delete-emptydir-data

## upgrade apply
kubeadm upgrade node --kubeconfig /etc/kubernetes/admin.conf

## kubelet upgrade
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubelet
mv kubelet /usr/local/bin/kubelet
chmod +x /usr/local/bin/kubelet
systemctl restart kubelet.service

## kubectl upgrade
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubectl
mv kubectl /usr/local/bin/kubectl
chmod +x /usr/local/bin/kubectl

# ################### worker-node-02 ###################
## kubeadm
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubeadm
mv kubeadm /usr/local/bin/kubeadm
chmod +x /usr/local/bin/kubeadm

## drain node
kubectl drain worker-node-02 --ignore-daemonsets --delete-emptydir-data

## upgrade apply
kubeadm upgrade node --kubeconfig /etc/kubernetes/admin.conf

## kubelet upgrade
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubelet
mv kubelet /usr/local/bin/kubelet
chmod +x /usr/local/bin/kubelet
systemctl restart kubelet.service

## kubectl upgrade
curl -LO https://dl.k8s.io/release/v1.34.6/bin/linux/amd64/kubectl
mv kubectl /usr/local/bin/kubectl
chmod +x /usr/local/bin/kubectl