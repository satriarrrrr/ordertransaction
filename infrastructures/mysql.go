package infrastructures

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// SQLConnection represents SQL connection
type SQLConnection interface {
	Open() (*sql.DB, error)
}

// MySQLConnection type of my sql connection
type MySQLConnection struct {
	User         string
	Password     string
	Host         string
	Port         int
	DatabaseName string
	Charset      string
	MaxIdle      int
}

// CreateMySQLConnection create my sql connection
func CreateMySQLConnection(name string) *MySQLConnection {
	return &MySQLConnection{
		User:         viper.GetString(fmt.Sprintf("database.%s.username", name)),
		Password:     viper.GetString(fmt.Sprintf("database.%s.password", name)),
		Host:         viper.GetString(fmt.Sprintf("database.%s.host", name)),
		Port:         viper.GetInt(fmt.Sprintf("database.%s.port", name)),
		DatabaseName: viper.GetString(fmt.Sprintf("database.%s.name", name)),
		Charset:      viper.GetString(fmt.Sprintf("database.%s.charset", name)),
		MaxIdle:      viper.GetInt(fmt.Sprintf("database.%s.max_idle", name)),
	}
}

// Open return database pool connection
func (conn *MySQLConnection) Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&loc=%s",
		conn.User,
		conn.Password,
		fmt.Sprintf("tcp(%s:%d)", conn.Host, conn.Port),
		conn.DatabaseName,
		"Local",
	))
	db.SetMaxIdleConns(conn.MaxIdle)
	return db, err
}
