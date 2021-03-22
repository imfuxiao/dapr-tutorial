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

