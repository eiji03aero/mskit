# mskit
- toolkit for microservices in go

# Todo

```
- review namings
  - none

- refactor
  - try simplify the initialization in cmd/app :(
    - gave up for now
    - kind of waiting for generics to see what it's gotta bring on table

- setup tests
  - unit tests
  - add comments on exporteds

- update aggregate
  - version
  - snapshot
```

# Concerns
- better implementation for saga?
  - should remain to be flexible (capable of using messaging/http/grpc)
- better way to detect not found on aggregates load with mongo

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
$ mskit generate domain:aggregate Order

# generate domain service
$ mskit generate domain:service Order

# generate proxy
$ mskit generate proxy Kitchen

# generate rpcendpoint
$ mskit generate rpcendpoint

# generate publisher
$ mskit generate publisher

# generate consumer
$ mskit generate consumer

# generate saga
$ mskit generate saga CreateOrder
```
