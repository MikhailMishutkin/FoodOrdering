package serviceR

import (
	repository "github.com/MikhailMishutkin/FoodOrdering/internal/restaurant/repository/postgres"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"go.uber.org/multierr"
	"time"
)

func (su *RestaurantUsecase) GetOrderList() (
	[]*types.OrderItem,
	[]*types.OrderByOffice,
	error,
) {
	var errs error

	t := repository.DateConv(time.Now())
	date := t.AddDate(0, 0, 1)

	//total items in orders
	ti, err := su.repoR.GetTotalOrders(date)
	if err != nil {
		errs = multierr.Append(errs, err)
	}

	//orders by offices
	offices, err := su.repoR.GetOfficesList()
	if err != nil {
		errs = multierr.Append(errs, err)
	}
	var ordersByOffice []*types.OrderByOffice
	for _, v := range offices {
		ibo, err := su.repoR.GetOrdersByOffice(date, v.OfficeUuid)
		if err != nil {
			errs = multierr.Append(errs, err)
		}
		v.Result = ibo
		ordersByOffice = append(ordersByOffice, v)
	}

	return ti, ordersByOffice, errs
}
