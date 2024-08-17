# OpenZiti

Runs OpenZiti's edge tunneler in run-host mode allowing to access talos nodes resources on the overlay network

https://openziti.io

https://openziti.io/docs/reference/tunnelers/docker/#use-case-hosting-openziti-services

## Installation

See [Installing Extensions](https://github.com/siderolabs/extensions#installing-extensions).

## Usage

Create an identity for the machine.
Manually enroll it using(for example using ziti edge enroll ...) and output the enrolled identity into a file or stdout
Copy the contents of the enrolled identity
Create `ExtensionServiceConfig` as the example below with replacing `JSON_IDENTITY` with the JSON of the identity you created and in mountPath replace `IDENTITY_FILE_NAME` with the filename of the identity(preserve the json in the filename)

```yaml
---
apiVersion: v1alpha1
kind: ExtensionServiceConfig
name: openziti
configFiles:
    - content: 'JSON_IDENTITY'
      mountPath: /var/lib/ziti/etc/identities/IDENTITY_FILE_NAME.json
```

Then apply the patch to your node's MachineConfigs
```bash
talosctl patch mc -p @openziti.talos.yaml
```

You will then be able to verify that it is in place with the following command
```bash
talosctl get extensionserviceconfigs

NODE            NAMESPACE   TYPE                     ID         VERSION
192.168.10.10   runtime     ExtensionServiceConfig   openziti   1
```

Example of creation of ziti service to serve talos api on the overlay network(note that 192.168.10.10 is the address of the node):
```
ziti edge create config talosctl-controlplane.intercept.v1 intercept.v1 '{"protocols":["tcp"],"addresses": ["talosctl-controlplane.ziti.internal"], "portRanges":[{"low": 50000, "high":50000}]}'
ziti edge create config talosctl-controlplane.host.v1 host.v1 '{"protocol": "tcp","address":"'"192.168.10.10"'", "port": 50000}'
ziti edge create service talosctrl-controlplane.svc --configs talosctl-controlplane.intercept.v1,talosctl-controlplane.host.v1
ziti edge create service-policy talosctl-controlplane.policy.dial Dial --service-roles "@talosctrl-controlplane.svc" --identity-roles "@macos"
ziti edge create service-policy talosctl-controlplane.policy.bind Bind --service-roles "@talosctrl-controlplane.svc" --identity-roles "@talos-cluster-test-identity"
```
