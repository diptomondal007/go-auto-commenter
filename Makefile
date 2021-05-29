build_all: build_mac build_linux build_linux_386 build_windows build_windows_386

build_mac: ## BUILD for Mac
		@export GOARCH="amd64"
		@export GOOS="darwin"
		@export CGO_ENABLED=1
		@go build -o mac_amd64 -v .

build_linux: ### BUILD for Linux
		@export GOARCH="amd64"
		@export GOOS="linux"
		@export CGO_ENABLED=0
		@go build -o linux_amd64 -v

build_linux_386: ### BUILD for Linux 386
		@export GOARCH="386"
		@export GOOS="linux"
		@export CGO_ENABLED=0
		@go build -o linux_i386 -v

build_windows_386: ###WINDOWS
		@export GOARCH="386"
		@export GOOS="windows"
		@export CGO_ENABLED=0
		@go build -o windows_386.exe -v


build_windows: ### WINDOWS
		@export GOARCH="amd64"
		@export GOOS="windows"
		@export CGO_ENABLED=0
		@go build -o windows_amd64.exe -v

clean:	### Clean built files
		@rm -f mac_amd64 linux_amd64 linux_i386 windows_386.exe windows_amd64.exe
