module order

go 1.13

require (
	common v0.0.0-00010101000000-000000000000
	github.com/eiji03aero/mskit v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.3.3
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
)

replace (
	common => ../common
	github.com/eiji03aero/mskit => ../../../../
)
