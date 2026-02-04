# Bare-Metal Kubernetes Bootstrap
*kubeadm · Ansible · GitOps*

This repository documents a **repeatable, automated, and production-oriented approach to bootstrapping Kubernetes on bare-metal infrastructure** using **kubeadm**, **Ansible**, and a **GitOps-ready design**.

The objective is not simply to “get a cluster running”.

The objective is to **expose and understand the Kubernetes fundamentals** that are typically hidden by managed services such as EKS, GKE, or AKS — and to build a solid baseline suitable for real operations, not demos.

---

## Why Bare Metal?

Managed Kubernetes platforms abstract away many critical details, including:

- Node identity and IP stability
- Certificate wiring and SAN configuration
- etcd topology and control-plane assumptions
- Low-level networking behavior

Those abstractions are convenient — but they also hide failure modes.

This project **intentionally removes those safety nets**.

By bootstrapping Kubernetes directly on bare metal, you are forced to think like a **platform or cloud provider**, not just a cluster consumer. That perspective directly improves how you design, debug, and operate Kubernetes in *any* environment.

---

## What This Project Focuses On

### 🧱 Infrastructure & Bootstrap
- Kubernetes bootstrapped using **kubeadm**
- Explicit `kubeadm` configuration (no hidden defaults)
- Static IP-based node identity to avoid certificate and control-plane instability
- Clear separation between control-plane and worker responsibilities

### 🤖 Automation First
- Node preparation fully automated using **Ansible**
- No manual SSH-driven setup
- Idempotent, repeatable provisioning
- Designed to be rebuilt from scratch without configuration drift

### 🔁 GitOps-Ready by Design
- Cluster bootstrap clearly separated from workload delivery
- **ArgoCD scaffolding** included for post-bootstrap operations
- Focus on **Day-2 operations**, not just Day-0 installation
- Declarative management as the source of truth

---

## Repository Structure

```text
.
├── ansible/            # Node preparation and OS-level configuration
├── kubeadm/            # kubeadm init/join configuration and templates
├── gitops/
│   └── argocd/         # GitOps scaffolding for declarative workload delivery
├── inventory/          # Ansible inventory and host definitions
└── README.md
```