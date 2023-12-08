package serviceR

import (
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"time"
)

func (su *RestaurantUsecase) CreateMenu(mc *types.MenuCreate) error {

	var slp []string
	slp = concantenateProducts(mc.Salads, slp)
	slp = concantenateProducts(mc.Garnishes, slp)
	slp = concantenateProducts(mc.Meats, slp)
	slp = concantenateProducts(mc.Soups, slp)
	slp = concantenateProducts(mc.Drinks, slp)
	slp = concantenateProducts(mc.Desserts, slp)

	date := time.Now()

	menuId, err := su.repoR.CreateMenuDate(mc)
	if err != nil {
		return err
	}
	for _, v := range slp {
		id, pt, err := su.repoR.SelectProductByName(v, date)
		if err != nil {
			return err
		}
		err = su.repoR.CreateMenu(menuId, id, pt)
		if err != nil {
			return err
		}
	}

	return err
}

func (su *RestaurantUsecase) GetMenu(t time.Time) (*types.Menu, error) {
	menu := &types.Menu{}
	var slp []*types.Product

	id, err := su.repoR.GetMenuId(t)
	if err != nil {
		return nil, err
	}
	pid, err := su.repoR.GetProductId(id)
	if err != nil {
		return nil, err
	}
	for _, v := range pid {
		p, err := su.repoR.GetMenu(v)
		if err != nil {
			return nil, err
		}
		slp = append(slp, p)
	}

	for _, v := range slp {
		switch {
		case v.Type == 1:
			menu.Salads = append(menu.Salads, v)
		case v.Type == 2:
			menu.Garnishes = append(menu.Garnishes, v)
		case v.Type == 3:
			menu.Meats = append(menu.Meats, v)
		case v.Type == 4:
			menu.Soups = append(menu.Soups, v)
		case v.Type == 5:
			menu.Drinks = append(menu.Drinks, v)
		default:
			menu.Desserts = append(menu.Desserts, v)
		}
	}

	menu.OpenAt, menu.ClosedAt, menu.CreatedAt, err = su.repoR.GetTimes(id)
	if err != nil {
		return nil, err
	}
	menu.Uuid = id
	return menu, err
}
