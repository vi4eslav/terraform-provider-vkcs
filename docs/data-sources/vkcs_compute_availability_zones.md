---
layout: "vkcs"
page_title: "vkcs: vkcs_compute_availability_zones"
description: |-
  Get a list of availability zones from VKCS
---

# vkcs_compute_availability_zones

Use this data source to get a list of availability zones from VKCS

## Example Usage

```terraform
data "vkcs_compute_availability_zones" "zones" {}
```

## Argument Reference
- `region` **String** (*Optional*) The `region` to fetch availability zones from, defaults to the provider's `region`

- `state` **String** (*Optional*) The `state` of the availability zones to match, default ("available").


## Attributes Reference
- `region` **String** See Argument Reference above.

- `state` **String** See Argument Reference above.

- `id` **String** Hash of the returned zone list.

- `names` **String** The names of the availability zones, ordered alphanumerically, that match the queried `state`


