package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/noskov-sergey/chat-server/internal/closer"
	"github.com/noskov-sergey/chat-server/internal/config"
	"log"

	chatApi "github.com/noskov-sergey/chat-server/internal/api/chats"
	chatRepository "github.com/noskov-sergey/chat-server/internal/repository/chats"
	messageRepository "github.com/noskov-sergey/chat-server/internal/repository/messages"
	userRepository "github.com/noskov-sergey/chat-server/internal/repository/user"
	chatUsecase "github.com/noskov-sergey/chat-server/internal/usecase/chats"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool *pgxpool.Pool
	cRep   *chatRepository.Repository
	uRep   *userRepository.Repository
	mRep   *messageRepository.Repository

	chatUsecase *chatUsecase.UseCase

	chatImpl *chatApi.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %v", err)
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGPRCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %v", err)
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PgPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) ChatRepository(ctx context.Context) *chatRepository.Repository {
	if s.cRep == nil {
		s.cRep = chatRepository.NewChatRepository(s.PgPool(ctx))
	}

	return s.cRep
}

func (s *serviceProvider) UserRepository(ctx context.Context) *userRepository.Repository {
	if s.uRep == nil {
		s.uRep = userRepository.NewUserRepository(s.PgPool(ctx))
	}

	return s.uRep
}

func (s *serviceProvider) MessageRepository(ctx context.Context) *messageRepository.Repository {
	if s.mRep == nil {
		s.mRep = messageRepository.NewMessagesRepository(s.PgPool(ctx))
	}

	return s.mRep
}

func (s *serviceProvider) ChatUsecase(ctx context.Context) chatUsecase.UseCase {
	if s.chatUsecase == nil {
		s.chatUsecase = chatUsecase.New(
			s.ChatRepository(ctx),
			s.UserRepository(ctx),
			s.MessageRepository(ctx),
		)
	}

	return *s.chatUsecase
}

func (s *serviceProvider) CImpl(ctx context.Context) *chatApi.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatApi.New(
			s.ChatUsecase(ctx),
		)
	}

	return s.chatImpl
}
