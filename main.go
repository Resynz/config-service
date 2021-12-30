/**
 * @Author: Resynz
 * @Date: 2021/12/28 11:08
 */
package main

import (
	"config-service/config"
	"config-service/server"
	"log"
)

func main() {
	log.Println("initial env ...")
	if err := config.InitEnv(); err != nil {
		log.Fatalf("init env failed! error:%s\n", err.Error())
	}
	log.Println("\033[42;30m DONE \033[0m[Config-Service] Start Success!")
	server.StartServer()
}
