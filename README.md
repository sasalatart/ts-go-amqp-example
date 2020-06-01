# TS-GO-AMQP Example

This monorepo project is a quick demo for a simple producer-consumer scenario using AMQP, and how it
can be used to achieve queued communication between different apps, even when written in different
languages.

## Components

### Producer

Written in [TypeScript](https://www.typescriptlang.org/), it sends a high ammount of messages in a
very small window of time. Each message is a JSON-encoded object with an `id` (string) and a
`message` (string) properties.

### Consumer

Written in [Go](https://golang.org/), it is in charge of receiving these messages and processing
them. In order to simulate a "resource-intensive" task, an additional random 5-15 seconds of time
are added to the time taken to process each message. For every logical CPU that the machine running
this demo has, a `goroutine` will be executed to achieve concurrency.

### Messages Broker

[RabbitMQ](https://www.rabbitmq.com/) was set up to queue the messages.

## Usage

1. Clone this repository.
2. To set up the **Messages Broker**, you may either:

   1. Use a cloud service and set the corresponding `AMQP_URI` ENV var, or
   2. Open a new shell and run:
      ```sh
      $ docker run -p 5672:5672 -p 8080:15672 rabbitmq:3.7.17-management-alpine
      ```

3. To start the **Consumer**:

   1. Open a new shell and cd into the `consumer` dir
   2. Run `go run .`

4. To start the **Producer**:

   1. Open a new shell and cd into the `producer` dir
   2. run `yarn install` to install dependencies
   3. run `yarn dev`

If you used the RabbitMQ docker image approach, you will be able to enter its dashboard by visiting
[http://localhost:8080](http://localhost:8080). The credentials should be `guest` for both username
and password. If you used a cloud service, or some other form of RabbitMQ setup, then the way to
access the dashboard may vary.
