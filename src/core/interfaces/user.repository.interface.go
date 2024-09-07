package interfaces

import (
	"etnasys.com/clientService/protogen/go/user"
)

type IUserRepository interface {
	Create(user *user.CreateUserRequest) (*user.User, error)
	Get(id *user.IdRequest) (*user.User, error)
	Search(param *user.SearchUserRequest) (*user.PaginatedUserResponse, error)
	Update(user *user.User) (*user.User, error)
	Delete(id *user.IdRequest) (*user.DeleteUserResponse, error)
}
