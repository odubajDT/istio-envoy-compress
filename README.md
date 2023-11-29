## Steps

1. install istio on cluster, easiest way via helm https://istio.io/latest/docs/setup/install/helm/
2. build the go service with ko `KO_DOCKER_REPO=docker.io/odubajdt ko build . --bare --tags test` -> use this image in manifests
3. create namespace and label it for istio envoy proxy injection `kubectl create ns istio-test && kubectl label namespace istio-test istio-injection=enabled`
4. adapt deployment image and deploy manifests `kubectl apply -f manifest.yaml -n istio-test`
5. Deploy debug pod `kubectl apply -f pod.yaml -n istio-test`
6. connect to the pod and execute `curl -H "Accept-Encoding: gzip" poc-istio-service.istio-test.svc.cluster.local:8080/long -o -` -> data should not be compressed
7. Deploy envoyfilter `kubectl apply -f envoyfilter.yaml -n istio-test`
8. execute the `curl` command from the pod again -> you should get compressed data

Note: commented version of envoyfilter with zstd does cause errors (unknown implementation) in the proxy sidecar logs, compression with zstd si not working, but is should (istio-1.20 is using envoy 1.28 which has zstd support), until now I do not know why it isn't working
