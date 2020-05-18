package resources

import (
	"context"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/rs/zerolog/log"
)

func (r *Resources) initDb(ctx context.Context) {
	db, err := gorm.Open("mysql", "test.db")
	if err != nil {
		log.Fatal().Err(err).Msg("Db connection failed")
	}

	r.Db = db

	go func() {
		<-ctx.Done()
		log.Info().Msg("The db connection closing...")
		_ = db.Close()
		log.Info().Msg("The db connection was successfully closed")
	}()
}