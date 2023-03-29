module webserver

go 1.19

replace webserver/server => ./server

require (
	go.uber.org/fx v1.19.2
	webserver/loggerfx v0.0.0-00010101000000-000000000000
	webserver/server v0.0.0-00010101000000-000000000000
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/dig v1.16.1 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/sys v0.0.0-20210903071746-97244b99971b // indirect
)

replace webserver/loggerfx => ./loggerfx
