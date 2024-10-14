package data

type Patient struct {
	Id      int
	Name    string
	Surname string
	Age     int
	Gender  string
}

type User struct {
	Id       int
	UserName string `json:"username`
	Password string `json:"password"`
}
