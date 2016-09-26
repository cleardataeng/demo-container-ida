ida: ida.go
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ida .

.PHONY: clean docker

clean:
	rm -rf *~ ida

docker: ida
	docker build -t rvandegrift/ida:latest .
