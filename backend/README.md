# Backend

This go module contains the backend that is basically a web socket relay.

| Endpoint        | Description                                                  | Format      |
| --------------- | ------------------------------------------------------------ | ----------- |
| `/robots`       | Get a list of the currently connected robots and their connection status. | `json`      |
| `ws/robot/:id`  | The robot firmware connects to this endpoint with its name as `:id`. This id will be shown in `/robots` | `websocket` |
| `ws/remote/:id` | The frontend app connects to this endpoint with a valid robot id from the `/robots` list. If successful both websockets are connected together. | `websocket` |
