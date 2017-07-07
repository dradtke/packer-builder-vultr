package main

import (
	"github.com/dradtke/packer-builder-vultr/vultr"
	"github.com/mitchellh/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(vultr.Builder))
	server.Serve()
}
