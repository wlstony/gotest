module github.com/data

require (
	github.com/rabbitmq v0.0.0-incompatible
	google.golang.org/genproto v0.0.0-20210804223703-f1db76f3300d
)

replace github.com/rabbitmq => ../../../rabbitmq

go 1.13
