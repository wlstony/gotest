module github.com/api

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/rabbitmq v0.0.0-incompatible
)

replace github.com/rabbitmq => ../../../rabbitmq

go 1.13
