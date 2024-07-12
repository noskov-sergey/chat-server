package main

import (
	"context"
	"flag"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/noskov-sergey/chat-server/internal/config"
	desc "github.com/noskov-sergey/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("CreateMethod User names: %v", req.GetUsernames())

	builder := sq.Insert("chats").
		PlaceholderFormat(sq.Dollar).
		Columns("id").
		Values(sq.Expr("DEFAULT")).
		Suffix("RETURNING id")

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("to sql: %v", err)
	}

	var insertedID int

	if err = s.pool.QueryRow(ctx, sqlQuery, args...).Scan(&insertedID); err != nil {
		log.Printf("query row: %v", err)
		return nil, fmt.Errorf("query row: %w", err)
	}

	for _, n := range req.Usernames {
		builder_user := sq.Insert("chat_user").
			PlaceholderFormat(sq.Dollar).
			Columns("chat_id", "username").
			Values(insertedID, n).
			Suffix("RETURNING id")

		sqlQuery, args, err = builder_user.ToSql()
		if err != nil {
			log.Fatalf("to sql: %v", err)
		}

		var insertedUserID int

		if err = s.pool.QueryRow(ctx, sqlQuery, args...).Scan(&insertedUserID); err != nil {
			log.Printf("query row: %v", err)
			return nil, fmt.Errorf("query row: %w", err)
		}
	}

	return &desc.CreateResponse{
		Id: int64(insertedID),
	}, nil
}

func (s *server) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteMethod - Chat id: %v", req.GetId())

	builder := sq.
		Delete("chats").
		PlaceholderFormat(sq.Dollar).
		Where("id = ?", int(req.Id))

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("to sql: %v", err)
	}

	if _, err = s.pool.Exec(ctx, sqlQuery, args...); err != nil {
		log.Printf("exec row: %v", err)
		return nil, fmt.Errorf("exec row: %w", err)
	}

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("SendMessage - To chat_id: %d, From id: %s, Text %sd", req.GetChatId(), req.GetFrom(), req.GetText())

	builder := sq.Insert("messages").
		PlaceholderFormat(sq.Dollar).
		Columns("username", "chat_id", "text").
		Values(req.From, req.ChatId, req.Text).
		Suffix("RETURNING id")

	sqlQuery, args, err := builder.ToSql()
	if err != nil {
		log.Fatalf("to sql: %v", err)
	}

	var insertedID int

	if err = s.pool.QueryRow(ctx, sqlQuery, args...).Scan(&insertedID); err != nil {
		log.Printf("query row: %v", err)
		return nil, fmt.Errorf("query row: %w", err)
	}

	return &emptypb.Empty{}, nil
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

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{pool: pool})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
