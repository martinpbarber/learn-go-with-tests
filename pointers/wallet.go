package wallet

// Wallet stores the number of Bitcoins.
type Wallet struct {
	balance int
}

// Balance returns the number of Bitcoins in a wallet.
func (w *Wallet) Balance() int {
	return w.balance
}

// Deposit adds bitcoins to the wallet.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
