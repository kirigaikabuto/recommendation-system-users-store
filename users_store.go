package recommendation_system_users_store

type UsersStore interface {
	Create(user *User) (*User, error)
	Update(user *UserUpdate) (*User, error)
	Delete(id string) error
	Get(id string) (*User, error)
	List() ([]User, error)
	GetByUsernameAndPassword(username, password string) (*User, error)
}
