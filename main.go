//08183742921

package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/tobi007/cdServer/db"
	"github.com/tobi007/cdServer/server"
	"github.com/tobi007/cdServer/config"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()


}
