# Default target
all: build release

# Set the name of the binary
LINUX_BINARY=cli-linux
MAC_ARM_BINARY=cli-mac-arm
WINDOWS_BINARY=cli-windows.exe

# Set the version of the binary
VERSION=v0.0.1

# Set the flags for the Go linker
LDFLAGS=-ldflags "-X main.version=$(VERSION)"

# Build the binaries for different platforms
build:
	GOOS=darwin GOARCH=arm64 go build $(LDFLAGS) -o $(MAC_ARM_BINARY)
	GOOS=linux GOARCH=amd64 go build $(LDFLAGS) -o $(LINUX_BINARY)
	GOOS=windows GOARCH=amd64 go build $(LDFLAGS) -o $(WINDOWS_BINARY)

# Tag the latest commit and create a release
release:
	git tag $(VERSION)
	git push origin $(VERSION)
	gh release create $(VERSION) -t $(VERSION) -n "Release $(VERSION)" $(MAC_ARM_BINARY) $(LINUX_BINARY) $(WINDOWS_BINARY) download.sh

# Clean the binaries
clean:
	rm -f $(MAC_ARM_BINARY) $(LINUX_BINARY) $(WINDOWS_BINARY)
