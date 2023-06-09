package service

import pb "github.com/MikhailMishutkin/FoodOrdering/proto/pkg/customer"

func (r *CustomerUsecase) CreateUser(user *pb.User) error {
	err := r.repoC.CreateUser(user)
	return err
}

func (r *CustomerUsecase) GetUserList(in *pb.GetUserListRequest) (*pb.GetUserListResponse, error) {
	res, err := r.repoC.GetUserList(in.OfficeUuid)
	return res, err
}
