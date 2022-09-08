package repositories

import (
	"context"
	"database/sql"
	"simpleProfile/model/domain"
)

type UserRepository interface {
	Register(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile
	Login(ctx context.Context, tx *sql.Tx, profile domain.Profile) (domain.Profile, error)
	Update(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile
	FindById(ctx context.Context, tx *sql.Tx, profileId int) (domain.Profile, error)
	Logout(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile
}
