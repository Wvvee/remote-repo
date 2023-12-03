package api

import (
	"fmt"
	"go_blog/utils"
)

func Encryt(a string) {
	a = "123456"
	passwd := utils.Md5Crypt(a, "mszlu")
	fmt.Println(passwd)

}

//2948c60d3d7938275daa3051034c8fe1
