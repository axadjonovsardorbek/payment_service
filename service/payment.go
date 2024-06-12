package service

import (
	"context"
	p "payment-service/genproto/payment"
	st "payment-service/storage/postgres"
)

type PaymentService struct {
	storage st.Storage
	p.UnimplementedPaymentServiceServer
}

func NewPaymentService(storage *st.Storage) *PaymentService {
	return &PaymentService{
		storage: *storage,
	}
}

func (s *PaymentService) Create(ctx context.Context, payment *p.PaymentCreateReq)(*p.PaymentGetByIdResp,error){
	
	resp, err := s.storage.PaymentS.Create(payment)

	if err != nil{
		return nil, err
	}

	return resp, nil
}

func (s *PaymentService) GetById(ctx context.Context, id *p.PaymentGetByIdReq)(*p.PaymentGetByIdResp, error){
	
	payment, err := s.storage.PaymentS.GetById(id)

	if err != nil{
		return nil, err
	}
	
	return payment, nil
}

func (s *PaymentService) Update(ctx context.Context, payment *p.PaymentUpdateReq)(*p.PaymentGetByIdResp, error){
	
	resp, err := s.storage.PaymentS.Update(payment)

	if err != nil{
		return nil, err
	}
	
	return resp, nil
}