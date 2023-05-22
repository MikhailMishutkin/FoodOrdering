gengw:
protoc -I/usr/local/include -I.    -I$GOPATH    -I$GOPATH/FoodOrdering/pkg/contracts-v0.2.1/third_party/googleapis    --grpc-gateway_out=logtostderr=true:.    pkg/contracts-v0.2.1/api/mediasoft-internship/final-task/contracts/restaurant/restaurant_product.proto

.PHONY: gengw