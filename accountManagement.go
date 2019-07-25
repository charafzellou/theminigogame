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