# External KVM Control for JetKVM
This is a plugin for the JetKVM to add support for controlling an external KVM via RS-232.

## Building
Run `GOOS=linux GOARCH=arm GOARM=7 go build .` to build the `jetkvm-plugin-serialkvm` binary.

Run `tar -czvf serialkvm.tar.gz manifest.json serialkvm` to build the plugin archive.
