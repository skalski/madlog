#build Win
GOOS=windows GOARCH=386 go build -o madlog_32.exe madlog.go
GOOS=windows GOARCH=amd64 go build -o madlog.exe madlog.go

#build *nix64
go build -o madlog madlog.go