package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"github.com/MikhailMishutkin/FoodOrdering/pkg/gormdb"
	"log"
)

func (cr *CustomerRepo) CreateUser(user *types.User) error {
	log.Println("CreateUser repo was invoked")
	gUser := &gormdb.User{
		Name:     user.Name,
		OfficeID: user.OfficeUuid,
	}

	err := cr.DB.Create(gUser).Error
	if err != nil {
		return fmt.Errorf("can't create office in gorm: %v\n", err)
	}
	return err
	return nil
}

func (cr *CustomerRepo) GetUserList(id int) ([]*types.User, error) {
	log.Printf("GetOfficeList Repository was invoked")
	usersList := &[]gormdb.User{}
	err := cr.DB.Where("office_id = ?", id).Find(usersList).Error
	if err != nil {
		return nil, fmt.Errorf("can't create office in gorm: %v\n", err)
	}

	tUsers := []*types.User{}

	for _, v := range *usersList {
		gOff := &gormdb.Office{}
		err = cr.DB.Where("id = ?", id).First(gOff).Error
		if err != nil {
			return nil, fmt.Errorf("can't find officeid in gorm (GetUserList): %v\n", err)
		}
		tUser := &types.User{}
		tUser.Uuid = v.ID
		tUser.Name = v.Name
		tUser.OfficeUuid = v.OfficeID
		tUser.OfficeName = gOff.Name
		tUser.CreatedAt = v.CreatedAt
		tUsers = append(tUsers, tUser)
	}

	return tUsers, err
}

//j, _ := json.Marshal(tUsers)
//fmt.Println(string(j))

//CreateUser without db
//m := officeMap
//
//for k, v := range m {
//user.Uuid = gen.RandomID()
//user.OfficeUuid = k
//user.OfficeName = v.Name
//}
//
//user.CreatedAt = timestamppb.Now()
//
//file, err := os.OpenFile("user.json", os.O_WRONLY, 0644)
//if err != nil {
//fmt.Println(err)
//}
//defer file.Close()
//
//j, err := json.Marshal(user)
//if err != nil {
//return err
//}
//file.Write(j)

//GetListUser without db
//var u *pb.User
//var slU []*pb.User
//data, err := os.OpenFile("user.json", os.O_RDONLY, 0644)
//if err != nil {
//log.Fatal("can't read user.json", err)
//}
//defer data.Close()
//
//m, err := io.ReadAll(data)
//if err != nil {
//log.Println("Can't read data from user.json: ", err)
//return nil, err
//}
//
//err = json.Unmarshal(m, &u)
//if err != nil {
//log.Fatal("cannot unmarshall data office.json", err)
//}
//slU = append(slU, u)
//resp := new(pb.GetUserListResponse)
//resp.Result = slU

//with pq driver
//var ofName string
//err := cr.DB.QueryRow("SELECT FROM office WHERE uuid = $1", user.OfficeUuid).Scan(&ofName)
//if err := cr.DB.QueryRow(
//"INSERT INTO user (name, office_uuid, office_name) VALUES ($1, $2, $3) RETURNING uuid",
//user.Name,
//user.OfficeUuid,
//ofName,
//).Scan(&user.Uuid); err != nil {
//return err
//}
//
//var slU []*types.User
//
//users, err := cr.DB.Query("SELECT * FROM user WHERE office_uuid= $1", id)
//if err != nil {
//return nil, fmt.Errorf("Error to get UserList from db: %s", err)
//
//}
//defer users.Close()
//
//for users.Next() {
//tp := &types.User{}
//if err = users.Scan(&tp.Uuid, &tp.Name, id, &tp.OfficeName, &tp.CreatedAt); err != nil {
//return nil, fmt.Errorf("trouble with row.Next Officelist: %v\n", err)
//}
//
//slU = append(slU, tp)
//
//}
