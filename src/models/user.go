package models

import "time"

type User struct {
	ID        int
	UUID      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Session struct {
	ID        int
	UUID      string
	Email     string
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

/**
	为已经存在的用户创建 Session
 */
func (user *User) CreateSession() (Session, error) {
	var session Session
	sql := "insert into sessions (uuid, email, user_id, created_at) values ($1, $2, $3, $4) returning id, uuid, email, user_id, created_at"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return session, err
	}
	defer statement.Close()

	err = statement.QueryRow(createUUID(), user.Email, user.ID, time.Now()).Scan(&session.ID, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		return session, err
	}
	return session, nil
}

/**
	从已经存在的用户中获取 Session
 */
func (user *User) Session() (Session, error) {
	var session Session

	err := DB.QueryRow("select id, uuid, email, user_id, created_at, updated_at from sessions where user_id = $1", user.ID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt, &session.UpdatedAt)
	if err != nil {
		return session, err
	}

	return session, nil
}

/**
	检查数据库中的 session 是否合法
 */
func (session *Session) Check() (bool, error) {
	var valid bool
	err := DB.QueryRow("select id, uuid, email, user_id, created_at, updated_at from sessions where uuid = $1", session.UUID).
		Scan(&session.ID, &session.UUID, &session.Email, &session.UserId, &session.CreatedAt, &session.UpdatedAt)
	if err != nil {
		valid = false
	}

	if session.ID != 0 {
		valid = true
	}

	return valid, err
}

/**
	从数据库中删除 session
 */
func (session *Session) DeleteByUUID() error {
	sql := "delete from sessions where uuid = $1"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(session.UUID)
	if err != nil {
		return err
	}
	return nil
}

/**
	从 session 中获取用户
 */
func (session *Session) User() (User, error) {
	var user User
	err := DB.QueryRow("select id, uuid, name, email, created_at, updated_at from users where id = $1", session.ID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

/**
	删除数据库中所有的 session
 */
func DeleteAllSessions() error {
	sql := "delete from sessions"
	_, err := DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}

/**
	创建一个 User，并保存到数据库中
 */
func (user *User) Create() error {
	sql := "insert into users (uuid, name, email, password, created_at, updated_at) values ($1, $2, $3, $4, $5) returning id, uuid, created_at"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	err = statement.QueryRow(createUUID(), user.Name, user.Email, Encrypt(user.Password), time.Now(), time.Now()).
		Scan(&user.ID, &user.UUID, &user.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

/**
	从数据库中删除 user
 */
func (user *User) Delete() error {
	sql := "delete from users where id = $1"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.ID); err != nil {
		return err
	}
	return nil
}

/**
	更新数据库中的用户信息
 */
func (user *User) Update() error {
	sql := "update users set name = $2, email = $3 where id = $1"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return err
	}
	defer statement.Close()
	if _, err := statement.Exec(user.ID, user.Name, user.Email); err != nil {
		return err
	}
	return nil
}

/**
	删除数据库中所有的用户信息
 */
func DeleteAllUsers() error {
	sql := "delete from users"
	if _, err := DB.Exec(sql); err != nil {
		return err
	}
	return nil
}

/**
	获取数据库中所有的用户信息
 */
func Users() ([]User, error) {
	rows, err := DB.Query("select id, uuid, name, email, password, created_at, updated_at from users")
	if err != nil {
		return nil, err
	}
	var users []User
	for rows.Next() {
		var user User
		if err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	rows.Close()
	return users, nil
}

/**
	使用邮箱查询用户信息
 */
func FindUserByEmail(email string) (User, error) {
	var user User
	sql := "select id, uuid, name, email, password, created_at, updated_at from users where email = $1"
	if err := DB.QueryRow(sql, email).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

/**
	使用邮箱查询用户信息
 */
func FindUserByUUID(uuid string) (User, error) {
	var user User
	sql := "select id, uuid, name, email, password, created_at, updated_at from users where uuid = $1"
	if err := DB.QueryRow(sql, uuid).Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}


