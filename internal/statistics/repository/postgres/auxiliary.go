package statrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/proto/pkg/restaurant"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strconv"
	"time"
)

func convertToTProduct(pr []*restaurant.Product) (pt []*types.Product, err error) {
	for _, v := range pr {
		uuid, err := strconv.Atoi(v.Uuid)
		if err != nil {
			return nil, fmt.Errorf("can't convert product uuid in Profit stathandler: %v\n", err)
		}
		product := &types.Product{
			Uuid:      uuid,
			Name:      v.Name,
			Type:      int(v.Type.Number()),
			Price:     v.Price,
			CreatedAt: timeAssert(v.CreatedAt),
		}
		pt = append(pt, product)
	}
	return pt, err
}

func timeAssert(ts *timestamppb.Timestamp) time.Time {
	return time.Unix(ts.Seconds, int64(ts.Nanos))
}
