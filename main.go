package main

import (
	"fmt"

	"github.com/superjcd/webdev_tookit/jwt"
)

func main() {
	infos := make(map[string]string, 0)
	infos["username"] = "jack"

	tokenString, err := jwt.GenerateToken(infos, 3)

	if err != nil {
		panic(err)
	}

	fmt.Println(tokenString)
	headers := "Bearer " + tokenString

	token, _ := jwt.GetTokenFromHeader(headers)

	customeInfos, _ := jwt.ParseCustomInfosFromToken(token)

	fmt.Println(customeInfos.Infos["username"])
}
