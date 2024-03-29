package server

import (
	"context"
	"github.com/Egor123qwe/GraphQL-project/internal/app/server/graphqlServer"
	"github.com/Egor123qwe/GraphQL-project/internal/app/server/grpcServer"
	"github.com/Egor123qwe/GraphQL-project/internal/config"
	"github.com/Egor123qwe/GraphQL-project/internal/handlers/grpcHandlers"
	"github.com/Egor123qwe/GraphQL-project/internal/scenarios"
	"github.com/Egor123qwe/GraphQL-project/internal/storage"
	"github.com/Egor123qwe/GraphQL-project/internal/storage/mongoDb"
	"golang.org/x/exp/slog"
	"golang.org/x/sync/errgroup"
	"log"
)

type server struct {
	config    *config.Config
	scenarios scenarios.User
	storage   storage.Storage
	ctx       context.Context
}

func New() *server {
	ctx := context.Background()

	storage, err := mongoDb.New(ctx)
	if err != nil {
		log.Fatalln("error in connection mongo : %w", err)
	}
	return &server{
		config:    config.New(),
		scenarios: scenarios.New(storage),
		storage:   storage,
		ctx:       ctx,
	}
}

func (s *server) Start() error {
	gr, ctx := errgroup.WithContext(s.ctx)

	gr.Go(func() error {
		return grpcServer.New(grpcHandlers.New(s.scenarios), s.config).Start(ctx)
	})
	gr.Go(func() error {
		return graphqlServer.New(s.scenarios, s.config).Start(ctx)
	})
	if err := gr.Wait(); err != nil {
		slog.Error(err.Error())
	}
	return nil
}
