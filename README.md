## Steps

1. install istio on cluster
2. build the go service with ko `KO_DOCKER_REPO=docker.io/odubajdt ko build . --bare --tags test`
3. create namespace and label it for istio envoy proxy injection `kubectl create ns istio-test && kubectl label namespace istio-test`
4. adapt deployment image and deploy manifests `kubectl apply -f manifest.yaml -n istio-test`
5. port forward the port
6. execute curl commands to the port 8080, use /short or /long