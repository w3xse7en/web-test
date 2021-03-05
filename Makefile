gen:
	make genapi

genapi:
	rm -rf pkg/api/*.go

	protoc -I api/proto \
--go_out pkg/api --go_opt paths=source_relative \
--go-grpc_out pkg/api --go-grpc_opt paths=source_relative \
--grpc-gateway_out pkg/api \
--grpc-gateway_opt logtostderr=true \
--grpc-gateway_opt paths=source_relative \
--grpc-gateway_opt generate_unbound_methods=true \
snowflake.proto
