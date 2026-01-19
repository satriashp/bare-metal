# Bare-Metal Kubernetes Bootstrap (kubeadm + Ansible + GitOps)

This repository contains a **documented and automated approach to bootstrapping Kubernetes on bare-metal infrastructure** using `kubeadm`, **Ansible**, and a **GitOps-ready structure**.

The goal of this project is not just to “get a cluster running”, but to build a **repeatable, predictable, and production-oriented baseline** that exposes the fundamentals usually hidden by managed Kubernetes services (EKS/GKE/AKS).

---

## Why Bare Metal?

Managed Kubernetes abstracts away many critical details:
- Node identity and IP stability
- Certificate wiring and SANs
- etcd endpoints and control plane assumptions
- Networking fundamentals

This project intentionally removes those abstractions.

By building Kubernetes on bare metal, you are forced to think like a **cloud provider**, not just a cluster user. That understanding directly translates to operating Kubernetes more reliably in any environment.

---

## What This Repository Covers

### ✅ Infrastructure & Bootstrap
- Kubernetes cluster bootstrapped using **kubeadm**
- Explicit `kubeadm` configuration files (no hidden defaults)
- Static IP-based node identity to avoid certificate and control plane instability

### ✅ Automation
- Node preparation fully automated with **Ansible**
- No manual SSH-driven setup
- Reproducible, idempotent provisioning

### ✅ GitOps-Ready Design
- Cluster bootstrap is clearly separated from workload delivery
- **ArgoCD scaffolding** included to manage workloads declaratively after bootstrap
- Designed for Day-2 operations, not just Day-0 installs

---

## Repository Structure

```text
.
├── ansible/            # Node preparation and system configuration
├── kubeadm/            # kubeadm init/join configuration and templates
├── gitops/
│   └── argocd/         # GitOps scaffolding for declarative workload management
├── inventory/          # Ansible inventory and host definitions
└── README.md
```

---
# Assumptions & Requirements

* Bare-metal or VM-based infrastructure

* Ubuntu Server (tested on 24.04 LTS)

* Static IP addresses assigned to all nodes

* SSH access to all nodes

* Basic familiarity with:

  * Kubernetes

  * Ansible

  * Linux networking

* Pi-hole installed on your network (or *other DNS server tools*)

# High-Level Workflow

1. Prepare nodes

    * Configure static IPs

    * Disable swap

    * Install container runtime

    * Install kubeadm, kubelet, kubectl

    * All handled via Ansible

2. Bootstrap control plane

    * Initialize Kubernetes using explicit kubeadm config (`./kubeadm` directory)

3. Join worker nodes

    * Workers join using generated tokens and CA cert hashes

    * No manual configuration drift

4. Enable GitOps

    * ArgoCD scaffolding prepared for declarative workload management

# What This Is (and Isn’t)
## This is:

* A reproducible Kubernetes bootstrap baseline

* A learning and reference platform for understanding Kubernetes internals

* A foundation for production-grade bare-metal or on-prem clusters

## This is not:

* A one-command installer

* A managed Kubernetes replacement

* A cloud-provider-specific solution

# Documentation

A human-readable walkthrough of the bootstrap process is available here:

📘 Notion Docs
https://thankful-baryonyx-8a2.notion.site/Bootstrap-kubernetes-with-kubeadm-2af05e33ac74802ea2e9c98567e4f2de

# Final Note

If you have only ever used managed Kubernetes, building a cluster this way—even once—will permanently change how you understand Kubernetes.

That perspective is the real value of this project.