package Models


type User struct {
	Id       int64
	Username string `xorm:"varchar(256) not null"`
	Password string `xorm:"varchar(256) not null"`
}

func NewUser(username , password string) *User  {
	newUser := new(User)
	newUser.Username = username
	newUser.Password = password
	return newUser
}