package services

import (
	"github.com/TeplyyMaksim/bookstore_users-api/domain/users"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/crypto_utils"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/date_utils"
	"github.com/TeplyyMaksim/bookstore_users-api/utils/errors_utils"
)

type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors_utils.HttpError)
	GetUser(int) (*users.User, *errors_utils.HttpError)
	UpdateUser (users.User, bool) (*users.User, *errors_utils.HttpError)
	DeleteUser(int) *errors_utils.HttpError
	SearchUser(string) (users.Users, *errors_utils.HttpError)
	LogInUser(users.LoginRequest) (*users.User, *errors_utils.HttpError)
}
var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct {}

func (*usersService) CreateUser(user users.User) (*users.User, *errors_utils.HttpError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowDBFormat()
	user.Password = crypto_utils.GetMd5(user.Password)
	user.Status = users.StatusActive
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (*usersService) GetUser(userId int) (*users.User, *errors_utils.HttpError) {
	result := users.User{ Id: userId }

	result.DateCreated = date_utils.GetNowDBFormat()
	if err := result.Get(); err != nil {
		return nil, err
	}

	return &result, nil
}

func (*usersService) UpdateUser (user users.User, isPartial bool) (*users.User, *errors_utils.HttpError) {
	current, err := UsersService.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
		if user.Status != "" {
			current.Status = user.Status
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
		current.Status = user.Status
	}


	if err = current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func (*usersService) DeleteUser(userId int) *errors_utils.HttpError {
	user := &users.User{ Id: userId }

	return user.Delete()
}

func (*usersService) SearchUser(status string) (users.Users, *errors_utils.HttpError) {
	return  users.FindByStatus(status)
}

func (*usersService) LogInUser(request users.LoginRequest) (*users.User, *errors_utils.HttpError) {
	dao := &users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password),
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return dao, nil
}