go-build:
	go build -o ./build/app ./src/backend/app.go
	cd src/frontend/ && npm run build
	cp ./config.ini ./build