package models

var users = []User{
	{
		Name:     "Peter",
		Email:    "peter@gmail.com",
		Password: "Peter_pass",
	},
	{
		Name:     "John",
		Email:    "john@gmail.com",
		Password: "john123456",
	},
}

func setup() {
	DeleteAllThreads()
	DeleteAllSessions()
	DeleteAllUsers()

}
