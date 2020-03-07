module kitchen

go 1.13

require (
	common v0.0.0-00010101000000-000000000000
	github.com/eiji03aero/mskit v0.0.0-20200306004028-eab1f19aeccb
	go.mongodb.org/mongo-driver v1.3.0
)

replace (
	common => ../common
	github.com/eiji03aero/mskit => ../../../../
)
