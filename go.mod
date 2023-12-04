module github.com/cdvelop/sessionhandler

go 1.20

require github.com/cdvelop/model v0.0.76

require (
	github.com/cdvelop/object v0.0.40
	github.com/cdvelop/strings v0.0.7 // indirect
	github.com/cdvelop/token v0.0.3
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/token => ../token

replace github.com/cdvelop/object => ../object
