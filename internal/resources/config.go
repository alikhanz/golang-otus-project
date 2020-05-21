package resources

import "github.com/alikhanz/golang-otus-project/internal/config"

func (r *Resources) initConfig() {
	r.Config = config.New()
}
