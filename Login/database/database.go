package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"login/config"
	"strconv"
)

// DB a global db object will be used across different packages
var DB *gorm.DB

// InitDatabase creates a sqlite db
func InitDatabase() (err error) {
	var erro error
	p := config.Config("DB_PORT")
	porta, erro := strconv.ParseUint(p, 10,32)

	s := config.Config("DB_PASSWORD")
	senha, erro := strconv.ParseUint(s, 10,32)

	StringDeConexao := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Config("DB_HOST"), porta, config.Config("DB_USER"), senha, config.Config("DB_NAME"))

	DB, erro = gorm.Open(postgres.Open(StringDeConexao))
	if erro != nil{
		panic("Falha ao concetar com o DATABASE")
	}
	fmt.Println("Conexão com o DATABASE foi aberta")
	return
}