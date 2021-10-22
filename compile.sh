CGO_ENABLED=0 
#compile for linux
GOOS=linux 
GOARCH=amd64 
go build -o ./tmp_uploader_linux main.go
#compile for windows
GOOS=windows
GOARCH=amd64
go build -o ./tmp_uploader_windows.exe main.go
#compile for macos intel x86_64
GOOS=darwin
GOARCH=amd64
go build -o ./tmp_uploader_macos_intel main.go
#compile for macos m1 arm64
GOOS=darwin
GOARCH=arm64
go build -o ./tmp_uploader_macos_arm64 main.go