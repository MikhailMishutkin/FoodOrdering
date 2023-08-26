package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"time"
)

func (su *RestaurantUsecase) CreateMenu(mc *types.MenuCreate) error {
	err := su.repoR.CreateMenu(mc)
	return err
}

func (su *RestaurantUsecase) GetMenu(t time.Time) (*types.Menu, error) {
	m, err := su.repoR.GetMenu(t)
	return m, err
}
