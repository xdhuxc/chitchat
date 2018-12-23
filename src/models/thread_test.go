package models

import "testing"

func ThreadDeleteAll() error {
	db := DB
	defer db.Close()
	sql := "delete from threads"
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

func TestCreatedThread(t *testing.T) {
	setup()
	if err := users[0].Create(); err != nil {
		t.Error(err, "Can not create user.")
	}

	thread, err := users[0].CreateThread("My First Thread")
	if err != nil {
		t.Error(err, "Can not create thread.")
	}

	if thread.UserID != users[0].ID {
		t.Error("User not linked with thread.")
	}
}
