# dapr-pubsub-example

This DAPR example shows how to simply create a struct with data and publish it (checkout). Secondly, (order-processor) picks up the message and unmarshals it back to a Go struct. DAPR is doing all the marshalling, its pretty straightforward.

Run this sample with `dapr run -f .`
You will need to have all the dapr setup up as defined here [Dapr Getting Started](https://docs.dapr.io/getting-started/)