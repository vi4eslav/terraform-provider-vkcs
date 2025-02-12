---
layout: "vkcs"
page_title: "Provider: VKCS Provider"
description: |-
  The VKCS provider is used to interact with VKCS services.
  The provider needs to be configured with the proper credentials before it can be used.
---

# VKCS Provider

The VKCS provider is used to interact with [VKCS services](https://mcs.mail.ru/). The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

Terraform 1.0 and later:

```terraform
# Configure the vkcs provider

terraform {
  required_providers {
    vkcs = {
      source = "vk-cs/vkcs"
      version = "~> 0.1.0"
    }
  }
}

# Create new compute instance
resource "vkcs_compute_instance" "myinstance"{
  # ...
}
```

## Authentication

The VKCS provider supports username/password authentication. Preconfigured provider file with `username` and `project_id` can be downloaded from [https://mcs.mail.ru/app/project](https://mcs.mail.ru/app/project) portal. Go to `Terraform` tab -> click on the "Download VKCS provider file".

```terraform
provider "vkcs" {
    username   = "USERNAME"
    password   = "PASSWORD"
    project_id = "PROJECT_ID"
}
```

## Argument Reference
- `auth_url` **String** (*Optional*) The Identity authentication URL.

- `password` **String** (*Optional* Sensitive) Password to login with.

- `project_id` **String** (*Optional*) The ID of Project to login with.

- `region` **String** (*Optional*) A region to use.

- `user_domain_id` **String** (*Optional*) The id of the domain where the user resides.

- `user_domain_name` **String** (*Optional*) The name of the domain where the user resides.

- `username` **String** (*Optional*) User name to login with.



