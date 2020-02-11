build:
	go build -o ../bot/bot cmd/bot/main.go 
	go build -o ../bot/notification cmd/notification/main.go 
	go build -o ../bot/initdb scripts/init_db.go
	mkdir ../bot/configs && mkdir ../bot/localization 
	cp configs/configs.toml ../bot/configs/configs.toml 
	cp -r localization/lang ../bot/localization