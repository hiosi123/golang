package design

import "fmt"

type IDBconnection interface {
	Connect()
}

type SqlConnection struct {
	connectionString string
}

func (con SqlConnection) Connect() {
	fmt.Println("sql" + con.connectionString)
}

type OracleConnection struct {
	connectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle" + con.connectionString)
}

type DBConnection struct {
	db IDBconnection
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}

func DesignPattern() {
	sqlConnection := SqlConnection{
		connectionString: "SQL Connection Succeed",
	}
	con := DBConnection{db: sqlConnection}
	con.db.Connect()

	orcConnection := OracleConnection{
		connectionString: "Oracle Connection Succeed",
	}
	con2 := DBConnection{db: orcConnection}
	con2.db.Connect()

}
