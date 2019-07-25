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