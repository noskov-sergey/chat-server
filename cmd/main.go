package main

import (
	"context"
	"flag"
	"github.com/jackc/pgx/v5/pgxpool"
	chatApi "github.com/noskov-sergey/chat-server/internal/api/chats"
	"github.com/noskov-sergey/chat-server/internal/config"
	chatRepository "github.com/noskov-sergey/chat-server/internal/repository/chats"
	messageRepository "github.com/noskov-sergey/chat-server/internal/repository/messages"
	userRepository "github.com/noskov-sergey/chat-server/internal/repository/user"
	chatUsecase "github.com/noskov-sergey/chat-server/internal/usecase/chats"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedChatV1Server
	pool *pgxpool.Pool
}

func main() {
	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	cfg, err := config.NewGPRCConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("failed to get pg config: %v", err)
	}

	lis, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	messageRep := messageRepository.NewMessagesRepository(pool)
	userRep := userRepository.NewUserRepository(pool)
	chatRep := chatRepository.NewChatRepository(pool)

	usecase := chatUsecase.New(chatRep, userRep, messageRep)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, chatApi.New(usecase))

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
