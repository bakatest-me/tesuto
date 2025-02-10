tidy:
	go mod tidy
build:
	go build -o tesuto cmd/cli/main.go
build-all:
	GOOS=linux GOARCH=amd64 go build -o out/tesuto ./cmd/cli/main.go
	cd out && tar -czvf tesuto-linux.tar.gz tesuto
	rm out/tesuto
	
	GOOS=linux GOARCH=arm64 go build -o out/tesuto ./cmd/cli/main.go
	cd out && tar -czvf tesuto-linux-arm64.tar.gz tesuto
	rm out/tesuto

	GOOS=darwin GOARCH=amd64 go build -o out/tesuto ./cmd/cli/main.go
	cd out && tar -czvf tesuto-darwin.tar.gz tesuto
	rm out/tesuto
	
	GOOS=windows GOARCH=amd64 go build -o out/tesuto.exe ./cmd/cli/main.go
	cd out && tar -czvf tesuto.exe.tar.gz tesuto.exe
	rm out/tesuto.exe
