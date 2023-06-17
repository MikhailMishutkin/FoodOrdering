package service

import pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"

func (cu *CustomerUsecase) CreateUser(user *pb.User) error {
	err := cu.repoC.CreateUser(user)
	return err
}

func (cu *CustomerUsecase) GetUserList(in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	res, err := cu.repoC.GetUserList(in.OfficeUuid)
	return res, err
}
