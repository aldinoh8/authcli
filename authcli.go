package authcli

import (
	"github.com/aldinoh8/authcli/cli"
	"github.com/aldinoh8/authcli/config"
	"github.com/aldinoh8/authcli/handler"
)

func New(dbUrl string) cli.Cli {
	db := config.InitDatabase(dbUrl)
	userHandler := handler.User{DB: db}

	return cli.Cli{UserHandler: userHandler}
}
