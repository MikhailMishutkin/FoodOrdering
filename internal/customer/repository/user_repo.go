package cusrepository

import (
	"encoding/json"
	"fmt"
	"github.com/MikhailMishutkin/FoodOrdering/microservices/gen"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
)

func (cr *CustomerRepo) CreateUser(user *pb.User) error {
	m := officeMap

	for k, v := range m {
		user.Uuid = gen.RandomID()
		user.OfficeUuid = k
		user.OfficeName = v.Name
	}

	user.CreatedAt = timestamppb.Now()

	file, err := os.OpenFile("user.json", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	j, err := json.Marshal(user)
	if err != nil {
		return err
	}
	file.Write(j)

	return err
}

func (cr *CustomerRepo) GetUserList(in string) (*pb.GetUserListResponse, error) {
	data, err := os.OpenFile("user.json", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("can't read user.json", err)
	}
	defer data.Close()

	m, err := io.ReadAll(data)
	if err != nil {
		log.Println("Can't read data from user.json: ", err)
		return nil, err
	}
	var u *pb.User
	var slU []*pb.User
	err = json.Unmarshal(m, &u)
	if err != nil {
		log.Fatal("cannot unmarshall data office.json", err)
	}
	slU = append(slU, u)
	resp := new(pb.GetUserListResponse)
	resp.Result = slU

	return resp, err
}
