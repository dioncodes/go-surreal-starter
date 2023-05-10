package user

import (
	"github.com/dioncodes/go-surreal-starter/pkg/env"
	"github.com/surrealdb/surrealdb.go"
)

type User struct {
	Id        string `json:"id,omitempty"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetUsers(env *env.Env) ([]User, error) {
	data, err := env.DB.Select("user")

	if err != nil {
		return []User{}, err
	}

	// Unmarshal data
	users := []User{}

	err = surrealdb.Unmarshal(data, &users)
	return users, err
}

func GetUserById(env *env.Env, id string) (User, error) {
	data, err := env.DB.Select("user:" + id)

	if err != nil {
		return User{}, err
	}

	// Unmarshal data
	user := new(User)
	err = surrealdb.Unmarshal(data, &user)

	return *user, err
}

func (user *User) Save(env *env.Env) error {
	if user.Id == "" {
		// Insert user
		data, err := env.DB.Create("user", user)
		if err != nil {
			return err
		}

		// Unmarshal data
		createdUser := make([]User, 1)
		err = surrealdb.Unmarshal(data, &createdUser)

		if err != nil {
			return err
		}

		user.Id = createdUser[0].Id
		return nil
	}

	// Update user
	_, err := env.DB.Update(user.Id, user)
	return err
}

func (user *User) Delete(env *env.Env) error {
	_, err := env.DB.Delete(user.Id)
	return err
}
