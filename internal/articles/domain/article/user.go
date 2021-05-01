package article

type User struct {
	id string
	avatar string
	firstName string
	lastName string
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
