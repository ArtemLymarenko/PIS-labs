package entity

type User struct {
	Id        string
	Email     string
	Password  string
	FirstName string
	LastName  string
}

func NewUser(id, email, password, firstname, lastname string) User {
	return User{
		Id:        id,
		Email:     email,
		Password:  password,
		FirstName: firstname,
		LastName:  lastname,
	}
}

func (u *User) Validate() error {
	return nil
}
