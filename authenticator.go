package main

import (
	"fmt"
	"time"
)

type Authenticator struct {
	accountRepository AccountRepository
}

func (a *Authenticator) setupAuth(userId string, name string, secret string) {
	err := a.accountRepository.SetUpAccount(userId, name, secret)
	if err != nil {
		fmt.Println("CHANGE ME")
	}
}

func (a *Authenticator) getOTPs(userId string) map[string]string {
	accounts, err := a.accountRepository.ListAccounts(userId)
	
	if err != nil {
		return nil
	}

	unixTime := time.Now().Unix()
	counter := unixTime / 30

	accToOtp := make(map[string]string)

	for k, v := range accounts {
		otp := generate(v, counter)
		accToOtp[k] = otp
	}

	return accToOtp
}