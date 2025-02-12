---
layout: "vkcs"
page_title: "vkcs: vkcs_compute_keypair"
description: |-
  Manages a keypair resource within VKCS.
---

# vkcs_compute_keypair

Manages a keypair resource within VKCS.

~> **Important Security Notice** The private key generated by this resource will be stored *unencrypted* in your Terraform state file. **Use of this resource for production deployments is *not* recommended**. Instead, generate a private key file outside of Terraform and distribute it securely to the system where Terraform will be run.

## Example Usage
### Import an Existing Public Key
```terraform
resource "vkcs_compute_keypair" "test-keypair" {
  name       = "my-keypair"
  public_key = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDAjpC1hwiOCCmKEWxJ4qzTTsJbKzndLotBCz5PcwtUnflmU+gHJtWMZKpuEGVi29h0A/+ydKek1O18k10Ff+4tyFjiHDQAnOfgWf7+b1yK+qDip3X1C0UPMbwHlTfSGWLGZqd9LvEFx9k3h/M+VtMvwR1lJ9LUyTAImnNjWG7TaIPmui30HvM2UiFEmqkr4ijq45MyX2+fLIePLRIF61p4whjHAQYufqyno3BS48icQb4p6iVEZPo4AE2o9oIyQvj2mx4dk5Y8CgSETOZTYDOR3rU2fZTRDRgPJDH9FWvQjF5tA0p3d9CoWWd2s6GKKbfoUIi8R/Db1BSPJwkqB"
}
```

### Generate a Public/Private Key Pair
```terraform
resource "vkcs_compute_keypair" "keypair" {
  name = "test-keypair"
}

output "public_key" {
  value = vkcs_compute_keypair.keypair.public_key
}

output "private_key" {
  value = vkcs_compute_keypair.keypair.private_key
  sensitive = true
}
```
## Argument Reference
- `name` **String** (***Required***) A unique name for the keypair. Changing this creates a new keypair.

- `public_key` **String** (*Optional*) A pregenerated OpenSSH-formatted public key. Changing this creates a new keypair. If a public key is not specified, then a public/private key pair will be automatically generated. If a pair is created, then destroying this resource means you will lose access to that keypair forever.

- `region` **String** (*Optional*) The region in which to obtain the Compute client. Keypairs are associated with accounts, but a Compute client is needed to create one. If omitted, the `region` argument of the provider is used. Changing this creates a new keypair.

- `value_specs` <strong>Map of </strong>**String** (*Optional*) Map of additional options.


## Attributes Reference
- `name` **String** See Argument Reference above.

- `public_key` **String** See Argument Reference above.

- `region` **String** See Argument Reference above.

- `value_specs` <strong>Map of </strong>**String** See Argument Reference above.

- `fingerprint` **String** The fingerprint of the public key.

- `id` **String** ID of the resource.

- `private_key` **String** The generated private key when no public key is specified.



## Import

Keypairs can be imported using the `name`, e.g.
```shell
terraform import vkcs_compute_keypair.my-keypair test-keypair
```
