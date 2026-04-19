package main

import (
	ctx "github.com/andygr1n1/go-revolution/internal/globalctx"
	appLogger "github.com/andygr1n1/go-revolution/internal/logger"
	utilities "github.com/andygr1n1/go-revolution/library/utilities"
)

func main() {
	logger := appLogger.NewLogger()

	ctx := ctx.GlobalContext{Logger: logger}

	utilities.FmtExample(ctx)
}
