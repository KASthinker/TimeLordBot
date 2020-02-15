.PHONY: all dir bot notification initDB

all: dir bot notification initDB

dir:
	@rm -rf bin
	@mkdir -p bin/bot/configs && mkdir -p bin/bot/localization
	@cp configs/helpconf.toml bin/bot/configs/helpconf.toml 
	@cp -r localization/lang bin/bot/localization
	@echo [OK]: Auxiliary files created.

bot:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/bot/bot cmd/bot/main.go
	@echo [OK]: The bot is built.

notification:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/bot/notification cmd/notification/main.go
	@echo [OK]: The notification is built.

initDB:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o bin/bot/init_db scripts/init_db.go
	@echo [OK]: The initDB is built.