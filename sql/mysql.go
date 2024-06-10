package sql

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func NewMysqlConnectionPool() *MysqlConnectionPool {
	return &MysqlConnectionPool{
		Username: "root",
		Password: "",
		DBName:   "example",
		IP:       "localhost",
		Port:     "3306",
	}
}

type MysqlConnectionPool struct {
	Username string
	Password string
	DBName   string
	IP       string
	Port     string
}

func (mcp *MysqlConnectionPool) Get() *sqlx.DB {
	return sqlx.MustOpen("mysql", mcp.getDSN())
}

func (mcp *MysqlConnectionPool) getDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mcp.Username, mcp.Password, mcp.IP, mcp.Port, mcp.DBName)
}

func (mcp *MysqlConnectionPool) SetUsername(username string) *MysqlConnectionPool {
	mcp.Username = username
	return mcp
}

func (mcp *MysqlConnectionPool) SetPassword(password string) *MysqlConnectionPool {
	mcp.Password = password
	return mcp
}

func (mcp *MysqlConnectionPool) SetDatabase(database string) *MysqlConnectionPool {
	mcp.DBName = database
	return mcp
}

func (mcp *MysqlConnectionPool) SetIP(ip string) *MysqlConnectionPool {
	mcp.IP = ip
	return mcp
}

func (mcp *MysqlConnectionPool) SetPort(port string) *MysqlConnectionPool {
	mcp.Port = port
	return mcp
}
