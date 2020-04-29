package main

import (
	"fmt"

	"github.com/eth_tools/mail"
)

func main() {
	//service.NewInterface()
	to := []string{"huan.cao@vonechain.com", "371902449@qq.com"}
	err := mail.SendMail(to, "test", "this is a test")
	if err != nil {
		fmt.Println(err)
	}
}
