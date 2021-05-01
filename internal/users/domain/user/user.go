package user

import "errors"

var (
	UserNotFound      = errors.New("user is not found")
	UserAlreadyExists = errors.New("user already exists")
)

type User struct {
	id string
	// avatar can be empty
	avatar    string
	firstName string
	lastName  string
}

func (u *User) Id() string {
	return u.id
}

func (u *User) Avatar() string {
	return u.avatar
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) LastName() string {
	return u.lastName
}

type Factory struct{}

func NewFactory() (Factory, error) {
	return Factory{}, nil
}

type NewUserInput struct {
	Id        string
	Avatar    string
	FirstName string
	LastName  string
}

func (f Factory) NewUser(in *NewUserInput) (*User, error) {
	if in.Id == "" {
		return nil, errors.New("user id is empty")
	}

	if in.FirstName == "" {
		return nil, errors.New("user first name is empty")
	}

	if in.LastName == "" {
		return nil, errors.New("user last name is empty")
	}

	return &User{
		id:        in.Id,
		avatar:    in.Avatar,
		firstName: in.FirstName,
		lastName:  in.LastName,
	}, nil
}

type UpdateUserInput struct {
	Id        string
	Avatar    string
	FirstName string
	LastName  string
}

func (u *User) SetAvatar(avatar string) {
	u.avatar = avatar
}

func (u *User) SetFirstName(firstName string) error {
	if firstName == "" {
		return errors.New("first name is empty")
	}

	u.firstName = firstName
	return nil
}

func (u *User) SetLastName(lastName string) error {
	if lastName == "" {
		return errors.New("last name is empty")
	}

	u.lastName = lastName
	return nil
}
