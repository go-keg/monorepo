# go.proxy
go.proxy:
	go env -w GO111MODULE=on
	go env -w GOPROXY=https://goproxy.cn,direct

# install dependents
install.dependents:
	go install github.com/google/wire/cmd/wire@latest
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/go-keg/keg/cmd/keg@latest
	go install github.com/go-keg/keg/cmd/protoc-gen-go-keg-error@latest

# init env
init:
	$(MAKE) go.proxy
	$(MAKE) install.dependents
	go mod tidy
	go mod download
	cp .env.example .env
	cp deploy/components/.env.example deploy/components/.env
