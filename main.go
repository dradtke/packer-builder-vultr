package main

import (
	"github.com/hashicorp/packer/packer/plugin"

	"github.com/dradtke/packer-builder-vultr/vultr"
)

func main() {
	plugin.RegisterBuilder(new(vultr.Builder))
}
