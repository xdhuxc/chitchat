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


