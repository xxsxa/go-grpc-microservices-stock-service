build:
	protoc -I. --go_out=plugins=micro:. \
	  proto/stock/stock.proto
	GOOS=linux GOARCH=amd64 go build -o stock-service && \
	docker build -t stock-service .
run:
		docker run -p 50051:50051 \
		    -e MICRO_SERVER_ADDRESS=:50051 \
		    stock-service