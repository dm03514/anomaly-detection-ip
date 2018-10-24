PKGS = $(shell go list ./... | grep -v /vendor/)
BUILD = "build"
ANOMALY_DETECTION_IP_BIN = "anomaly-detection-ip"
IPSETS_BIN = "ipsets"

fmt:
	go fmt $(PKGS)

lint:
	echo "lint"

build-tools:
	echo "build tools"

build:
	go build -o $(BUILD)/$(ANOMALY_DETECTION_IP_BIN) ./cmd/anomalydetectionip
	go build -o $(BUILD)/$(IPSETS_BIN) ./cmd/ipsets

test-unit:
	go test $(PKGS) -v

docker-dev-image:
	# WARNING WARNING - THIS IS A DEV IMAGE!!!!
	# No security, cleanup or optimization has been performed!
	docker build -t anomaly-detection-ip .

start-dev:
	# update-ipsets enable dshield
	# update-ipsets
	docker run -it \
		-v /var/run/docker.sock:/host/var/run/docker.sock \
		-v /dev:/host/dev \
		-v /proc:/host/proc:ro \
		-v /boot:/host/boot:ro \
		-v /lib/modules:/host/lib/modules:ro \
		-v /usr:/host/usr:ro \
		--net=host \
		--cap-add=ALL \
		--privileged \
		anomaly-detection-ip /bin/bash

start-test-stack:
	docker-compose down && docker-compose up
