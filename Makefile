go-build:
	go build -o ./build/app ./src/backend/app.go
	go build -o ./build/monitor ./src/backend/services/monitor.go
	cp ./config.ini ./build
	cd src/frontend/ && npm run build