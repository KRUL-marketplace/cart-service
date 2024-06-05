LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.53.3

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@v0.10.1
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.15.2
	GOBIN=$(LOCAL_BIN) go install github.com/rakyll/statik@v0.1.7

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go get -u github.com/gosimple/slug
	go get -u github.com/joho/godotenv
	go get -u github.com/jackc/pgx/v4
	go get -u github.com/pkg/errors
	go get -u github.com/georgysavva/scany/pgxscan
	go get -u google.golang.org/grpc
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/runtime
	go get -u github.com/grpc-ecosystem/grpc-gateway/v2/utilities
	go get -u github.com/rs/cors
	go get -u github.com/Masterminds/squirrel
	go get -u github.com/lib/pq
	go get -u github.com/google/uuid
	go get -u github.com/rakyll/statik
	go get -u github.com/KRUL-marketplace/product-catalog-service

generate-cart-service-api:
	mkdir -p pkg
	mkdir -p pkg/cart-service
	protoc --proto_path api --proto_path vendor.protogen \
	--go_out=pkg/cart-service --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/cart-service --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	--validate_out lang=go:pkg/cart-service --validate_opt=paths=source_relative \
	--plugin=protoc-gen-validate=bin/protoc-gen-validate \
	--grpc-gateway_out=pkg/cart-service --grpc-gateway_opt=paths=source_relative \
	--plugin=protoc-gen-grpc-gateway=bin/protoc-gen-grpc-gateway \
	--openapiv2_out=allow_merge=true,merge_file_name=api:pkg/swagger \
	--plugin=protoc-gen-openapiv2=bin/protoc-gen-openapiv2 \
	api/cart.proto

generate:
	mkdir -p pkg/swagger
	make generate-cart-service-api
	$(LOCAL_BIN)/statik -src=pkg/swagger/ -include='*.css,*.html,*.js,*.json,*.png'

vendor-proto:
	@if [ ! -d vendor.protogen/validate ]; then \
		   mkdir -p vendor.protogen/validate &&\
		   git clone https://github.com/envoyproxy/protoc-gen-validate vendor.protogen/protoc-gen-validate &&\
		   mv vendor.protogen/protoc-gen-validate/validate/*.proto vendor.protogen/validate &&\
		   rm -rf vendor.protogen/protoc-gen-validate ;\
	fi
	@if [ ! -d vendor.protogen/google ]; then \
		   git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
		   mkdir -p  vendor.protogen/google/ &&\
		   mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
		   rm -rf vendor.protogen/googleapis ;\
	fi
	@if [ ! -d vendor.protogen/protoc-gen-openapiv2 ]; then \
		   mkdir -p vendor.protogen/protoc-gen-openapiv2/options &&\
		   git clone https://github.com/grpc-ecosystem/grpc-gateway vendor.protogen/openapiv2 &&\
		   mv vendor.protogen/openapiv2/protoc-gen-openapiv2/options/*.proto vendor.protogen/protoc-gen-openapiv2/options &&\
		   rm -rf vendor.protogen/openapiv2 ;\
	fi
	@if [ ! -d vendor.protogen/grpc-graphql-gateway ]; then \
		   mkdir -p vendor.protogen/graphql-gateway &&\
		   git clone https://github.com/ysugimoto/grpc-graphql-gateway.git vendor.protogen/grpc-graphql-gateway &&\
		   mv vendor.protogen/grpc-graphql-gateway/include/graphql/*.proto vendor.protogen/graphql-gateway &&\
		   rm -rf vendor.protogen/grpc-graphql-gateway ;\
	fi
