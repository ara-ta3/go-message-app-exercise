GO=$(shell which go)

run:
	$(GO) run main.go

ping:
	curl localhost:8888/ping

