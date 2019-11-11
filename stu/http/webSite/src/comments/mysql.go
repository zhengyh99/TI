package comments

import (
	"database/sql"
	"fmt"
)

var dbSource = "root:123@tcp(localhost:3306)/database"

type MySQLDB struct {
	db     *sql.DB
	stmt   *sql.Stmt
	rows   *sql.Rows
	result sql.Result
	rowsAffected
}

func NewMySQLDB() (db *MySQLDB) {
	mydb, err := sql.Open("mysql", dbSource)
	if err != nil {
		return &MySQLDB{db: mydb}
	}
	fmt.Println("sql open error:", err)
	return nil
}

func (mydb *MySQLDB) Close() {
	if mydb.rows != nil {
		mydb.rows.Close()
	}
	if mydb.stmt != nil {
		mydb.stmt.Close()
	}
	if mydb.db != nil {
		mydb.db.Close()
	}
}
//Prepare创建一个准备好的状态用于之后的查询和命令。返回值可以同时执行多个查询和命令。
func (mydb *MySQLDB) prepare(sql) (err error) {
	mydb.stmt, err = mydb.db.Prepare(sql)
	if err != nil {
		fmt.Println("DB prepare error:", err)
		return err
	}
	return nil
}
//Exec使用提供的参数执行准备好的命令状态，返回Result类型的该状态执行结果的总结。
func (mydb *MySQLDB) runDML(sql string, args ...interface{}) (err error) {

	err = mydb.prepare(sql)
	mydb.result, err = mydb.stmt.Exec(args...)
	if err != nil {
		fmt.Println("stmt exec error:", err)
	}
	return

}

//Query使用提供的参数执行准备好的查询状态，返回Rows类型查询结果。
func (mydb *MySQLDB) runDQL(sql string, args ...interface{}) (err error) {

	err = mydb.prepare(sql)
	mydb.rows, err = mydb.stmt.Query(args...)
	if err != nil {
		fmt.Println("stmt exec error:", err)
	}
	return

}
// LastInsertId返回一个数据库生成的回应命令的整数。
// 当插入新行时，一般来自一个"自增"列。
func (mydb *MySQLDB) Insert(sql string, args ...interface{}) (lastId int64, err error) {
	err = mydb.runDML(sql, args...)
	if err != nil {
		return -1, err
	}
	return mydb.result.LastInsertId(), nil
}
// RowsAffected返回被update、insert或delete命令影响的行数。
func (mydb *MySQLDB) UpdOrDel(sql string, args ...interface{}) (rowsAffected int64, err error) {
	err = mydb.runDML(sql, args...)
	if err != nil {
		return -1, err
	}
	return mydb.result.RowsAffected(), nil
}
//Rows是查询的结果。它的游标指向结果集的第零行，使用Next方法来遍历各行结果
func (mydb *MySQLDB) Select(sql string, args ...interface{}) (rows *sql.Rows, err error) {
	err = mydb.runDQL(sql, args...)
	if err != nil {
		return nil, err
	}
	return mydb.rows, nil
}
