package resources

import (
	"github.com/setest/pet-shop/api-gateway/internal/errors"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type R struct {
	DB *sqlx.DB
}

type Config struct {
	//DiagPort    int    `envconfig:"DIAG_PORT" default:"8081" required:"true"`
	RESTAPIPort int    `envconfig:"PORT" default:"8080" required:"true"`
	DBURL       string `mapstructure:"DBURL" required:"true"`
}

func New(c Config) (*R, error) {

	db, err := sqlx.Connect("postgres", c.DBURL)
	if err != nil {
		return nil, errors.ErrConfigIncorrectWithMsg(err)
	}

	return &R{
		DB: db,
	}, nil
}

func (r *R) Release() error {
	return r.DB.Close()
}
