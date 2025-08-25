IMAGE_NAME=app-service
IMAGE_TAG=latest

build_image:
	docker build -t $(IMAGE_NAME):$(IMAGE_TAG) ./app-service

load_image:
	 minikube image load $(IMAGE_NAME):$(IMAGE_TAG)
install_metrics:
	minikube addons enable metrics-server

install_app:
	helm install app-service ./helm/app/

uninstall_app:
	helm uninstall app-service

install_hpa:
	#kubectl apply -f ./helm/hpa/hpa.yaml && kubectl get hpa
	kubectl delete -f ./helm/hpa/hpa.yaml

install_sm:
	kubectl apply -f ./helm/monitoring/sm/service_monitor.yaml


install_prometheus_adapter:
	helm install prometheus-adapter prometheus-community/prometheus-adapter -f ./helm/monitoring/adapter/values.yaml

install_hpa_rpc:
	kubectl apply -f ./helm/monitoring/adapter/hpa.yaml


#kube-prometheus-stack has been installed. Check its status by running:
#  kubectl --namespace default get pods -l "release=prometheus-operator"
#
#Get Grafana 'admin' user password by running:
#
#  kubectl --namespace default get secrets prometheus-operator-grafana -o jsonpath="{.data.admin-password}" | base64 -d ; echo
#
#Access Grafana local instance:
#
#  export POD_NAME=$(kubectl --namespace default get pod -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=prometheus-operator" -oname)
#  kubectl --namespace default port-forward $POD_NAME 3000
