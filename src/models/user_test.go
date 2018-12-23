package models

import (
	"database/sql"
	"testing"
)

func TestUserCreate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	if users[0].ID == 0 {
		t.Error("No id or created_at in user")
	}

	u, err := FindUserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "User not created.")
	}

	if users[0].Email != u.Email {
		t.Error("User retrieved is not the same as the one created.")
	}
}

func TestUserDelete(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	if err := users[0].Delete(); err != nil {
		t.Error(err, "Can not delete user.")
	}

	_, err := FindUserByEmail(users[0].Email)
	if err != sql.ErrNoRows {
		t.Error(err, "User not deleted.")
	}
}

func TestUserUpdate(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	users[0].Name = "Random User"
	if err := users[0].Update(); err != nil {
		t.Error(err, "Can not update user.")
	}

	u, err := FindUserByEmail(users[0].Email)
	if err != nil {
		t.Error(err, "Can get user.")
	}
	if u.Name != "Random User" {
		t.Error(err, "User not updated.")
	}
}

func TestFindUserByUUID(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}
	u, err := FindUserByUUID(users[0].UUID)
	if err != nil {
		t.Error(err, "User not created.")
	}
	if users[0].Email != u.Email {
		t.Errorf("User retrieved is not the same as the one created.")
	}
}

func TestUsers(t *testing.T) {
	setup()
	for _, user := range users {
		if err := user.Create(); err != nil {
			t.Error(err, "Can not create user.")
		}
	}
	u, err := Users()
	if err != nil {
		t.Error(err, "Can not retrieve users.")
	}

	if len(u) != 2 {
		t.Error(err, "Wrong number of users retrieved.")
	}

	if u[0].Email != users[0].Email {
		t.Error(u[0], users[0], "Wrong user retrieved.")
	}
}

func TestUser_CreateSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}
	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Can not create session.")
	}
	if session.UserId != users[0].ID {
		t.Error("User not linked with session.")
	}
}

func TestGetSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Can not create session.")
	}

	s, err := users[0].Session()
	if err != nil {
		t.Error(err, "Can not get session.")
	}

	if s.ID == 0 {
		t.Error("No session retrieved.")
	}

	if s.ID != session.ID {
		t.Error("Different session retrieved.")
	}
}

func TestValidSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Can not create session.")
	}
	uuid := session.UUID
	s := Session{UUID: uuid}
	valid, err := s.Check()
	if err != nil {
		t.Error(err, "Can not check session.")
	}
	if valid != true {
		t.Error(err, "Session is not valid.")
	}
}

func TestInvalidSession(t *testing.T) {
	setup()
	s := Session{UUID: "123"}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is not valid but is validated.")
	}
	if valid == true {
		t.Error(err, "Session is valid.")
	}
}

func TestDeleteSession(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	session, err := users[0].CreateSession()
	if err != nil {
		t.Error(err, "Can not create session.")
	}

	err = session.DeleteByUUID()
	if err != nil {
		t.Error(err, "Can not delete session.")
	}

	s := Session{UUID: session.UUID}
	valid, err := s.Check()
	if err == nil {
		t.Error(err, "Session is valid even though deleted.")
	}
	if valid == true {
		t.Error(err, "Session is not deleted.")
	}
}
