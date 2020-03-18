package main

import (
	"os"

	"github.com/DevYukine/cdn/cdn"
)

func main() {
	cdn.NewHTTPServer(os.Getenv("CONTENT_ROOT"), os.Getenv("PORT"), os.Getenv("TOKEN")).Serve()
}
