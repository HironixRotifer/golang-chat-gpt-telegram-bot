package reportserver

import (
	"log"
	"net"

	"github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/logger"

	types "github.com/HironixRotifer/golang-chat-gpt-telegram-bot/pkg/bottypes"
)

type MessageSender interface {
	SendReportToUser(dt []types.UserDataReportRecord, userID int64, reportKey string) error
}

// server is used to implement UserReportsReciverServer.
type server struct {
	api.UnimplementedUserReportsReciverServer
	msgModel MessageSender
}

// StartReportServer запуск сервиса, слушающего сервис формирования отчетов.
func StartReportServer(msgModel MessageSender) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Fatal("failed to listen", "err", err)
	}
	s := grpc.NewServer()
	api.RegisterUserReportsReciverServer(s, &server{msgModel: msgModel})
	log.Printf("server listening at %v", lis.Addr())
	go func() {
		if err := s.Serve(lis); err != nil {
			logger.Fatal("failed to serve", "err", err)
		}
	}()
}
