.PHONY: init-helm
init-helm:
	helm repo add dapr https://dapr.github.io/helm-charts/
	helm repo update

.PHONY: minikube
minikube:
	helm install dapr dapr/dapr \
		--version=1.0.1 \
		--namespace dapr-system \
		--create-namespace \
		--values $${PWD}/deploy/helm/values.yaml \
		--wait

.PHONY: app-ping
app-ping:
	dapr run --app-id app-ping \
		-d $${PWD}/deploy/config \
		--dapr-http-port 3500 \
		--app-port 50001 \
		--log-level debug \
 		go run ./cmd/ping/main.go

.PHONY: app-pong
app-pong:
	dapr run --app-id app-pong \
		-d $${PWD}/deploy/config \
		--log-level debug \
 		go run ./cmd/pong/main.go

.PHONY: app-client
app-client:
	dapr run --app-id app-client \
		-d $${PWD}/deploy/config \
		--log-level debug \
 		go run ./cmd/client/main.go

APIROOTS=github.com/imfuxiao/dapr-tutorial/pkg/id-server/v1
current_dir := $(notdir $(patsubst %/,%,$(dir $(mkfile_path))))
.PHONY: protobuf
protobuf:
	go-to-protobuf \
	--only-idl=true \
	--apimachinery-packages="" \
	--proto-import="${current_dir}pkg/id-server" \
	--packages="${APIROOTS}" \
	--go-header-file "${current_dir}hack/boilerplate.go.txt" \
	--keep-gogoproto=true \
#	--output-base="${current_dir}"