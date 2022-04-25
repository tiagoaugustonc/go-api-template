package orm

import (
	"example.com/cadastro-power/pkg/infra/config"
	"gorm.io/gorm"
)

func NewDB(config *config.Config) (*gorm.DB, error) {

	dialect, _ := newDialect(config)

	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
