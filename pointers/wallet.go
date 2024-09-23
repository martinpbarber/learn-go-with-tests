package wallet

import "fmt"

// Bitcoin represents the number of Bitcoins.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoins.
type Wallet struct {
	balance Bitcoin
}

// Balance returns the number of Bitcoins in a wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit adds Bitcoins to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw removes Bitcoins from a wallet.
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
