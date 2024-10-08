run.%:
	$(eval COMPONENT:= $*)
	docker compose -f ./deploy/_components/docker-compose.yaml up -d --remove-orphans $(COMPONENT)

start.%:
	$(eval COMPONENT:= $*)
	docker compose -f ./deploy/_components/docker-compose.yaml start $(COMPONENT)

stop.%:
	$(eval COMPONENT:= $*)
	docker compose -f ./deploy/_components/docker-compose.yaml stop $(COMPONENT)

restart.%:
	$(eval COMPONENT:= $*)
	docker compose -f ./deploy/_components/docker-compose.yaml restart $(COMPONENT)

recreate.%:
	$(eval COMPONENT:= $*)
	docker compose -f ./deploy/_components/docker-compose.yaml up -d --build --force-recreate --remove-orphans $(COMPONENT)

# start all components
start.all:
	@$(MAKE) run.nacos --no-print-directory
	@$(MAKE) run.kafka --no-print-directory
	@$(MAKE) run.jaeger --no-print-directory
	@$(MAKE) run.prom-node-exporter --no-print-directory
	@$(MAKE) run.prometheus --no-print-directory
	@$(MAKE) run.grafana --no-print-directory

# stop all components
stop.all:
	@$(MAKE) stop.nacos --no-print-directory
	@$(MAKE) stop.kafka --no-print-directory
	@$(MAKE) stop.jaeger --no-print-directory
	@$(MAKE) stop.prom-node-exporter --no-print-directory
	@$(MAKE) stop.prometheus --no-print-directory
	@$(MAKE) stop.grafana --no-print-directory

ps:
	docker compose -f ./deploy/_components/docker-compose.yaml ps

logs:
	docker compose -f ./deploy/_components/docker-compose.yaml logs

logs.%:
	$(eval COMPONENT:= $*)
	@sudo docker compose -f ./deploy/_components/docker-compose.yaml logs $(COMPONENT)

kafka.recreate:
	$(MAKE) recreate.zookeeper
	$(MAKE) recreate.kafka
	$(MAKE) recreate.kafka-ui
