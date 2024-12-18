package _struct

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go-ddd/pkg/config"
)

type EtlApp struct {
	Db        *pgxpool.Pool
	Config    *config.Schema
}