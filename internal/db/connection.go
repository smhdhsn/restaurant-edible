package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/smhdhsn/restaurant-menu/internal/config"
)

// Connect creates a database connection.
func Connect(c config.DBConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		c.User,
		c.Pass,
		c.Host,
		c.Port,
		c.Name,
	)

	return gorm.Open(mysql.Open(dsn))
}
