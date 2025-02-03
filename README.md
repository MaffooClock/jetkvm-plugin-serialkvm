# External KVM Control for JetKVM

This is a plugin for [JetKVM](https://jetkvm.com/) to add support for controlling an external KVM via RS-232.

A [serial port extension](https://jetkvm.com/docs/peripheral-devices/extension-port) is required to use this plugin, however the official extension device is not yet available yet.  You can build one using a MAX3232 (or similar) IC to convert between the UART in the JetKVM to RS-232 used on the external KVM device.  See [discussion #101](https://github.com/jetkvm/kvm/discussions/101) for the pinouts to build your own. 

> [!NOTE]
> The plugin system is currently in [PR #10](https://github.com/jetkvm/kvm/pull/10) and not yet part of the core functionality.
> 
> However, it can be tested by overwriting the binary from the [Nevexo/jetkvm-kvm](https://github.com/Nevexo/jetkvm-kvm) fork. 

I shamelessly stole @tutman96's [Tailscale plugin](https://github.com/tutman96/jetkvm-plugin-tailscale) to use as a starting point for this plugin.


## To-do:

- [x] Basic functionality for connecting to, configuring, and writing to a local serial port

- [ ] Configure number of inputs, and the serial command for each one <sup>1</sup>

- [ ] Display a dropdown in the Action Bar for switching the input on the remote KVM <sup>2</sup>

**Notes:**
1. This (likely) depends on the "Basic JSON configuration management" bullet in [PR #10](https://github.com/jetkvm/kvm/pull/10), which is noted to be in a separate PR.  Until then, the only way will be to manually edit a JSON configuration file.
2. There is not currently a way to add functionality to the web UI via the proposed plugin system as it sits currently.  However, there is [a comment](https://github.com/jetkvm/kvm/discussions/9#discussioncomment-11756044) that proposes the idea of a simple UI framework that might serve this purpose.

Thus, since it has external dependencies that have not been developed, this plugin is not usable at the moment.


## Building

1. Build the `serialkvm` binary:
   ```bash
   GOOS=linux GOARCH=arm GOARM=7 go build .
   ``` 

2. Build the plugin archive:
   ```bash
   tar -czvf serialkvm.tar.gz manifest.json serialkvm
   ```


## Testing

Since the plugin system won't allow you to re-upload the plugin (it'll complain that the version is the same and abort).  But that's kind of the hard way anyway.  This is (slightly) easier:

1. Disable the plugin (which stops the process)
2. Replace the binary from step 1 in [Building](#building):
   ```bash
   cat serialkvm | ssh root@jetkvm "cat > /userdata/jetkvm/plugins/extracts/<uuid>/serialkvm"
   ```
3. Re-enable the plugin (which will re-start the process using the new binary)
