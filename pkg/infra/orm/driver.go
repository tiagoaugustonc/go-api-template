package orm

import (
	"fmt"

	"example.com/cadastro-power/pkg/infra/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDialect(config *config.Config) (gorm.Dialector, error) {

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Database.Cadastro.UserName,
		config.Database.Cadastro.Password,
		config.Database.Cadastro.Host,
		config.Database.Cadastro.Port,
		config.Database.Cadastro.Database,
	)

	return mysql.Open(dns), nil
}
