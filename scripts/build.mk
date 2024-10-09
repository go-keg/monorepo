
admin.build: admin.build
admin.image: admin.image
admin.publish: admin.publish
admin.deploy: admin.deploy

api.build: api.build
api.image: api.image
api.publish: api.publish
api.deploy: api.deploy

.PHONY: generate
generate:
	wire ./cmd/...

build.all: generate
	go build -ldflags "-X main.Version=$(GIT_VERSION)" -o ./bin/ ./cmd/...

test:
	go test -v ./internal/... -cover

%.gen:
	$(eval SERVICE:= $*)
	go generate ./cmd/$(GIT_VERSION)/main.go

%.build:
	$(eval SERVICE:= $*)
	@echo "build: $(SERVICE):$(GIT_VERSION)-$(GIT_BRANCH)"
	go env -w CGO_ENABLED=0 GOOS=linux GOARCH=amd64
	go build -ldflags "-X main.Version=$(GIT_VERSION)" -o ./bin/$(SERVICE) ./cmd/$(SERVICE)/

%.image:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).build
	docker build -t $(SERVICE):$(GIT_VERSION)-$(GIT_BRANCH) -f ./deploy/build/$(SERVICE)/Dockerfile .

%.publish:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).image
	@echo "publish $(SERVICE)"
	docker tag $(SERVICE):$(GIT_VERSION)-$(GIT_BRANCH) $(IMAGE_REGISTRY)/$(SERVICE):$(GIT_VERSION)-$(GIT_BRANCH)
	docker push $(IMAGE_REGISTRY)/$(SERVICE):$(GIT_VERSION)-$(GIT_BRANCH)

%.deploy:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).publish
	@echo "deploy $(SERVICE)"
	@echo "TODO deploy to k8s"