# Generate manifest file
helm upgrade --install --namespace monitoring --create-namespace monitoring oci://ghcr.io/prometheus-community/charts/kube-prometheus-stack --values=value-override.yaml --version=77.13.0

helm template \
  monitoring \
  oci://ghcr.io/prometheus-community/charts/kube-prometheus-stack \
  --namespace monitoring \
  --create-namespace \
  --version 77.13.0 \
  --values values.override.yaml \
  --include-crds \
  > monitoring-manifest.yaml
