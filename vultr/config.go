package vultr

import (
	"github.com/hashicorp/packer/common"
	"github.com/hashicorp/packer/helper/communicator"
	"github.com/hashicorp/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	ctx                 interpolate.Context
	Comm                communicator.Config `mapstructure:",squash"`

	APIKey string `mapstructure:"api_key"`
	// RootPass string `mapstructure:"root_pass"`

	Description     string `mapstructure:"description"`
	RegionID        int    `mapstructure:"region_id"`
	RegionName      string `mapstructure:"region_name"`
	RegionCode      string `mapstructure:"region_code"`
	PlanID          int    `mapstructure:"plan_id"`
	PlanName        string `mapstructure:"plan_name"`
	OSID            int    `mapstructure:"os_id"`
	OSName          string `mapstructure:"os_name"`
	ScriptID        int    `mapstructure:"script_id"`
	SnapshotID      string `mapstructure:"snapshot_id"`
	SSHKey          string `mapstructure:"ssh_key_id"`
	SSHUsername     string `mapstructure:"ssh_username"`
	SSHPassword     string `mapstructure:"ssh_password"`
	ShutdownCommand string `mapstructure:"shutdown_command"`

	/*
		DatacenterID   int    `mapstructure:"datacenter_id"`
		DatacenterName string `mapstructure:"datacenter_name"`

		PlanID   int    `mapstructure:"plan_id"`
		PlanName string `mapstructure:"plan_name"`

		DistributionID   int    `mapstructure:"distribution_id"`
		DistributionName string `mapstructure:"distribution_name"`

		KernelID   int    `mapstructure:"kernel_id"`
		KernelName string `mapstructure:"kernel_name"`

		DiskSize int    `mapstructure:"disk_size"`
		RootPass string `mapstructure:"root_pass"`

		Label string

		Description string // optional
		RootSSHKey  string // optional
		PaymentTerm int    // optional
	*/

	interCtx interpolate.Context
}
