release:
	GOOS=linux GOARCH=amd64 go build -o hashverifier-for-lyra2rev2-linux
	GOOS=darwin GOARCH=amd64 go build -o hashverifier-for-lyra2rev2-darwin
	GOOS=windows GOARCH=amd64 go build -o hashverifier-for-lyra2rev2-windows.exe
