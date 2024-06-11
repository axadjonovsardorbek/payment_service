package storage

import (
	p "payment-service/genproto/payment"
)

type PaymentI interface {
    Create(*p.PaymentCreateReq)(*p.PaymentGetByIdResp, error)
	GetById(*p.PaymentGetByIdReq)(*p.PaymentGetByIdResp, error)
    Update(*p.PaymentUpdateReq)(*p.PaymentGetByIdResp, error)
}

