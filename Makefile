APP?=go-k8s-microservice
PROJECT?=github.com/Abdujabbor/minikube-go-microservice
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_DATE?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
RELEASE?=0.0.1
GOOS?=linux
GOARCH?=amd64
LDFLAGS?=-ldflags "-X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.BuildDate=${BUILD_DATE}"
CONTAINER_IMAGE?=docker.io/abdujabbor1987/${APP}


clean:
	rm -f ${APP}



build: clean
	go build $(LDFLAGS) -o ${APP}

container: build
	docker build -t $(CONTAINER_IMAGE):$(RELEASE) .

push: container
	docker push $(CONTAINER_IMAGE):$(RELEASE)

test: 
	go test -v ./...