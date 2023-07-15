package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	"log"
)

func (cr *CustomerRepo) CreateOffice(office *types.Office) error {
	fmt.Println("Repo CreateOffice was invoked")

	gOffice := &gormdb.Office{
		Name:    office.Name,
		Address: office.Address,
	}

	err := cr.DB.Create(gOffice).Error
	if err != nil {
		return fmt.Errorf("can't create office in gorm: %v\n", err)
	}
	return err
}

func (cr *CustomerRepo) GetOfficeList() ([]*types.Office, error) {
	log.Printf("GetOfficeList Repository was invoked")
	officesList := &[]gormdb.Office{}
	err := cr.DB.Find(officesList).Error
	if err != nil {
		return nil, fmt.Errorf("can't create office in gorm: %v\n", err)
	}
	tOffs := []*types.Office{}

	for _, v := range *officesList {
		tOff := &types.Office{}
		tOff.Uuid = v.ID
		tOff.Name = v.Name
		tOff.Address = v.Address
		tOff.CreatedAt = v.CreatedAt
		tOffs = append(tOffs, tOff)
	}
	return tOffs, nil
}

// CreateOffice without db
//m := officeMap
//var slO []*pb.Office
//file, err := os.OpenFile("office.json", os.O_WRONLY, 0644)
//if err != nil {
//fmt.Println(err)
//}
//defer file.Close()
//
//other := &pb.Office{}
//err = copier.Copy(other, office)
//if err != nil {
//return fmt.Errorf("cannot copy product data: %w", err)
//}
//
//m[other.Uuid] = other
//for _, v := range m {
//slO = append(slO, v)
//}
//j, err := json.Marshal(slO)
//if err != nil {
//return err
//}
//file.Write(j)

//GetOfficeList without db
//data, err := os.OpenFile("office.json", os.O_RDONLY, 0644)
//if err != nil {
//log.Fatal("can't read office.json", err)
//}
//defer data.Close()
//
//m, err := io.ReadAll(data)
//if err != nil {
//log.Println("Can't read data from Office.json: ", err)
//return sl, err
//}
//
//err = json.Unmarshal(m, &sl)
//if err != nil {
//log.Fatal("cannot unmarshall data office.json", err)
//}

// with pq driver
//
//if err := cr.DB.QueryRow(
//"INSERT INTO office (name, adress) VALUES ($1, $2) RETURNING uuid",
//office.Name,
//office.Address,
//).Scan(&office.Uuid); err != nil {
//return err
//}

//var sl []*types.Office
//offices, err := cr.DB.Query("SELECT * FROM office")
//if err != nil {
//return nil, fmt.Errorf("Error to get OfficeList from db: %s", err)
//
//}
//defer offices.Close()
//
//for offices.Next() {
//tp := &types.Office{}
//if err = offices.Scan(&tp.Uuid, &tp.Name, &tp.Address, &tp.CreatedAt); err != nil {
//return nil, fmt.Errorf("trouble with row.Next Officelist: %v\n", err)
//}
//
//sl = append(sl, tp)
//
//}
