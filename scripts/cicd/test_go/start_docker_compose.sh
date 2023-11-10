set -e

# Install yq - used for parsing docker-compose.yaml
if ! command -v yq &> /dev/null; then
    wget https://github.com/mikefarah/yq/releases/download/${VERSION}/${BINARY} -O /usr/bin/yq && chmod +x /usr/bin/yq
fi

# Boot up services
docker compose up -d