// Package demo
// Time : 2023/8/15 23:43
// Author : PTJ
// File : main
// Software: GoLand
package main

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, _ := ethclient.Dial("http://1.15.74.74:8545")
	log.Println(client)
	//r := client.Client()
	//log.Println(r)
	number, _ := client.BlockNumber(context.Background())
	log.Println(number)
	//account := common.HexToAddress("0xde411a67AF3061cd13ff67aB511A3E8454456D9C")
	////balance, err := client.BalanceAt(context.Background(), account, nil)
	////if err != nil {
	////	log.Fatal(err)
	////}
	//
	////fmt.Println(balance) // 25893180161173005034
	//
	//blockNumber := big.NewInt(1)
	//balance, err := client.BalanceAt(context.Background(), account, blockNumber)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println(balance) // 25729324269165216042
}
