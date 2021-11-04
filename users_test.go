package recommendation_system_users_store

import (
	"fmt"
	setdata_common "github.com/kirigaikabuto/setdata-common"
	"testing"
)

var (
	store          UsersStore
	service        UserService
	err            error
	Id             = "122222"
	commandHandler setdata_common.CommandHandler
	username       = "3333"
	password       = "secret"
)

//testService

func TestUserService_CreateUser(t *testing.T) {
	config := PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "setdatauser",
		Password: "123456789",
		Database: "setdata",
		Params:   "sslmode=disable",
	}
	store, err = NewPostgresUsersStore(config)
	if err != nil {
		t.Error(err)
		return
	}
	service = NewUserService(store)
	commandHandler = setdata_common.NewCommandHandler(service)
	response, err := commandHandler.ExecCommand(&CreateUserCommand{
		Username:  username,
		Password:  password,
		Email:     "2222",
		FirstName: "222",
		LastName:  "2222",
	})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(response)
	Id = response.(*User).Id
}

func TestUserService_UpdateUser(t *testing.T) {
	cmd := &UpdateUserCommand{Id: Id, LastName: "tleugazy", Email: "tleugazy98@gmail.com", FirstName: "122334444"}
	response, err := commandHandler.ExecCommand(cmd)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(response)
}

func TestUserService_GetUserByUsernameAndPassword(t *testing.T) {
	cmd := &GetUserByUsernameAndPassword{
		Username: username,
		Password: password,
	}
	response, err := commandHandler.ExecCommand(cmd)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(response)
}

//test store
//func TestUsersStore_Create(t *testing.T) {
//	config := PostgresConfig{
//		Host:     "localhost",
//		Port:     5432,
//		User:     "setdatauser",
//		Password: "123456789",
//		Database: "setdata",
//		Params:   "sslmode=disable",
//	}
//	store, err = NewPostgresUsersStore(config)
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	_, err = store.Get(Id)
//	if err == nil {
//		store.Delete(Id)
//	}
//	newUser, err := store.Create(&User{
//		Id:        Id,
//		Username:  "456",
//		Password:  "123",
//		Email:     "678",
//		LoginType: setdata_common.Email,
//		FirstName: "123",
//		LastName:  "423",
//	})
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(newUser)
//}

//func TestUsersStore_Update(t *testing.T) {
//	user := &User{
//		Id:        "",
//		Username:  "123",
//		Password:  "123",
//		Email:     "123",
//		LoginType: "123",
//		FirstName: "123",
//		LastName:  "123",
//	}
//	updatedUser, err := store.Update(&UserUpdate{
//		Id:        Id,
//		Username:  &user.Username,
//		Password:  &user.Username,
//		Email:     &user.Username,
//		LoginType: nil,
//		FirstName: &user.Username,
//		LastName:  &user.Username,
//	})
//	if err != nil {
//		store.Delete(Id)
//		t.Error(err)
//		return
//	}
//	fmt.Println(updatedUser)
//}
//
//func TestUsersStore_List(t *testing.T) {
//	users, err := store.List()
//	if err != nil {
//		t.Error(err)
//		return
//	}
//	fmt.Println(users)
//}
