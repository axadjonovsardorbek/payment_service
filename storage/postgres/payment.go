package postgres

import (
	"database/sql"
	"payment-service/config/logger"
	p "payment-service/genproto/payment"

	"github.com/google/uuid"
)

type PaymentRepo struct {
	db     *sql.DB
	Logger *logger.Logger
}

func NewPaymentRepo(db *sql.DB, logger *logger.Logger) *PaymentRepo {
	return &PaymentRepo{db: db, Logger: logger}
}

func (pr *PaymentRepo) Create(payment *p.PaymentCreateReq)(*p.PaymentGetByIdResp,error){

	id := uuid.New().String()
	res := p.PaymentGetByIdResp{}

	query := `
	INSERT INTO transactions(
		id,
		reservation_id,
		amount,
		payment_method,
		payment_status
	) VALUES ($1, $2, $3, $4, $5)
		RETURNING
		id,
		reservation_id,
		amount,
		payment_method,
		payment_status
	`

	row := pr.db.QueryRow(query, id, payment.ReservationId, payment.Amount, payment.PaymentMethod, payment.PaymentStatus)

	err := row.Scan(
		&res.Id, 
		&res.ReservationId,
		&res.Amount,
		&res.PaymentMethod,
		&res.PaymentStatus,
	)

	if err != nil {
		pr.Logger.ERROR.Println("Error while creating transaction")
		return nil, err
	}

	pr.Logger.INFO.Println("Successfully created transaction")

	return &res, nil
}

func (pr *PaymentRepo) GetById(id *p.PaymentGetByIdReq)(*p.PaymentGetByIdResp, error){
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	// |---------------------------------------------------------------|
	return nil, nil
}

func (pr *PaymentRepo) Update(payment *p.PaymentUpdateReq)(*p.PaymentGetByIdResp, error){
	res := p.PaymentGetByIdResp{}

	query := `
	UPDATE transactions SET
		reservation_id=$1,
		amount=$2,
		payment_method=$3,
		payment_status=$4
	WHERE 
		id=$5
	AND 
		deleted_at = 0
	RETURNING
		id,
		reservation_id,
		amount,
		payment_method,
		payment_status
	`

	row := pr.db.QueryRow(query, payment.Payment.ReservationId, payment.Payment.Amount, payment.Payment.PaymentMethod, payment.Payment.PaymentStatus, payment.Id)

	err := row.Scan(
		&res.Id, 
		&res.ReservationId,
		&res.Amount,
		&res.PaymentMethod,
		&res.PaymentStatus,
	)

	if err != nil {
		pr.Logger.ERROR.Println("Error while updating transaction")
		return nil, err
	}

	pr.Logger.INFO.Println("Successfully updated transaction")
	
	return &res, nil
}