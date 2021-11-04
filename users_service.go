package recommendation_system_users_store

import (
	"github.com/google/uuid"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type UserService interface {
	CreateUser(cmd *CreateUserCommand) (*User, error)
	UpdateUser(cmd *UpdateUserCommand) (*User, error)
	DeleteUser(cmd *DeleteUserCommand) error
	GetUser(cmd *GetUserCommand) (*User, error)
	ListUser(cmd *ListUserCommand) ([]User, error)
	GetUserByUsernameAndPassword(cmd *GetUserByUsernameAndPassword) (*User, error)
}

type userService struct {
	userStore UsersStore
}

func NewUserService(uStore UsersStore) UserService {
	return &userService{userStore: uStore}
}

func (u *userService) CreateUser(cmd *CreateUserCommand) (*User, error) {
	user := &User{
		Id:        uuid.New().String(),
		Username:  cmd.Username,
		FirstName: cmd.FirstName,
		LastName:  cmd.LastName,
		Email:     cmd.Email,
		Age:       cmd.Age,
		Country:   cmd.Country,
	}
	hashedPassword, err := setdata_common.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword
	return u.userStore.Create(user)
}

func (u *userService) UpdateUser(cmd *UpdateUserCommand) (*User, error) {
	if cmd.Id == "" {
		return nil, ErrUserIdNotProvided
	}
	oldUser, err := u.GetUser(&GetUserCommand{
		Id: cmd.Id,
	})
	if err != nil {
		return nil, err
	}
	userUpdate := &UserUpdate{Id: cmd.Id}
	if cmd.Password != "" {
		hashedPassword, err := setdata_common.HashPassword(cmd.Password)
		if err != nil {
			return nil, err
		}
		if oldUser.Password != hashedPassword {
			userUpdate.Password = &hashedPassword
		}
	}
	if cmd.FirstName != "" && cmd.FirstName != oldUser.FirstName {
		userUpdate.FirstName = &cmd.FirstName
	}
	if cmd.LastName != "" && cmd.LastName != oldUser.LastName {
		userUpdate.LastName = &cmd.LastName
	}
	if cmd.Username != "" && cmd.Username != oldUser.Username {
		userUpdate.Username = &cmd.Username
	}
	if cmd.Email != "" && cmd.Email != oldUser.Email {
		userUpdate.Email = &cmd.Email
	}
	if cmd.Age != 0 && cmd.Age != oldUser.Age {
		userUpdate.Username = &cmd.Username
	}
	return u.userStore.Update(userUpdate)
}

func (u *userService) DeleteUser(cmd *DeleteUserCommand) error {
	return u.userStore.Delete(cmd.Id)
}

func (u *userService) GetUser(cmd *GetUserCommand) (*User, error) {
	return u.userStore.Get(cmd.Id)
}

func (u *userService) ListUser(cmd *ListUserCommand) ([]User, error) {
	return u.userStore.List()
}

func (u *userService) GetUserByUsernameAndPassword(cmd *GetUserByUsernameAndPassword) (*User, error) {
	return u.userStore.GetByUsernameAndPassword(cmd.Username, cmd.Password)
}
