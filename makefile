tidy:
	go mod tidy
build:
	go build -o tesuto cmd/cli/main.go
build-all:
	export VERSION=0.0.1
	# Build the application for multiple platforms
	GOOS=linux GOARCH=amd64 go build -o out/tesuto-linux ./cmd/cli/main.go
	GOOS=linux GOARCH=arm64 go build -o out/tesuto-linux-arm64 ./cmd/cli/main.go
	GOOS=darwin GOARCH=amd64 go build -o out/tesuto-darwin ./cmd/cli/main.go
	GOOS=windows GOARCH=amd64 go build -o out/tesuto.exe ./cmd/cli/main.go

	tar -czvf out/tesuto-linux-${VERSION}.tar.gz out/tesuto-linux
	tar -czvf out/tesuto-linux-arm64-${VERSION}.tar.gz out/tesuto-linux-arm64
	tar -czvf out/tesuto-darwin-${VERSION}.tar.gz out/tesuto-darwin
	tar -czvf out/tesuto.exe-${VERSION}.tar.gz out/tesuto.exe

	rm out/tesuto-linux
	rm out/tesuto-linux-arm64
	rm out/tesuto-darwin
	rm out/tesuto.exe
