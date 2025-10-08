package main

type Account struct {
	Id         int    `json:"id" gorm:"primary_key"`
	Number     string `json:"number"`
	MoneyCount int    `json:"money_count"`
}
