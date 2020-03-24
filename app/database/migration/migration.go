package migration

import "myproject/app/config"
import "myproject/app/model"
func New(config *config.Config)  {
	config.Db.AutoMigrate(&model.User{})
}