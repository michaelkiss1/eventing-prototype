module eventing-prototype

go 1.14

require internal/awsTools v0.0.1
replace internal/awsTools => ./internal/awsTools

require internal/events v0.0.1
replace internal/events => ./internal/events

require (
	github.com/aws/aws-sdk-go v1.40.52 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/mustafaturan/bus/v3 v3.0.3 // indirect
)


