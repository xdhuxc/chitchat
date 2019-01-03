package models

import (
	"time"
)

type Thread struct {
	ID        int
	UUID      string
	Topic     string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID        int
	UUID      string
	Body      string
	UserId    string
	ThreadId  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (thread *Thread) UpdatedAtDate() string {
	return thread.UpdatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

func (post *Post) UpdatedAtDate() string {
	return post.UpdatedAt.Format("Jan 2, 2006 at 3:04pm")
}

/**
获取某个主题下面文章的数量
*/
func (thread *Thread) Count() int {
	var count int
	rows, err := DB.Query("select count(*) from posts where thread_id = $1", thread.ID)
	if err != nil {
		return count
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return count
		}
	}
	rows.Close()
	return count
}

/**
获取某个话题下面的所有文章
*/
func (thread *Thread) Posts() ([]Post, error) {
	var posts []Post
	rows, err := DB.Query("select id, uuid, body, user_id, thread_id, created_at, updated_at from posts where thread_id = $1", thread.ID)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		if err = rows.Scan(&post.ID, &post.UUID, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt, &post.UpdatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	rows.Close()
	return posts, nil
}

/**
创建一个话题
*/
func (user *User) CreateThread(topic string) (Thread, error) {
	var thread Thread
	sql := "insert into threads (uuid, topic, user_id, created_at, updated_at) values ($1, $2, $3, $4, $5) return id, uuid, topic, user_id, created_at, updated_at"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return thread, err
	}

	defer statement.Close()

	if err = statement.QueryRow(createUUID(), topic, user.ID, time.Now(), time.Now()).Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt, &thread.UpdatedAt); err != nil {
		return thread, err
	}
	return thread, nil
}

/**
为某个话题增加一条记录
*/
func (user *User) CreatePost(thread Thread, body string) (Post, error) {
	var post Post
	sql := "insert into posts (uuid, body, user_id, thread_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6) returning id, uuid, body, user_id, thread_id, created_at, updated_at"
	statement, err := DB.Prepare(sql)
	if err != nil {
		return post, err
	}
	defer statement.Close()

	if err = statement.QueryRow(createUUID(), body, user.ID, thread.ID, time.Now()).Scan(&post.ID, &post.UUID, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt, &post.UpdatedAt); err != nil {
		return post, err
	}
	return post, err
}

/**
获取所有的话题并返回之
*/
func Threads() ([]Thread, error) {
	var threads []Thread
	rows, err := DB.Query("select id, uuid, topic, user_id, created_at, updated_at from threads order by created desc")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var thread Thread
		if err = rows.Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt, &thread.UpdatedAt); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	rows.Close()
	return threads, nil
}

/**
Get a thread by it's UUID
*/
func FindThreadByUUID(uuid string) (Thread, error) {
	var thread Thread
	if err := DB.QueryRow("select id, uuid, topic, user_id, created_at, updated_at from threads where uuid = $1", uuid).
		Scan(&thread.ID, &thread.UUID, &thread.Topic, &thread.UserID, &thread.CreatedAt, &thread.UpdatedAt); err != nil {
		return thread, err
	}
	return thread, nil
}

/**
Get the user who started this thread
*/
func (thread *Thread) User() (User, error) {
	var user User
	if err := DB.QueryRow("select id, uuid, name, email, created_at, updated_at from users where id = $1", thread.UserID).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

/**
Get the user who wrote the post
*/
func (post *Post) User() (User, error) {
	var user User
	if err := DB.QueryRow("select id, uuid, name, email, created_at, updated_at from users where id = $1", post.UserId).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}
	return user, nil
}

func DeleteAllThreads() error {
	db := DB
	defer db.Close()
	sql := "delete from threads"
	_, err := db.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
