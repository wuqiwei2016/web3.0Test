package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Employee 结构体定义
type Employee struct {
	ID         int    `db:"id"`
	Name       string `db:"name"`       //名称
	Department string `db:"department"` // 部门
	Salary     int    `db:"salary"`     // 薪资
}

// 查询技术部所有员工
func GetTechEmployees(db *sqlx.DB) ([]Employee, error) {
	var employees []Employee

	query := `SELECT id, name, department, salary FROM employees WHERE department = ?`
	err := db.Select(&employees, query, "技术部")

	if err != nil {
		return nil, fmt.Errorf("查询技术部员工失败: %w", err)
	}

	return employees, nil
}

// 查询工资最高的员工
func GetHighestPaidEmployee(db *sqlx.DB) (*Employee, error) {
	var employee Employee

	query := `SELECT id, name, department, salary FROM employees ORDER BY salary DESC LIMIT 1`
	err := db.Get(&employee, query)

	if err != nil {
		return nil, fmt.Errorf("查询最高工资员工失败: %w", err)
	}

	return &employee, nil
}

func main() {
	// 连接SQLite数据库
	db, err := sqlx.Connect("sqlite3", "./gorm.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 查询技术部所有员工
	techEmployees, err := GetTechEmployees(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("技术部员工:")
	for _, emp := range techEmployees {
		fmt.Printf("ID: %d, Name: %s, Department: %s, Salary: %d\n",
			emp.ID, emp.Name, emp.Department, emp.Salary)
	}

	// 查询工资最高的员工
	highestPaid, err := GetHighestPaidEmployee(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n工资最高的员工: ID: %d, Name: %s, Department: %s, Salary: %d\n",
		highestPaid.ID, highestPaid.Name, highestPaid.Department, highestPaid.Salary)
}
