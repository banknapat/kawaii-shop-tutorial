package main

import (
	"os"

	"github.com/banknapat/kawaii-shop-tutorial/modules/servers"

	"github.com/banknapat/kawaii-shop-tutorial/config"
	"github.com/banknapat/kawaii-shop-tutorial/pkg/databases"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

func main() {
	cfg := config.LoadConfig(envPath())

	db := databases.DbConnect(cfg.Db())
	defer db.Close()

	servers.NewServer(cfg, db).Start()
}