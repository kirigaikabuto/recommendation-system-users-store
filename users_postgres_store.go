package recommendation_system_users_store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

var marketplaceAppRepoQueries = []string{
	`CREATE TABLE IF NOT EXISTS users(
		id TEXT,
		username TEXT,
		password TEXT,
		email TEXT,
		first_name TEXT,
		last_name TEXT,
		age integer,
		country TEXT,
		PRIMARY KEY(id)
	);`,
}

type usersStore struct {
	db *sql.DB
}

func NewPostgresUsersStore(cfg PostgresConfig) (UsersStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range marketplaceAppRepoQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &usersStore{db: db}
	return store, nil
}

func (u *usersStore) Create(user *User) (*User, error) {
	result, err := u.db.Exec("INSERT INTO users (id, username, password, email, first_name, last_name, age, country) "+
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		user.Id, user.Username, user.Password, user.Email, user.FirstName, user.LastName, user.Age, user.Country,
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateUserUnknown
	}
	return user, nil
}

func (u *usersStore) Update(user *UserUpdate) (*User, error) {
	q := "update users set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if user.FirstName != nil {
		cnt++
		parts = append(parts, "first_name = $"+strconv.Itoa(cnt))
		values = append(values, user.FirstName)
	}
	if user.LastName != nil {
		cnt++
		parts = append(parts, "last_name = $"+strconv.Itoa(cnt))
		values = append(values, user.LastName)
	}
	if user.Password != nil {
		cnt++
		parts = append(parts, "password = $"+strconv.Itoa(cnt))
		values = append(values, user.Password)
	}
	if user.Email != nil {
		cnt++
		parts = append(parts, "email = $"+strconv.Itoa(cnt))
		values = append(values, user.Email)
	}
	if user.Username != nil {
		cnt++
		parts = append(parts, "username = $"+strconv.Itoa(cnt))
		values = append(values, user.Username)
	}
	if user.Age != nil {
		cnt++
		parts = append(parts, "age = $"+strconv.Itoa(cnt))
		values = append(values, user.Age)
	}
	if user.Country != nil {
		cnt++
		parts = append(parts, "country = $"+strconv.Itoa(cnt))
		values = append(values, user.Country)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	q = q + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, user.Id)
	result, err := u.db.Exec(q, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrUserNotFound
	}
	return u.Get(user.Id)

}

func (u *usersStore) Delete(id string) error {
	result, err := u.db.Exec("delete from users where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrUserNotFound
	}
	return nil
}

func (u *usersStore) Get(id string) (*User, error) {
	user := &User{}
	err := u.db.QueryRow("select id, username, password, email, first_name, last_name, age, country "+
		"from users where id = $1 limit 1", id).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Age, &user.Country)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *usersStore) List() ([]User, error) {
	users := []User{}
	var values []interface{}
	q := "select id, username, password, email, first_name, last_name, age, country from users"
	//cnt := 1
	rows, err := u.db.Query(q, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Age, &user.Country)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *usersStore) GetByUsernameAndPassword(username, password string) (*User, error) {
	user := &User{}
	err := u.db.QueryRow("select id, username, password, email, first_name, last_name, age, country "+
		"from users where username = $1 limit 1", &username).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.FirstName, &user.LastName, &user.Age, &user.Country)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	} else if err != nil {
		return nil, err
	}
	//fmt.Println("old user", user)
	//fmt.Println(password)
	//fmt.Println([]byte(password))
	//fmt.Println([]byte(user.Password))
	//compare := setdata_common.CheckPasswordHash(password, user.Password)
	//fmt.Println(compare)
	//if !compare {
	//	return nil, ErrUserPasswordNotCorrect
	//}
	return user, nil
}
