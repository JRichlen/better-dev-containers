install:
	go get
	go build

test:
	@echo "TODO: Implement testing..."

build:
	go build
	cp "${CURDIR}/better-dev-container.exe" "${GOPATH}/bin/better-dev-container.exe"
	cp "${CURDIR}/better-dev-container.exe" "${GOPATH}/bin/bdev.exe"
	bdev go build

