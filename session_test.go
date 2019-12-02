package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

const (
	DefaultTimeout = 24 * 7 * time.Hour
)

type Session struct {
	SessionId string `json:"-"`
	Uid       int    `json:"uid"`
	Email     string `json:"email"`
	Expire    int64  `json:"expire"`
	Signature string `json:"signature"`
}

func TestGenerate(t *testing.T) {
	session, err := Generate(1, "mail@gmail.com", time.Now().Add(DefaultTimeout).Unix())
	if err != nil {
		t.Fatal(err)
	}

	b, _ := json.Marshal(session)
	t.Log(string(b))
	t.Log(session.SessionId)
}

func Generate(uid int, email string, expireTs int64) (session *Session, err error) {
	if expireTs <= time.Now().Unix() {
		expireTs = time.Now().Add(DefaultTimeout).Unix()
	}

	signature := newSignature(uid, email, expireTs)

	session = &Session{
		Uid:       uid,
		Email:     email,
		Expire:    expireTs,
		Signature: signature,
	}

	data, err := json.Marshal(&session)
	if err != nil {
		return nil, nil
	}

	session.SessionId = base64.URLEncoding.EncodeToString(data)
	return session, nil
}

func newSignature(uid int, email string, expireSeconds int64) (signature string) {
	secretKey := Sha256Hex([]byte(strconv.Itoa(uid)))
	signature = Sha256Hex([]byte(fmt.Sprintf("%d:%s:%d:%s", uid, email, expireSeconds, secretKey)))
	return
}

func Sha256Hex(data []byte) string {
	return hex.EncodeToString(Sha256(data))
}

func Sha256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}
