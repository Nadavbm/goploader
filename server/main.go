package main

import "github.com/nadavbm/goploader/server/api"

func main() {
	s := api.NewServer()

	s.StartServer()
}
