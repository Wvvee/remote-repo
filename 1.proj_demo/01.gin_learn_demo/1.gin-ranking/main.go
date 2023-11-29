package main

import (
	"gin-ranking/router"
)

func main() {
	r := router.Router()

	// defer recover panic

	r.Run(":9999")
}
