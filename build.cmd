SET GOOS=windows
SET GOARCH=amd64
go build -o releases/sweetssl.exe
SET GOOS=linux
go build -o releases/sweetssl