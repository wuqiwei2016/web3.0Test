package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint    `gorm:"primaryKey"`
	Balance float64 // 余额
}

type Transaction struct {
	ID            uint    `gorm:"primaryKey"`
	FromAccountID uint    // 转出账户ID
	ToAccountID   uint    // 转入账户ID
	Amount        float64 // 金额
}

/*
*实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务
 */
func main() {
	// 连接数据库
	db, err := gorm.Open(sqlite.Open("transfer.db"), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败")
	}

	// 自动迁移表结构
	err = db.AutoMigrate(&Account{}, &Transaction{})
	if err != nil {
		panic("迁移表结构失败")
	}

	// 假设账户A初始余额200，账户B初始余额100
	db.Create(&Account{ID: 1, Balance: 200})
	db.Create(&Account{ID: 2, Balance: 100})

	// 执行转账
	err = TransferMoney(db, 1, 2, 100)
	if err != nil {
		fmt.Println("转账失败:", err)
	} else {
		fmt.Println("转账成功")
	}
}

func TransferMoney(db *gorm.DB, fromID, toID uint, amount float64) error {
	// 开始事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// 1. 检查转出账户余额是否足够
	var fromAccount Account
	if err := tx.First(&fromAccount, fromID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("查询转出账户失败: %v", err)
	}

	if fromAccount.Balance < amount {
		tx.Rollback()
		return fmt.Errorf("账户余额不足")
	}

	// 2. 扣除转出账户金额
	if err := tx.Model(&Account{}).Where("id = ?", fromID).
		Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("扣除金额失败: %v", err)
	}

	// 3. 增加转入账户金额
	if err := tx.Model(&Account{}).Where("id = ?", toID).
		Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("增加金额失败: %v", err)
	}

	// 4. 记录交易信息
	transaction := Transaction{
		FromAccountID: fromID,
		ToAccountID:   toID,
		Amount:        amount,
	}
	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("记录交易失败: %v", err)
	}

	// 提交事务
	return tx.Commit().Error
}
