---
layout: "vkcs"
page_title: "vkcs: vkcs_lb_member"
description: |-
  Manages a member resource within VKCS.
---

# vkcs_lb_member

Manages a member resource within VKCS.

## Example Usage
```terraform
resource "vkcs_lb_member" "member_1" {
	address       = "192.168.199.23"
	pool_id       = "935685fb-a896-40f9-9ff4-ae531a3a00fe"
	protocol_port = 8080
}
```
## Argument Reference
- `address` **String** (***Required***) The IP address of the member to receive traffic from the load balancer. Changing this creates a new member.

- `pool_id` **String** (***Required***) The id of the pool that this member will be assigned to. Changing this creates a new member.

- `protocol_port` **Number** (***Required***) The port on which to listen for client traffic. Changing this creates a new member.

- `admin_state_up` **Boolean** (*Optional*) The administrative state of the member. A valid value is true (UP) or false (DOWN). Defaults to true.

- `name` **String** (*Optional*) Human-readable name for the member.

- `region` **String** (*Optional*) The region in which to obtain the Loadbalancer client. If omitted, the `region` argument of the provider is used. Changing this creates a new member.

- `subnet_id` **String** (*Optional*) The subnet in which to access the member. Changing this creates a new member.

- `weight` **Number** (*Optional*) A positive integer value that indicates the relative portion of traffic that this member should receive from the pool. For example, a member with a weight of 10 receives five times as much traffic as a member with a weight of 2. Defaults to 1.


## Attributes Reference
- `address` **String** See Argument Reference above.

- `pool_id` **String** See Argument Reference above.

- `protocol_port` **Number** See Argument Reference above.

- `admin_state_up` **Boolean** See Argument Reference above.

- `name` **String** See Argument Reference above.

- `region` **String** See Argument Reference above.

- `subnet_id` **String** See Argument Reference above.

- `weight` **Number** See Argument Reference above.

- `id` **String** ID of the resource.



## Import

Load Balancer Pool Member can be imported using the Pool ID and Member ID separated by a slash, e.g.:

```shell
terraform import vkcs_lb_member.member_1 c22974d2-4c95-4bcb-9819-0afc5ed303d5/9563b79c-8460-47da-8a95-2711b746510f
```
