module accounting

go 1.13

require (
	github.com/eiji03aero/mskit v0.0.0-20200329125241-9a73db08f856
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
)

replace (
	common => ../common
	github.com/eiji03aero/mskit => ../../../../
)
