package vultr

import (
	"net"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/mitchellh/multistep"
	"golang.org/x/crypto/ssh"
)

func commHost(state multistep.StateBag) (string, error) {
	return state.Get("server").(vultr.Server).MainIP, nil
}

func sshConfig(state multistep.StateBag) (*ssh.ClientConfig, error) {
	c := state.Get("config").(Config)
	server := state.Get("server").(vultr.Server)

	config := &ssh.ClientConfig{
		User: c.SSHUsername,
		Auth: make([]ssh.AuthMethod, 0),
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil // accept anything
		},
	}

	if c.OSID == SnapshotOSID || c.OSID == CustomOSID {
		config.Auth = append(config.Auth, ssh.Password(c.SSHPassword))
	} else {
		config.Auth = append(config.Auth, ssh.Password(server.DefaultPassword))
	}

	return config, nil
}
