---
page_title: "cloudflare_waiting_room_event Resource - Cloudflare"
subcategory: ""
description: |-
  Provides a Cloudflare Waiting Room Event resource.
---

# cloudflare_waiting_room_event (Resource)

Provides a Cloudflare Waiting Room Event resource.

## Example Usage

```terraform
# Waiting Room Event
resource "cloudflare_waiting_room_event" "example" {
  zone_id              = "ae36f999674d196762efcc5abb06b345"
  waiting_room_id      = "d41d8cd98f00b204e9800998ecf8427e"
  name                 = "foo"
  event_start_time     = "2006-01-02T15:04:05Z"
  event_end_time       = "2006-01-02T20:04:05Z"
}
```
<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `event_end_time` (String) ISO 8601 timestamp that marks the end of the event.
- `event_start_time` (String) ISO 8601 timestamp that marks the start of the event. Must occur at least 1 minute before `event_end_time`.
- `name` (String) A unique name to identify the event. Only alphanumeric characters, hyphens, and underscores are allowed.
- `waiting_room_id` (String) The Waiting Room ID the event should apply to.
- `zone_id` (String) The zone identifier to target for the resource.

### Optional

- `custom_page_html` (String) This is a templated html file that will be rendered at the edge.
- `description` (String) A description to let users add more details about the event.
- `disable_session_renewal` (Boolean) Disables automatic renewal of session cookies.
- `new_users_per_minute` (Number) The number of new users that will be let into the route every minute.
- `prequeue_start_time` (String) ISO 8601 timestamp that marks when to begin queueing all users before the event starts. Must occur at least 5 minutes before `event_start_time`.
- `queueing_method` (String) The queueing method used by the waiting room. Available values: `"fifo"`, `"random"`, `"passthrough"`, `"reject"`.
- `session_duration` (Number) Lifetime of a cookie (in minutes) set by Cloudflare for users who get access to the origin.
- `shuffle_at_event_start` (Boolean) Users in the prequeue will be shuffled randomly at the `event_start_time`. Requires that `prequeue_start_time` is not null. Defaults to `false`.
- `suspended` (Boolean) If suspended, the event is ignored and traffic will be handled based on the waiting room configuration.
- `total_active_users` (Number) The total number of active user sessions on the route at a point in time.

### Read-Only

- `created_on` (String) Creation time.
- `id` (String) The ID of this resource.
- `modified_on` (String) Last modified time.

## Import

Import is supported using the following syntax:
```shell
# Use the Zone ID, Waiting Room ID, and Event ID to import.
$ terraform import cloudflare_waiting_room_event.default <zone_id>/<waiting_room_id>/<waiting_room_event_id>
```