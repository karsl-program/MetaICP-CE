package main

import (
	"io"
	"log"
	"metaicp/router"
	"os"
	"strconv"
)

func GetPort() uint {
	config_f, err := os.Open("port")

	if err != nil {
		log.Println("[main] Warning: can not read `port` file to run on port, will use 8080 port...")
		return 8080
	}

	data, errr := io.ReadAll(config_f)

	if errr != nil {
		log.Println("[main] Warning: read `port` file error, will run 8080...")
		return 8080
	}

	port, erra := strconv.Atoi(string(data))

	if erra != nil || port <= 0 {
		log.Println("[main] Warning: invalid `port`, will try run 8080...")
		return 8080
	}

	return uint(port)
}

func main() {
	app := router.InitRouter()

	if err := app.Run(":" + strconv.Itoa(int(GetPort()))); err != nil {
		panic("[main] Error: run error ( " + err.Error() + " )")
	}
}
