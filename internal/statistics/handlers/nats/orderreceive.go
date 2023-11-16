package natsstatistics

import (
	"context"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (n *NatsSubStat) OrderReceive(order *types.OrderRequest) error {
	log.Println("OrderReceiveStat was invoked")

	resProducts, err := n.ClientProduct.GetProductList(context.Background(), &restaurant.GetProductListRequest{})
	if err != nil {
		code := codes.Internal
		return status.Errorf(code, "GetProductList calling by Stat.Profit went down witn error, cannot save products in db: %v\n", err)
	}

	products, err := convertToProduct(resProducts.Result)
	if err != nil {
		code := codes.Internal
		return status.Errorf(code, "convertToProduct went down witn error, cannot save products in db: %v\n", err)
	}

	err = n.Jm.DataSaveService(order, products)
	if err != nil {
		log.Println(err)
	}

	return err
}
