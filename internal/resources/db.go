package resources

import (
	"context"

	_ "github.com/jinzhu/gorm/dialects/mysql" // for mysql connect.

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

func (r *Resources) initDB(ctx context.Context) {
	p := make(map[string]string)
	p["charset"] = "utf8mb4"

	config := mysql.NewConfig()
	config.User = r.Config.DBUser
	config.Passwd = r.Config.DBPassword
	config.Net = "tcp"
	config.Addr = r.Config.DBHost
	config.DBName = r.Config.DBName
	config.Collation = "utf8mb4_unicode_ci"
	config.Params = p
	config.ParseTime = true

	db, err := gorm.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal().Err(err).Msg("DB connection failed")
	}
	db = db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8")

	r.DB = db

	go func() {
		<-ctx.Done()
		log.Info().Msg("The db connection closing...")
		_ = db.Close()
		log.Info().Msg("The db connection was successfully closed")
	}()
}
