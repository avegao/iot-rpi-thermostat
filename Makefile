VCS_REF = $(shell git rev-parse --verify HEAD)
BUILD_DATE = $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

build-arm6:
	GOOS=linux GOARCH=arm GOARM=6 go build -v -o build/arm32v6/iot-thermostat main.go

build-arm6-docker:
	docker build \
		-t registry.gitlab.com/avegao/iot/thermostat/arm32v6\:develop \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		-f ./Dockerfile.arm32v6 .

build-arm7:
	docker build \
		-t registry.gitlab.com/avegao/iot/thermostat/arm32v7\:develop \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		-f ./Dockerfile.arm32v7 .

build-amd64:
	docker build \
		-t registry.gitlab.com/avegao/iot/thermostat\:develop \
		--build-arg="VCS_REF=$(VCS_REF)" \
		--build-arg="BUILD_DATE=$(BUILD_DATE)" \
		.

build: build-arm6 build-arm6-docker build-arm7 build-amd64
