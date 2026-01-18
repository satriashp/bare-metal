# Create namespace
kubectl create namespace "argocd"

# Install core component
kubectl apply -n "argocd" -f ./GitOps/argocd/core-install-v3.2.5.yaml

#  Generate key secret
KEY=$(head -c 32 /dev/urandom | base64)
kubectl patch secret argocd-secret -n "argocd" --type='merge' -p="{\"stringData\":{\"server.secretkey\":\"$(head -c 32 /dev/urandom | base64)\"}}"


# Generate cloudflare api token secret
kubectl create secret generic cloudflare-api-token-secret \
  --from-literal=api-token="${CLOUDFLARE_API_TOKEN}" \
  -n cert-manager \
  --dry-run=client -o yaml | kubectl apply -f -

# Generate admin password for pi.hole
kubectl create secret --namespace external-dns generic pihole-password \
    --from-literal EXTERNAL_DNS_PIHOLE_PASSWORD=${EXTERNAL_DNS_PIHOLE_PASSWORD}