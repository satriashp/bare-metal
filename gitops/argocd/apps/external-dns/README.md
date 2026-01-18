# Generate admin password for pi.hole
kubectl create secret --namespace external-dns generic pihole-password \
    --from-literal EXTERNAL_DNS_PIHOLE_PASSWORD=${EXTERNAL_DNS_PIHOLE_PASSWORD}