package app

import (
	"context"

	"github.com/noskov-sergey/platform-common/pkg/closer"
	"github.com/noskov-sergey/platform-common/pkg/db"
	"github.com/noskov-sergey/platform-common/pkg/db/pg"
	"github.com/noskov-sergey/platform-common/pkg/db/transaction"

	"github.com/noskov-sergey/chat-server/internal/config"
	"log"

	chatApi "github.com/noskov-sergey/chat-server/internal/api/chats"
	chatRepository "github.com/noskov-sergey/chat-server/internal/repository/chats"
	messageRepository "github.com/noskov-sergey/chat-server/internal/repository/messages"
	userRepository "github.com/noskov-sergey/chat-server/internal/repository/user"
	chatUsecase "github.com/noskov-sergey/chat-server/internal/usecase/chats"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager
	cRep      *chatRepository.Repository
	uRep      *userRepository.Repository
	mRep      *messageRepository.Repository

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

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %v", err)
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.HTTPConfig {
	if s.swaggerConfig == nil {
		cfg, err := config.NewSWAGGERConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %v", err)
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %v", err)
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) ChatRepository(ctx context.Context) *chatRepository.Repository {
	if s.cRep == nil {
		s.cRep = chatRepository.NewChatRepository(s.DBClient(ctx))
	}

	return s.cRep
}

func (s *serviceProvider) UserRepository(ctx context.Context) *userRepository.Repository {
	if s.uRep == nil {
		s.uRep = userRepository.NewUserRepository(s.DBClient(ctx))
	}

	return s.uRep
}

func (s *serviceProvider) MessageRepository(ctx context.Context) *messageRepository.Repository {
	if s.mRep == nil {
		s.mRep = messageRepository.NewMessagesRepository(s.DBClient(ctx))
	}

	return s.mRep
}

func (s *serviceProvider) ChatUsecase(ctx context.Context) *chatUsecase.UseCase {
	if s.chatUsecase == nil {
		s.chatUsecase = chatUsecase.New(
			s.ChatRepository(ctx),
			s.UserRepository(ctx),
			s.MessageRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatUsecase
}

func (s *serviceProvider) CImpl(ctx context.Context) *chatApi.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatApi.New(
			s.ChatUsecase(ctx),
		)
	}

	return s.chatImpl
}
