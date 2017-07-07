package vultr

import vultr "github.com/JamesClonk/vultr/lib"

type Artifact struct {
	SnapshotID string
	apiKey     string
}

func (a Artifact) BuilderId() string             { return "vultr" }
func (a Artifact) Files() []string               { return nil }
func (a Artifact) Id() string                    { return a.SnapshotID }
func (a Artifact) String() string                { return "Snapshot: " + a.SnapshotID }
func (a Artifact) State(name string) interface{} { return nil }
func (a Artifact) Destroy() error {
	return vultr.NewClient(a.apiKey, nil).DeleteSnapshot(a.SnapshotID)
}
