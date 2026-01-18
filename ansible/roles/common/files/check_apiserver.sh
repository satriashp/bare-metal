#!/bin/bash
# Simple health check for kube-apiserver
curl --silent --max-time 1 https://127.0.0.1:6443/healthz --insecure > /dev/null
