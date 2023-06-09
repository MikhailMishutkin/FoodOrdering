package cusrepository

import (
	"encoding/json"
	"fmt"
	pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"
	"github.com/jinzhu/copier"
	"io"
	"log"
	"os"
)

func (cr *CustomerRepo) CreateOffice(office *pb.Office) error {
	fmt.Println("Repo CreateOffice was invoked")
	m := officeMap
	var slO []*pb.Office
	file, err := os.OpenFile("office.json", os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	other := &pb.Office{}
	err = copier.Copy(other, office)
	if err != nil {
		return fmt.Errorf("cannot copy product data: %w", err)
	}

	m[other.Uuid] = other

	for _, v := range m {
		slO = append(slO, v)
	}
	j, err := json.Marshal(slO)
	if err != nil {
		return err
	}
	file.Write(j)

	return err
}

func (cr *CustomerRepo) GetOfficeList() ([]*pb.Office, error) {
	log.Printf("GetOfficeList Repository was invoked")

	//var officeInst *pb.Office
	var sl []*pb.Office
	data, err := os.OpenFile("office.json", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal("can't read office.json", err)
	}
	defer data.Close()

	m, err := io.ReadAll(data)
	if err != nil {
		log.Println("Can't read data from Office.json: ", err)
		return sl, err
	}

	//office := &pb.Office{}

	err = json.Unmarshal(m, &sl)
	if err != nil {
		log.Fatal("cannot unmarshall data office.json", err)
	}

	//for _, v := range sl{
	//	officeInst = v
	//	sl = append(sl, v)
	//}

	return sl, err
}
