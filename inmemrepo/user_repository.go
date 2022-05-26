package inmemrepo

import (
	"fmt"
	"go-app-service-test/domain/model"
)

type UserRepository struct {
	users []model.User
}

func (ur UserRepository) FindByID(id model.UserID) (model.User, error) {
	for _, user := range ur.users {
		if user.ID() == id {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("user not found for user_id: %s", id)
}

func (ur UserRepository) FindByName(name model.UserName) (model.User, error) {
	for _, user := range ur.users {
		if user.Name() == name {
			return user, nil
		}
	}

	return model.User{}, fmt.Errorf("user not found for user_name: %s", name)
}

func (ur *UserRepository) Create(user model.User) (model.User, error) {
	ur.users = append(ur.users, user)

	return user, nil
}

func (ur *UserRepository) Update(user model.User) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (ur *UserRepository) Delete(user model.User) error {
	//TODO implement me
	panic("implement me")
}
