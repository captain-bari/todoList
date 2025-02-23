package database

import "fmt"

func NewDatabaseConnection() *database {
	return &database{}
}

type conn struct {
}

func (c *conn) sendQuery(query string) {
	fmt.Println("MOCK-Database : SEND Query :", query)
}

func (c *conn) close() error {
	fmt.Println("MOCK-Database : close connection")
	return nil
}

type database struct {
	conn *conn
}

func (d *database) CloseConnection() error {
	return d.conn.close()
}

func (d *database) AddList() {
	d.conn.sendQuery("Add list")
}

func (d *database) GetListDetails(uuid string) {
	d.conn.sendQuery("GetListDetails:" + uuid)
}
