package postgres

import (
	"database/sql"
	"payment-service/config/logger"
	p "payment-service/genproto/payment"
)

type PaymentRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewPaymentRepo(db *sql.DB, logger *logger.Logger) *PaymentRepo {
	return &PaymentRepo{db: db, Logger: logger}
}

func (pr *PaymentRepo) Create(payment *p.PaymentCreateReq)(*p.PaymentGetByIdResp,error){
	return nil, nil
}

func (pr *PaymentRepo) GetById(id *p.PaymentGetByIdReq)(*p.PaymentGetByIdResp, error){
	return nil, nil
}

func (pr *PaymentRepo) Update(payment *p.PaymentUpdateReq)(*p.PaymentGetByIdResp, error){
	return nil, nil
}