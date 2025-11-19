BUILD_TAG=$(shell keg image tag)
BUILD_GOARCH := $(shell go env GOARCH)
BUILD_GOOS := $(shell go env GOOS)
#BASE_IMAGE=$(IMAGE_REGISTRY)/monorepo:$(BUILD_TAG)
BASE_IMAGE=ubuntu:latest

build-tag:
	echo "$(BUILD_TAG)"

k8s.config.%:
	$(eval NAMESPACE:= $*)
	@keg k8s gen config --namespace=$(NAMESPACE)

generate:
	wire ./cmd/...

%.gen:
	$(eval SERVICE:= $*)
	go generate ./cmd/$(SERVICE)/main.go
	go generate ./internal/app/$(SERVICE)/service/graphql/generate.go

%.build:
	$(eval SERVICE:= $*)
	@echo "build: $(SERVICE):$(GIT_VERSION) platform: $(BUILD_GOOS)/$(BUILD_GOARCH)"
	#CGO_ENABLED=1 CC=x86_64-unknown-linux-gnu-gcc GOOS=$(BUILD_GOOS) GOARCH=$(BUILD_GOARCH) go build -ldflags "-X main.Version=$(GIT_VERSION)" -o ./bin/$(SERVICE) ./cmd/$(SERVICE)/
	CGO_ENABLED=0 GOOS=$(BUILD_GOOS) GOARCH=$(BUILD_GOARCH) go build -ldflags "-X main.Version=$(GIT_VERSION)" -o ./bin/$(SERVICE) ./cmd/$(SERVICE)/

%.build-windows:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).build BUILD_GOOS=windows BUILD_GOARCH=amd64

%.build-linux:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).build BUILD_GOOS=linux BUILD_GOARCH=amd64

%.build-mac:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).build BUILD_GOOS=darwin BUILD_GOARCH=arm64

%.image:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).build-linux
	docker build --platform=linux/amd64 --build-arg IMAGE=$(BASE_IMAGE) -t $(SERVICE):$(BUILD_TAG) -f ./deploy/build/$(SERVICE)/Dockerfile .

%.publish:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).image
	docker tag $(SERVICE):$(BUILD_TAG) $(IMAGE_REGISTRY)/$(SERVICE):$(BUILD_TAG)
	docker push $(IMAGE_REGISTRY)/$(SERVICE):$(BUILD_TAG)

%.publish.only:
	$(eval SERVICE:= $*)
	docker tag $(SERVICE):$(BUILD_TAG) $(IMAGE_REGISTRY)/$(SERVICE):$(BUILD_TAG)
	docker push $(IMAGE_REGISTRY)/$(SERVICE):$(BUILD_TAG)

%.deploy:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).publish
	@keg k8s deployment update-image -n $(SERVICE) -t $(BUILD_TAG)

%.deploy.gen:
	$(eval SERVICE:= $*)
	@$(MAKE) $(SERVICE).gen
	@$(MAKE) $(SERVICE).publish
	@keg k8s deployment update-image -n $(SERVICE) -t $(BUILD_TAG)

%.deploy.only:
	$(eval SERVICE:= $*)
	@keg k8s deployment update-image -n $(SERVICE) -t $(BUILD_TAG)

%.deploy_all_in_one:
	$(eval SERVICE:= $*)
	docker build --build-arg SERVICE=$(SERVICE) --build-arg IMAGE=$(BASE_IMAGE) VERSION=$(GIT_VERSION) -t $(SERVICE):$(BUILD_TAG) -f ./deploy/build/Dockerfile .
	docker tag $(SERVICE):$(BUILD_TAG) $(IMAGE_REGISTRY)/$(SERVICE):$(BUILD_TAG)
	docker push $(IMAGE_REGISTRY)/$(SERVICE):$(BUILD_TAG)

# 推送公共镜像
common_image.publish:
	docker build -t monorepo:latest -t monorepo:$(BUILD_TAG) -f ./deploy/build/common/Dockerfile .
	docker tag monorepo:latest $(IMAGE_REGISTRY)/monorepo:latest
	docker tag monorepo:$(BUILD_TAG) $(IMAGE_REGISTRY)/monorepo:$(BUILD_TAG)
	docker push $(IMAGE_REGISTRY)/monorepo:latest
	docker push $(IMAGE_REGISTRY)/monorepo:$(BUILD_TAG)
