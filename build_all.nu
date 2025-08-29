$env.CGO_ENABLED = "0"

GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o lzpb-windows-amd64.exe .
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o lzpb-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -o lzpb-linux-arm64 .
GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s" -o lzpb-darwin-arm64 .