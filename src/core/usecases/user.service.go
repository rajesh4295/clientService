package usecases

import (
	"context"

	"etnasys.com/clientService/protogen/go/user"
	"etnasys.com/clientService/src/core/interfaces"
)

type UserService struct {
	user.UnimplementedUserServiceServer
	userRepo interfaces.IUserRepository
}

func NewUserService(userRepo interfaces.IUserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (service *UserService) Create(ctx context.Context, req *user.CreateUserRequest) (*user.User, error) {
	return service.userRepo.Create(req)
}

func (service *UserService) Get(ctx context.Context, req *user.IdRequest) (*user.User, error) {
	return service.userRepo.Get(req)
}
func (service *UserService) Search(ctx context.Context, req *user.SearchUserRequest) (*user.PaginatedUserResponse, error) {
	return service.userRepo.Search(req)
}

func (service *UserService) Update(ctx context.Context, req *user.User) (*user.User, error) {
	return service.userRepo.Update(req)
}

func (service *UserService) Delete(ctx context.Context, req *user.IdRequest) (*user.DeleteUserResponse, error) {
	return service.userRepo.Delete(req)
}
