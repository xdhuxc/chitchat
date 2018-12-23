package models

import (
	"database/sql"
	"log"
	"crypto/rand"
	"fmt"
	"crypto/sha1"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "dbname=chitchat sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return
}

/**
	create a random UUID with from RFC 4122
	adapted from http://github.com/nu7hatch/gouuid
 */
func createUUID() string {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalf("Can not generate UUID, %v", err)
	}

	u[8] = (u[8] | 0x40) & 0x7F

	u[6] = (u[6] & 0xF) | (0x4 << 4)

	return fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
}

/**
	使用 SHA-1 哈希 文本
 */
func Encrypt(plaintext string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
}

