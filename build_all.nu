$env.CGO_ENABLED = "0"

GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o lzb-windows-amd64.exe .
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o lzb-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -o lzb-linux-arm64 .
GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s" -o lzb-darwin-arm64 .