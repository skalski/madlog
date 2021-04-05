cd src/app/

#build Win
GOOS=windows GOARCH=amd64 go build -o ../../bin/win/madlog.exe madlog.go

#build OSX
GOOS=darwin GOARCH=amd64 go build -o ../../bin/macos_amd/madlog madlog.go
GOOS=darwin GOARCH=amd64 go build -o ../../bin/macos_arm/madlog madlog.go

#build *nix64
go build -o ../../bin/linux/madlog madlog.go