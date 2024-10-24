package main

import (
	"context"
	"github.com/ethereum/go-ethereum/log"
	"github.com/zhulida1234/go-rpc-service/common/opio"
	"os"
)

var (
	GitCommit = ""
	GitData   = ""
)

func main() {
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))
	app := NewCli(GitCommit, GitData)
	ctx := opio.WithInterruptBlocker(context.Background())
	if err := app.RunContext(ctx, os.Args); err != nil {
		log.Error("Application failed")
		os.Exit(1)
	}

}
