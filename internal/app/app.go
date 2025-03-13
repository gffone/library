package app

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"library/config"
	"library/db"
	generated "library/generated/api/library"
	"library/internal/controller"
	"library/internal/usecase/library"
	repository "library/internal/usecase/repository/postgres"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func Run(logger *zap.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	dbPool, err := pgxpool.New(ctx, cfg.PG.URL)
	if err != nil {
		logger.Error("failed to connect to database", zap.Error(err))
		os.Exit(-1)
	}
	defer dbPool.Close()

	db.SetupPostgres(dbPool, logger)

	repo := repository.NewPostgresRepository(dbPool)
	useCases := library.New(logger, repo, repo)

	ctrl := controller.New(logger, useCases, useCases)

	go runRest(ctx, cfg, logger)
	go runGrpc(cfg, logger, ctrl)

	<-ctx.Done()

	time.Sleep(time.Second * 5)
}

func runRest(ctx context.Context, cfg *config.Config, logger *zap.Logger) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	address := "localhost:" + cfg.GRPC.GatewayPort

	err := generated.RegisterLibraryHandlerFromEndpoint(ctx, mux, address, opts)
	if err != nil {
		logger.Error("can not register grpc gateway", zap.Error(err))
		os.Exit(-1)
	}
	getewayPort := ":" + cfg.GRPC.GatewayPort
	logger.Info("start grpc gateway", zap.String("address", getewayPort))

	if err = http.ListenAndServe(getewayPort, mux); err != nil {
		logger.Error("start grpc gateway failed", zap.Error(err))
	}
}

func runGrpc(cfg *config.Config, logger *zap.Logger, libraryService generated.LibraryServer) {
	port := ":" + cfg.GRPC.Port
	lis, err := net.Listen("tcp", port)

	if err != nil {
		logger.Error("can not open tcp socket", zap.Error(err))
		os.Exit(-1)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	generated.RegisterLibraryServer(s, libraryService)

	logger.Info("start grpc server", zap.String("address", port))

	if err = s.Serve(lis); err != nil {
		logger.Error("start grpc server failed", zap.Error(err))
	}
}
