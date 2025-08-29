$env.CGO_ENABLED = "0"

mkdir dist
GOOS=windows GOARCH=amd64 go build -ldflags "-w -s" -o dist/lzpb-windows-amd64.exe .
GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o dist/lzpb-linux-amd64 .
GOOS=linux GOARCH=arm64 go build -ldflags "-w -s" -o dist/lzpb-linux-arm64 .
GOOS=darwin GOARCH=arm64 go build -ldflags "-w -s" -o dist/lzpb-darwin-arm64 .