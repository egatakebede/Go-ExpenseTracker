package main


import (
	"github.com/gin-gonic/gin"
"net/http"	
)


func main() {
	r := gin.Default()
     


	r.Run(":9080")
}