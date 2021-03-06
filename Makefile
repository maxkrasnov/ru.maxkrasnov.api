# обрабатывает команду make dev - сделан для тестирования
dev:
	go get github.com/gin-gonic/gin
	go get github.com/jinzhu/gorm
	go get github.com/jinzhu/gorm/dialects/postgres
	go run ./cmd/ru.maxkrasnov.api/main.go
# обрабатывает комнаду make run - запускает сборку проекта и устанавливает нужные зависимости
run:
	go get github.com/gin-gonic/gin
	go get github.com/jinzhu/gorm
	go get github.com/jinzhu/gorm/dialects/postgres
	cd build&&go build ./../cmd/ru.maxkrasnov.api/main.go&&./main
# обрабатывает команду make deploy - запускает сборку докер контейнеров
deploy:
	cd deployments&&docker-compose build
	export APP_VERSION=$(APP_VERSION)&&cd deployments&&docker-compose up -d