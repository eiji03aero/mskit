# mskit
- toolkit for microservices in go

# Todo
- review namings
  - none for now

- cli
  - generate saga

- saga
  - cancelorder saga

- refactor
  - try simplify the initialization in cmd/app :(
  - logging
    - should add logging on
      - publish, consume, rpcendpoint, rpcclient
  - should think about how eventstore and sagastore dealing with restoring data from json
  - registry could be utilized more?

- rpc
  - add timeout

- refactor services
  - align the structure

- setup tests
  - unit tests
  - add comments on exporteds
  - github action
  - reference
    - https://github.com/jetbasrawi/go.cqrs

- update aggregate
  - version
  - snapshot

# Concerns
- better implementation for saga?
  - should remain to be flexible (capable of using messaging/http/grpc)

# Features
- Event sourcing
- Pub/Sub
- RPC with other service
- Saga

# CLI

```sh
# initialize service
$ mskit init order

# cd into service directory
$ cd order

# generate aggregate
$ mskit generate aggregate Order

# generate proxy
$ mskit generate proxy Kitchen

# generate rpcendpoint
$ mskit generate rpcendpoint

# generate publisher
$ mskit generate publisher

# generate consumer
$ mskit generate consumer
```
