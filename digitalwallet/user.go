package digitalwallet

import "sync"

type User struct {
	ID       string
	name     string
	email    string
	password string
	accounts []*Account
	mu       sync.RWMutex
}

func NewUser(id, name, email, password string) *User {
	return &User{
		ID:       id,
		name:     name,
		email:    email,
		password: password,
		accounts: make([]*Account, 0),
	}
}

func (u *User) AddAccount(account *Account) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.accounts = append(u.accounts, account)
}
