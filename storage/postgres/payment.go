package postgres

import (
	"database/sql"
	"payment-service/config/logger"
)

type PaymentRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewPaymentRepo(db *sql.DB, logger *logger.Logger) *PaymentRepo {
	return &PaymentRepo{db: db, Logger: logger}
}
