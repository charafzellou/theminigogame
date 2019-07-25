package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

var (
	accountsList []assetAccount
	loggedAccount int
)

type assetAccount struct {
	AccountId     string
	AccountPwd    [32]byte
	Player        assetPlayer
}

func signIn() bool{
	accountId := ""
	clearPassword := ""

	fmt.Print("Account Id: ")
	_, _ = fmt.Scanf("%s\n", &accountId)

	fmt.Print("Password: \033[38;5;232m")
	_, _ = fmt.Scanf("%s\n", &clearPassword)
	fmt.Print("\033[39;49m")
	password := sha256.Sum256([]byte(clearPassword))

	if idxUser := login(accountId, password); idxUser != -1 {
		loggedAccount = idxUser
		return true
	}

	return false
}

func register(){
	newAccount := assetAccount{}

	for {
		fmt.Print("Account Id: ")
		_, _ = fmt.Scanf("%s\n", &newAccount.AccountId)
		if accountExist(newAccount.AccountId) && newAccount.AccountId != "" {
			println("Id already taken. Please retry")
		} else {
			break
		}
	}

	clearPassword := ""
	fmt.Print("Password: \033[38;5;232m")
	_, _ = fmt.Scanf("%s\n", &clearPassword)
	fmt.Print("\033[39;49m")
	newAccount.AccountPwd = sha256.Sum256([]byte(clearPassword))

	newAccount.Player = createPlayer()
	accountsList = append(accountsList, newAccount)
	saveAccounts()
}

func accountExist(accountId string) bool{
	for _, account := range accountsList {
		if account.AccountId == accountId {
			return true
		}
	}
	return false
}

func login(account string, password [32]byte) int{
	for idx, listedAccount := range accountsList {
		if listedAccount.AccountId == account && listedAccount.AccountPwd == password {
			return idx
		}
	}
	return -1
}

// not pure att all
func getAccounts(){
	content, err := ioutil.ReadFile(".data/accounts.secure")
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &accountsList)
	if err != nil {
		log.Fatal(err)
	}
}
// not pure either
func saveAccounts(){
	jsonAccounts, err := json.Marshal(accountsList)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(".data/accounts.secure", jsonAccounts, 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Account successfully saved")
	}
}