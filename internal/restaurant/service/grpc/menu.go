package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"go.uber.org/multierr"
	"time"
)

func (su *RestaurantUsecase) CreateMenu(mc *types.MenuCreate) error {
	var errs error
	var slp []string
	slp = concantenateProducts(mc.Salads, slp)
	slp = concantenateProducts(mc.Garnishes, slp)
	slp = concantenateProducts(mc.Meats, slp)
	slp = concantenateProducts(mc.Soups, slp)
	slp = concantenateProducts(mc.Drinks, slp)
	slp = concantenateProducts(mc.Desserts, slp)

	date := time.Now()

	menuId, err := su.repoR.CreateDate(mc.OnDate)
	if err != nil {
		multierr.Append(errs, err)
	}

	for _, v := range slp {
		id, pt, err := su.repoR.SelectProductByName(v, date)
		if err != nil {
			multierr.Append(errs, err)
		}

		err = su.repoR.CreateMenu(id, menuId, pt)
		if err != nil {
			multierr.Append(errs, err)
		}
	}

	return errs
}

func (su *RestaurantUsecase) GetMenu(t time.Time) (*types.Menu, error) {
	m, err := su.repoR.GetMenu(t)
	return m, err
}
