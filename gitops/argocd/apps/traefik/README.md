# Generate manifest file from helm template command
helm template \
  traefik \
  traefik/traefik \
  --namespace traefik \
  --create-namespace \
  --version 38.0.2 \
  --values values.override.yaml \
  --include-crds \
  > traefik-manifest.yaml
