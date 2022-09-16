package log

import "fmt"

func PrintLogo(version, author string) {
	logo := "   ____                               _       _\n" +
		"  / ___| ___  ___  ___  ___  ___ _ __(_) __ _| |\n" +
		" | |  _ / _ \\/ __|/ _ \\/ __|/ _ \\ '__| |/ _` | |\n" +
		" | |_| | (_) \\__ \\ (_) \\__ \\  __/ |  | | (_| | |\n" +
		"  \\____|\\___/|___/\\___/|___/\\___|_|  |_|\\__,_|_|"
	fmt.Println(logo)
	fmt.Printf("version: %s   author: %s\n", version, author)
	fmt.Println("Get Ysoserial Payload by Golang")
}
