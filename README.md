Vultr builder plugin for Packer
===============================

A Packer builder for creating Vultr snapshots.

## Building

Due to vendoring issues, at the moment this package must be built from within
Packer's source code. So, fetch this repo and Packer:

```sh
$ go get -d github.com/hashicorp/packer
$ go get -d github.com/dradtke/packer-builder-vultr
```

Copy the contents of `vultr/` to Packer's source tree:

```sh
$ cp -r ${GOPATH:-~/go}/src/github.com/dradtke/packer-builder-vultr/vultr ${GOPATH:-~/go}/src/github.com/hashicorp/packer/builder/
```

Then open up Packer's file `command/plugin.go` and add Vultr as a new builder:
 - in `import` secion add: `vultrbuilder "github.com/hashicorp/packer/builder/vultr"`
 - in `Builders` map add: `"vultr":   new(vultrbuilder.Builder),`
 
Then you can `go install` Packer, and it will have support for the "vultr"
plugin.

## Configuration

Check out `testdata/vultr.json` for an example Packer file that uses the
Vultr builder.
