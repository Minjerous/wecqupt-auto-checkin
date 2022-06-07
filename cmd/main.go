package main

import (
	"log"
	"wcy_health_check/consts"
	"wcy_health_check/work"
)

func main() {
	log.Println("WeCY Auto Heath Check_in Start")
	log.Print(consts.PiKaQu)
	work.DaKaWork()
}
