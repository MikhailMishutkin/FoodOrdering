package cusrepository

import (
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/internal/types"
	"log"
)

func (cr *CustomerRepo) CreateUser(user *types.User) error {
	log.Println("CreateUser repo was invoked")
	var ofName string
	err := cr.DB.QueryRow("SELECT FROM office WHERE uuid = $1", user.OfficeUuid).Scan(&ofName)
	if err := cr.DB.QueryRow(
		"INSERT INTO user (name, office_uuid, office_name) VALUES ($1, $2, $3) RETURNING uuid",
		user.Name,
		user.OfficeUuid,
		ofName,
	).Scan(&user.Uuid); err != nil {
		return err
	}

	return err
}

func (cr *CustomerRepo) GetUserList(id int) ([]*types.User, error) {
	log.Printf("GetOfficeList Repository was invoked")
	var slU []*types.User

	users, err := cr.DB.Query("SELECT * FROM user WHERE office_uuid= $1", id)
	if err != nil {
		return nil, fmt.Errorf("Error to get UserList from db: %s", err)

	}
	defer users.Close()

	for users.Next() {
		tp := &types.User{}
		if err = users.Scan(&tp.Uuid, &tp.Name, id, &tp.OfficeName, &tp.CreatedAt); err != nil {
			return nil, fmt.Errorf("trouble with row.Next Officelist: %v\n", err)
		}

		slU = append(slU, tp)

	}

	return slU, err
}

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
