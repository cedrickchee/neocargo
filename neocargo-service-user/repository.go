// Interact with PostgreSQL database

package main

import (
	"context"

	pb "github.com/haxorbit/neocargo/neocargo-service-user/proto/user"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// Database models

// User model
type User struct {
	ID       string `sql:"id"`
	Name     string `sql:"name"`
	Email    string `sql:"email"`
	Company  string `sql:"company"`
	Password string `sql:"password"`
}

// Repository is the API for user repository.
type Repository interface {
	GetAll(ctx context.Context) ([]*User, error)
	Get(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) error
	GetByEmail(ctx context.Context, email string) (*User, error)
}

// PostgresRepository implementation
type PostgresRepository struct {
	db *sqlx.DB
}

// NewPostgresRepository creates a new PostgreSQL repository by wrapping sqlx DB
// driver.
func NewPostgresRepository(db *sqlx.DB) *PostgresRepository {
	return &PostgresRepository{db}
}

// Marshalling and unmarshalling functions
//
// Convert between the protobuf definition generated structs, and our internal
// database models.

// MarshalUserCollection is ...
func MarshalUserCollection(users []*pb.User) []*User {
	u := make([]*User, len(users))
	for _, val := range users {
		u = append(u, MarshalUser(val))
	}
	return u
}

// MarshalUser is ...
func MarshalUser(user *pb.User) *User {
	return &User{
		ID:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

// UnmarshalUserCollection is ...
func UnmarshalUserCollection(users []*User) []*pb.User {
	u := make([]*pb.User, len(users))
	for _, val := range users {
		u = append(u, UnmarshalUser(val))
	}
	return u
}

// UnmarshalUser is ...
func UnmarshalUser(user *User) *pb.User {
	return &pb.User{
		Id:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Company:  user.Company,
		Password: user.Password,
	}
}

// GetAll gets all user.
func (r *PostgresRepository) GetAll(ctx context.Context) ([]*User, error) {
	users := make([]*User, 0)
	if err := r.db.GetContext(ctx, users, "select * from users"); err != nil {
		return users, err
	}

	return users, nil
}

// Get gets a user.
func (r *PostgresRepository) Get(ctx context.Context, id string) (*User, error) {
	var user *User
	if err := r.db.GetContext(ctx, &user, "select * from users where id = $1", id); err != nil {
		return nil, err
	}

	return user, nil
}

// Create creates a user.
func (r *PostgresRepository) Create(ctx context.Context, user *User) error {
	user.ID = uuid.NewV4().String()
	query := "insert into users (id, name, email, company, password) values ($1, $2, $3, $4, $5)"
	_, err := r.db.ExecContext(ctx, query, user.ID, user.Name, user.Email, user.Company, user.Password)
	return err
}

// GetByEmail gets a user by email address.
func (r *PostgresRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	query := "select * from users where email = $1"
	var user *User
	if err := r.db.GetContext(ctx, &user, query, email); err != nil {
		return nil, err
	}
	return user, nil
}
