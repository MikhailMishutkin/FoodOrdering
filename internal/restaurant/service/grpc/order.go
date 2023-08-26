package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"time"
)

func (su *RestaurantUsecase) GetOrderList(
	o []*types.Office,
	u []*types.User,

) (
	[]*types.OrderItem,
	[]*types.OrderByOffice,
	error,
) {
	date := time.Now()
	t, tbo, err := su.repoR.GetOrderList(date, o, u)
	return t, tbo, err
}
