package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/TestardR/geo-port/internal/application/service"
	"github.com/TestardR/geo-port/internal/domain"
	portCLI "github.com/TestardR/geo-port/internal/infrastructure/cli"
	inMemoryStore "github.com/TestardR/geo-port/internal/infrastructure/in_memory_store"
	"github.com/urfave/cli/v2"
)

const AddOrUpdateFilePathArgument = "filepath"

type StdLogger interface {
	Printf(format string, v ...interface{})
}

func RunAsCLIHandler(
	cliCtx *cli.Context,
	consoleOutput StdLogger,
) error {
	ctx := cliCtx.Context

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	consoleOutput.Printf("CLI handler mode")

	// Rough argument validation
	filePathArgument := cliCtx.String(AddOrUpdateFilePathArgument)
	if len(filePathArgument) == 0 {
		return fmt.Errorf("missing --%s option", AddOrUpdateFilePathArgument)
	}

	portSvc := service.NewPortService(
		inMemoryStore.NewPortStore(make(map[domain.PortID]domain.Port)),
		domain.NewPortValidator(),
	)

	// TODO: implement shared logger infrastructure adapter
	// Industry recommendation: https://github.com/uber-go/zap
	// Explore value brought by new sLog package, go 1.21
	cliHandler := portCLI.NewPortCLI(
		portSvc,
		consoleOutput,
		slog.New(slog.NewJSONHandler(os.Stderr, nil)),
	)

	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-stop
		consoleOutput.Printf("Stopping prematurely geo-port service")
		cancel()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		cliHandler.Handle(ctx, filePathArgument)
	}()

	wg.Wait()

	fmt.Println("Exiting geo-port service")

	return nil
}
