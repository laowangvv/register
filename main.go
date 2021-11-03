package main

import (
	"ginEssential/common"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	db := common.InitDB()
	log.Println(db)
	r := gin.Default()
	r = CollectRouter(r)
	panic(r.Run())
}



