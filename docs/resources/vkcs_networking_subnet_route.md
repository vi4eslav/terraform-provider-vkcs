---
layout: "vkcs"
page_title: "vkcs: vkcs_networking_subnet_route"
description: |-
  Creates a routing entry on a VKCS subnet.
---

# vkcs_networking_subnet_route

Creates a routing entry on a VKCS subnet.

## Example Usage
```terraform
resource "vkcs_networking_router" "router_1" {
  name           = "router_1"
  admin_state_up = "true"
}

resource "vkcs_networking_network" "network_1" {
  name           = "network_1"
  admin_state_up = "true"
}

resource "vkcs_networking_subnet" "subnet_1" {
  network_id = "${vkcs_networking_network.network_1.id}"
  cidr       = "192.168.199.0/24"
}

resource "vkcs_networking_subnet_route" "subnet_route_1" {
  subnet_id        = "${vkcs_networking_subnet.subnet_1.id}"
  destination_cidr = "10.0.1.0/24"
  next_hop         = "192.168.199.254"
}
```
## Argument Reference
- `destination_cidr` **String** (***Required***) CIDR block to match on the packet’s destination IP. Changing this creates a new routing entry.

- `next_hop` **String** (***Required***) IP address of the next hop gateway. Changing this creates a new routing entry.

- `subnet_id` **String** (***Required***) ID of the subnet this routing entry belongs to. Changing this creates a new routing entry.

- `region` **String** (*Optional*) The region in which to obtain the networking client. A networking client is needed to configure a routing entry on a subnet. If omitted, the `region` argument of the provider is used. Changing this creates a new routing entry.

- `sdn` **String** (*Optional*) SDN to use for this resource. Must be one of following: "neutron", "sprut". Default value is "neutron".


## Attributes Reference
- `destination_cidr` **String** See Argument Reference above.

- `next_hop` **String** See Argument Reference above.

- `subnet_id` **String** See Argument Reference above.

- `region` **String** See Argument Reference above.

- `sdn` **String** See Argument Reference above.

- `id` **String** ID of the resource.



## Import

Routing entries can be imported using a combined ID using the following format: ``<subnet_id>-route-<destination_cidr>-<next_hop>``

```shell
terraform import vkcs_networking_subnet_route.subnet_route_1 686fe248-386c-4f70-9f6c-281607dad079-route-10.0.1.0/24-192.168.199.25
```
