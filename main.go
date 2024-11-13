package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os/signal"
	"rest/config"
	"rest/handler"
	"rest/usecase/base"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// bad practice (sandbox only version)
	cfg := config.NewConfig(8881)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		panic("err")
		// log.Fatal
	}

	baseService := base.NewUsecase(ctx)
	ctrl := handler.New(baseService)
	engine := handler.Router(ctrl)

	server := &http.Server{
		Handler:     engine,
		BaseContext: func(net.Listener) context.Context { return ctx },
	}

	wg := new(sync.WaitGroup)
	defer wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()

		for _, shutdowner := range []interface{ Shutdown(context.Context) error }{server} {
			if err := shutdowner.Shutdown(shutdownCtx); err != nil {
				fmt.Println("serious problem with shutind down")
				return
			}
		}
		fmt.Println("gracefully shuted down")
	}()

	switch err = server.Serve(lis); {
	case err == nil, errors.Is(err, http.ErrServerClosed):
	default:
		return
	}
}
