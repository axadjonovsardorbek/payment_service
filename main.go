package main

import (
	"log"
	"net"
	"path/filepath"
	"runtime"

	cf "payment-service/config"
	"payment-service/config/logger"

	pb "payment-service/genproto/payment"
	service "payment-service/service"
	"payment-service/storage/postgres"

	"google.golang.org/grpc"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

func main() {
	config := cf.Load()
	logger := logger.NewLogger(basepath, config.LOG_PATH) // Don't forget to change your log path
	em := cf.NewErrorManager(logger)
	db, err := postgres.NewPostgresStorage(config, logger)
	em.CheckErr(err)
	defer db.Db.Close()

	listener, err := net.Listen("tcp", config.PAYMENT_PORT)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterPaymentServiceServer(s, service.NewPaymentService(db))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
