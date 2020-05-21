package resources

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
)

func (r *Resources) initDb(ctx context.Context) {
	p := make(map[string]string)
	p["charset"] = "utf8mb4"

	config := mysql.NewConfig()
	config.User = r.Config.DbUser
	config.Passwd = r.Config.DbPassword
	config.Net = "tcp"
	config.Addr = r.Config.DbHost
	config.DBName = r.Config.DbName
	config.Collation = "utf8mb4_unicode_ci"
	config.Params = p
	config.ParseTime = true

	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal().Err(err).Msg("Db connection failed")
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")

	r.Db = db

	go func() {
		<-ctx.Done()
		log.Info().Msg("The db connection closing...")
		_ = db.Close()
		log.Info().Msg("The db connection was successfully closed")
	}()
}