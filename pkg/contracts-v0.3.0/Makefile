
API_DIR =

.PHONY: gen-proto
gen-proto:
	mkdir -p pkg/contracts/restaurant
	protoc \
		-I api/mediasoft-internship/final-task/contracts/restaurant \
		-I third_party/googleapis \
		-I third_party/envoyproxy/protoc-gen-validate \
		--go_out=./pkg/contracts/restaurant --go_opt=paths=source_relative \
        --go-grpc_out=./pkg/contracts/restaurant  --go-grpc_opt=paths=source_relative \
        --validate_out="lang=go:./pkg/contracts/restaurant" --validate_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/contracts/restaurant --grpc-gateway_opt=paths=source_relative \
        --openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/contracts/restaurant \
		./api/mediasoft-internship/final-task/contracts/restaurant/*.proto

	mkdir -p pkg/contracts/customer
	protoc \
		-I api/mediasoft-internship/final-task/contracts/customer \
		-I third_party/googleapis \
		-I third_party/envoyproxy/protoc-gen-validate \
		--go_out=./pkg/contracts/customer --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/contracts/customer  --go-grpc_opt=paths=source_relative \
		--validate_out="lang=go:./pkg/contracts/customer" --validate_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/contracts/customer --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/contracts/customer \
		./api/mediasoft-internship/final-task/contracts/customer/*.proto

	mkdir -p pkg/contracts/statistics
	protoc \
		-I api/mediasoft-internship/final-task/contracts/statistics \
		-I third_party/googleapis \
		-I third_party/envoyproxy/protoc-gen-validate \
		--go_out=./pkg/contracts/statistics --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/contracts/statistics  --go-grpc_opt=paths=source_relative \
		--validate_out="lang=go:./pkg/contracts/statistics" --validate_opt=paths=source_relative \
		--grpc-gateway_out=./pkg/contracts/statistics --grpc-gateway_opt=paths=source_relative \
		--openapiv2_out=use_go_templates=true,json_names_for_fields=false,allow_merge=true,merge_file_name=api:./pkg/contracts/statistics \
		./api/mediasoft-internship/final-task/contracts/statistics/*.proto
