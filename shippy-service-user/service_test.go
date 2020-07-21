package main

import (
	"context"
	"strings"
	"testing"

	pb "github.com/haxorbit/shippy/shippy-service-user/proto/user"
)

var (
	user = &pb.User{
		Id:    "user001",
		Email: "john@foo.bar",
	}
)

type MockRepo struct{}

func (repo *MockRepo) GetAll(ctx context.Context) ([]*User, error) {
	var users []*User
	return users, nil
}

func (repo *MockRepo) Get(ctx context.Context, id string) (*User, error) {
	var user *User
	return user, nil
}

func (repo *MockRepo) Create(ctx context.Context, user *User) error {
	return nil
}

func (repo *MockRepo) GetByEmail(ctx context.Context, email string) (*User, error) {
	var user *User
	return user, nil
}

func newInstance() authable {
	repo := &MockRepo{}
	return &TokenService{repo}
}

func TestCanCreateToken(t *testing.T) {
	srv := newInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}

	if token == "" {
		t.Fail()
	}

	if len(strings.Split(token, ".")) != 3 {
		t.Fail()
	}
}

func TestCanDecodeToken(t *testing.T) {
	srv := newInstance()
	token, err := srv.Encode(user)
	if err != nil {
		t.Fail()
	}
	claims, err := srv.Decode(token)
	if err != nil {
		t.Logf("Decode failed err: %v", err)
		t.Fail()
	}
	if claims.User == nil {
		t.Fail()
	}
	if claims.User.Email != "john@foo.bar" {
		t.Fail()
	}
}

func TestThrowsErrorIfTokenInvalid(t *testing.T) {
	srv := newInstance()
	_, err := srv.Decode("blah123.another321.random098")
	if err == nil {
		t.Fail()
	}
}
