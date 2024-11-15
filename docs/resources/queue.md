---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "rabbitmq_queue Resource - terraform-provider-rabbitmq"
subcategory: ""
description: |-
  The rabbitmq_queue resource creates and manages a queue.
---

# rabbitmq_queue (Resource)

The rabbitmq_queue resource creates and manages a queue.

## Example Usage

```terraform
# Create a virtual host
resource "rabbitmq_vhost" "example" {
  name = "myvhost"
}

# Create a user
resource "rabbitmq_user" "example" {
  name     = "myuser"
  password = "foobar"
  tags     = ["administrator", "management"]
}

# Create a permission
resource "rabbitmq_permissions" "example" {
  user  = rabbitmq_user.example.name
  vhost = rabbitmq_vhost.example.name

  permissions {
    configure = ".*"
    write     = ".*"
    read      = ".*"
  }
}

# Create a queue
resource "rabbitmq_queue" "example" {
  name  = "myqueue"
  vhost = rabbitmq_permissions.example.vhost

  settings {
    durable     = false
    auto_delete = true
    arguments = {
      "x-queue-type" : "quorum",
    }
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the queue.
- `settings` (Block List, Min: 1, Max: 1) The settings of the queue. The structure is described below. (see [below for nested schema](#nestedblock--settings))

### Optional

- `vhost` (String) The vhost to create the resource in. Defaults to `/`.

### Read-Only

- `id` (String) The ID of this resource.

<a id="nestedblock--settings"></a>
### Nested Schema for `settings`

Optional:

- `arguments` (Map of String) Additional key/value settings for the queue. All values will be sent to RabbitMQ as a string. If you require non-string values, use `arguments_json`.
~> **Note:** Either this or `arguments_json` must be specified but not both.
- `arguments_json` (String) A nested JSON string which contains additional settings for the queue. This is useful for when the arguments contain non-string values.
~> **Note:** Either this or `arguments` must be specified but not both.
- `auto_delete` (Boolean) Whether the queue will self-delete when all consumers have unsubscribed. Defaults to `false`.
- `durable` (Boolean) Whether the queue survives server restarts. Defaults to `false`.