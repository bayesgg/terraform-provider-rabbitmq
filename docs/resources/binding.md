---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "rabbitmq_binding Resource - terraform-provider-rabbitmq"
subcategory: ""
description: |-
  The rabbitmq_binding resource creates and manages a binding relationship between a queue an exchange.
---

# rabbitmq_binding (Resource)

The `rabbitmq_binding` resource creates and manages a binding relationship between a queue an exchange.

## Example Usage

```terraform
# Create a vhost
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

# Create an exchange
resource "rabbitmq_exchange" "example" {
  name  = "myexchange"
  vhost = rabbitmq_vhost.example.vhost

  settings {
    type        = "fanout"
    durable     = false
    auto_delete = true
  }
}

# Create a queue
resource "rabbitmq_queue" "test" {
  name  = "myqueue"
  vhost = rabbitmq_vhost.example.vhost

  settings {
    durable     = true
    auto_delete = false
  }
}

# Create a binding
resource "rabbitmq_binding" "example" {
  source           = rabbitmq_exchange.example.name
  vhost            = rabbitmq_vhost.example.name
  destination      = rabbitmq_queue.example.name
  destination_type = "queue"
  routing_key      = "#"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination` (String) The destination queue or exchange.
- `destination_type` (String) The type of destination. Possible values are `queue` and `exchange`.
- `source` (String) The source exchange.
- `vhost` (String) The vhost to create the resource in.

### Optional

- `arguments` (Map of String) Additional key/value arguments for the binding.
~> **Note:** Either this or `arguments` must be specified but not both.
- `arguments_json` (String) A nested JSON string which contains additional settings for the binding. This is useful for when the arguments contain non-string values.
~> **Note:** Either this or `arguments` must be specified but not both.
- `routing_key` (String) A routing key for the binding.

### Read-Only

- `id` (String) The ID of this resource.
- `properties_key` (String) A unique key to refer to the binding.