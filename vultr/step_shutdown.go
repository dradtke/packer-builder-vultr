package vultr

import (
	"time"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"golang.org/x/crypto/ssh"
)

type stepShutdown struct{}

func (s *stepShutdown) Run(state multistep.StateBag) multistep.StepAction {
	c := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)
	server := state.Get("server").(vultr.Server)

	ui.Say("Preparing the server for a graceful shutdown...")
	config, err := sshConfig(state)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	client, err := ssh.Dial("tcp", server.MainIP+":22", config)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	session, err := client.NewSession()
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	if c.ShutdownCommand == "" {
		c.ShutdownCommand = "shutdown -P now"
	}
	ui.Say("Shutting down server...")
	session.Run(c.ShutdownCommand)

	ui.Say("Sleeping to ensure that server is shut down...")
	time.Sleep(3 * time.Second)

	return multistep.ActionContinue
}

func (s *stepShutdown) Cleanup(state multistep.StateBag) {
}
