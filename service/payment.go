package service

import (
	r "payment-service/genproto/payment"
	st "payment-service/storage/postgres"
)

type PaymentService struct {
	storage st.Storage
	r.UnimplementedPaymentServiceServer
}

func NewPaymentService(storage *st.Storage) *PaymentService {
	return &PaymentService{
		storage: *storage,
	}
}