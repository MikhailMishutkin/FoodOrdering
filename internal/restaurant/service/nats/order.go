package natsrestservice

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
)

func (r *RestNATSService) DataSaveService(
	order *types.OrderRequest,
	offices []*types.Office,
	users []*types.User,
) error {
	err := r.nrr.SaveOfficeList(offices)
	if err != nil {
		return err
	}
	err = r.nrr.SaveUserList(users)
	if err != nil {
		return err
	}

	var slOI []*types.OrderItem

	slOI = append(slOI, order.Salads...)
	slOI = append(slOI, order.Garnishes...)
	slOI = append(slOI, order.Meats...)
	slOI = append(slOI, order.Soups...)
	slOI = append(slOI, order.Drinks...)
	slOI = append(slOI, order.Desserts...)

	userUuid := order.UserUuid

	err = r.nrr.ReceiveOrder(slOI, userUuid)
	return err
}
