package vultr

import (
	"time"

	vultr "github.com/JamesClonk/vultr/lib"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
)

type stepCreate struct {
	v *vultr.Client
}

func (s *stepCreate) Run(state multistep.StateBag) multistep.StepAction {
	c := state.Get("config").(Config)
	ui := state.Get("ui").(packer.Ui)

	opts := &vultr.ServerOptions{
		Script:               c.ScriptID,
		DontNotifyOnActivate: true,
	}
	if c.OSID == SnapshotOSID {
		opts.Snapshot = c.SnapshotID
	}
	server, err := s.v.CreateServer("Snapshotting: "+c.Description, c.RegionID, c.PlanID, c.OSID, opts)
	if err != nil {
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}
	state.Put("server-creation-time", time.Now())

	ui.Message("Server " + server.ID + " created, waiting for it to become active...")
	for server.Status != "active" {
		time.Sleep(1 * time.Second)
		if server, err = s.v.GetServer(server.ID); err != nil {
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	state.Put("server", server)
	return multistep.ActionContinue
}

func (s *stepCreate) Cleanup(state multistep.StateBag) {
	ui := state.Get("ui").(packer.Ui)
	server := state.Get("server").(vultr.Server)
	startTime := state.Get("server-creation-time").(time.Time)

	wait := (5 * time.Minute) - time.Now().Sub(startTime)
	if wait > 0 {
		ui.Say("Vultr requires you to wait 5 minutes before destroying a server, we have " + wait.String() + " left...")
		time.Sleep(wait)
	}

	ui.Say("Destroying server " + server.ID)
	if err := s.v.DeleteServer(server.ID); err != nil {
		state.Put("error", err)
	}
}
