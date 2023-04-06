.PHONY : run_http run_grpc pwd mod

run_http: pwd mod
	@rm -rf luwu_http_server
	@wire ./cmd/http_server/wire.go
	GO111MODULE=on go build -o luwu_http_server ./cmd/http_server
	@./luwu_http_server

run_grpc: pwd mod
	@rm -rf luwu_grpc_server
	GO111MODULE=on go build -o luwu_grpc_server ./cmd/grpc_server
	@./luwu_grpc_server

pwd:
	@pwd

mod:
	@go mod tidy
