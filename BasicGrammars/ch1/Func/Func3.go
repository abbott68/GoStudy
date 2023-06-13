package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/rand"
	"time"
)

// 定义银行账户类型
type BankAccout struct {
	AccoutNumber string
	Balance      float64
}

// 函数：生成随机的银行账户号码
func generateAccountNumber() string {
	rand.Seed(time.Now().UnixNano())
	accountNumber := ""
	for i := 0; i < 8; i++ {
		accountNumber += string(rand.Intn(10) + '0')
	}
	return accountNumber
}

// 函数：创建一个新的银行账户
func createAccout() BankAccout {
	accountNumber := generateAccountNumber()
	account := BankAccout{
		AccoutNumber: accountNumber,
		Balance:      0,
	}
	return account
}

// 函数：存款
func deposit(account *BankAccout, amount float64) {
	account.Balance += amount
}

// 函数：取款
func withdraw(account *BankAccout, amount float64) bool {
	if amount > account.Balance {
		return false
	}
	account.Balance -= amount
	return true
}

// 主函数
func main() {

	// 创建一个新的银行账户
	account := createAccout()
	fmt.Println("Account Number:", account.AccoutNumber)
	fmt.Println("Initial Blance:", account.Balance)

	// 生成随机的银行账户号码
	deposit(&account, 1000)
	fmt.Println("Blance after deposit:", account.Balance)

	// 取款
	withdrawAmount := 500
	success := withdraw(&account, float64(withdrawAmount))
	if success {
		fmt.Println("Withdrawal of", withdrawAmount, "successful")
	} else {
		fmt.Println("Insufficient balance for withdrawal")
	}
	fmt.Println("Remaining Balance:", account.Balance)

}
