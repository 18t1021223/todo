package config

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func ConnectDB(env *Env, embedMigrations embed.FS) *sql.DB {
	url := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		env.DBUser, env.DBPass, env.DBHost, env.DBPort, env.DBName, env.DBParams,
	)
	connect, err := sql.Open(env.DBDriver, url)
	if err != nil {
		zap.L().Error("Failed ConnectDB", zap.Error(err))
		panic(err)
	}

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect(env.DBDriver); err != nil {
		zap.L().Error("Failed migration", zap.Error(err))
		panic(err)
	}
	if err := goose.Up(connect, "."); err != nil {
		log.Fatal(err)
	}
	zap.L().Info("Migrations completed!")
	return connect
}
