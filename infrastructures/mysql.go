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
	Port         string
	DatabaseName string
	Charset      string
	MaxIdle      int
}

// CreateMySQLConnection create my sql connection
func CreateMySQLConnection(name string) *MySQLConnection {
	return &MySQLConnection{
		User:         viper.GetString(fmt.Sprintf("db.%s.user", name)),
		Password:     viper.GetString(fmt.Sprintf("db.%s.password", name)),
		Host:         viper.GetString(fmt.Sprintf("db.%s.host", name)),
		Port:         viper.GetString(fmt.Sprintf("db.%s.port", name)),
		DatabaseName: viper.GetString(fmt.Sprintf("db.%s.name", name)),
		Charset:      viper.GetString(fmt.Sprintf("db.%s.charset", name)),
		MaxIdle:      viper.GetInt(fmt.Sprintf("db.%s.max_idle", name)),
	}
}

// Open return database pool connection
func (conn *MySQLConnection) Open() (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@%s/%s?parseTime=true&loc=%s",
		conn.User,
		conn.Password,
		fmt.Sprintf("tcp(%s:%s)", conn.Host, conn.Port),
		conn.DatabaseName,
		"Local",
	))
	db.SetMaxIdleConns(conn.MaxIdle)
	return db, err
}
