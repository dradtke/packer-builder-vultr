Vultr builder plugin for Packer
===============================

A Packer builder for creating Vultr snapshots.

## Deprecation Notice

This repository is not being actively maintained, and I recommend all
users of this plugin switch over to Vultr's because it will be better
maintained and supported: https://github.com/vultr/packer-builder-vultr

## Installing

- Ensure you have the Packer plugin directory:

`mkdir -p ~/.packer.d/plugins`

- Install the plugin binary:

`GO111MODULE=on go get github.com/dradtke/packer-builder-vultr`

- Move the plugin from where Go put it to the plugin directory:

`mv $(go env GOBIN)/packer-builder-vultr ~/.packer.d/plugins/`

That's it. Builder type `vultr` should now work.

## Configuration

Check out `testdata/vultr.json` for an example Packer file that uses the
Vultr builder.
