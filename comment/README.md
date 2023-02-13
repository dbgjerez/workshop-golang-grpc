
# Applications

## Server

### Configuration
ENV_CONFIG_JSON_LOCATION
SERVICE_VERSION
SERVICE_NAME
SERVICE_BUILD_TIME

### Build
```zsh
SERVICE_NAME=comment-service
VERSION=0.3
SERVICE_BUILD_TIME=$(date '+%Y/%m/%d %H:%M:%S')
podman build \
    --no-cache \
    --build-arg version=$VERSION \
    --build-arg serviceName=$SERVICE_NAME \
    --build-arg buildTime=$SERVICE_BUILD_TIME \
    -t quay.io/dborrego/$SERVICE_NAME:$VERSION \
    -f Containerfile
```