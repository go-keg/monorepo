k8s.config.%:
	$(eval NAMESPACE:= $*)
	@y-deploy k8s gen config --namespace=$(NAMESPACE)

k8s.config.qa: k8s.config.qa
k8s.config.live: k8s.config.live
k8s.config.local: k8s.config.local
