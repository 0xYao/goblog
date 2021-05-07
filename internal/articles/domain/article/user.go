package article

import "errors"

type User struct {
	id        string
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

type NewUserInput struct {
	Id        string
	Avatar    string
	FirstName string
	LastName  string
}

func NewUser(in NewUserInput) (*User, error) {
	if in.Id == "" {
		return nil, errors.New("user id is empty")
	}

	return &User{
		id:        in.Id,
		avatar:    in.Avatar,
		firstName: in.FirstName,
		lastName:  in.LastName,
	}, nil
}
