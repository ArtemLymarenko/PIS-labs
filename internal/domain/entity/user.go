package entity

type User struct {
	Id       string
	Email    string
	Password string
}

func NewUser(id, email, password string) User {
	return User{
		Id:       id,
		Email:    email,
		Password: password,
	}
}

func (u *User) Validate() error {
	return nil
}
