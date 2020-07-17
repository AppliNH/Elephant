package commands

import "fmt"

func Help() {

	help := `
Commands:

		walk	Start a stack by providing the path to your docker-compose file
		stomp	Stop a stack by providing an elephant name
		ls	Allows you to see all the elephants and the containers they carry
		help	Displays this help`

	fmt.Println(help)
}
