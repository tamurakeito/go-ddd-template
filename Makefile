APP_ENV ?= development
BUILD_DIR = docker
APP_NAME = main

build: generate-env
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIR)/$(APP_NAME) src/main.go

generate-env:
	@echo "Generating .env file for $(APP_ENV) environment..."
	@if [ "$(APP_ENV)" = "production" ]; then \
		SECRET_KEY=$$(LC_ALL=C tr -dc 'A-Za-z0-9' </dev/urandom | head -c 32); \
		echo "APP_ENV=production" > $(BUILD_DIR)/.env; \
		echo "JWT_SECRET_KEY=$$SECRET_KEY" >> $(BUILD_DIR)/.env; \
	else \
		SECRET_KEY=$$(LC_ALL=C tr -dc 'A-Za-z0-9' </dev/urandom | head -c 32); \
		echo "APP_ENV=development" > $(BUILD_DIR)/.env; \
		echo "JWT_SECRET_KEY=$$SECRET_KEY" >> $(BUILD_DIR)/.env; \
	fi

start: build
	cd docker && docker-compose up --build

deploy: build APP_ENV=production
	eval "$(ssh-agent -s)"
	ssh-add ~/.ssh/gcp_tamurakeito_key
	rsync -avz docker/ tamurakeito@xx.xxx.xx.xx:/home/tamurakeito/go-ddd-template
	rm -r go-ddd-template/main

ssh:
	ssh -i ~/.ssh/gcp_tamurakeito_key tamurakeito@xx.xxx.xx.xx

mock:
	mockgen -source=src/domain/repository/repository.go -destination=mocks/repository/mock_repository.go -package=mocks
	mockgen -source=src/domain/repository/line/repository.go -destination=mocks/repository/line/mock_repository.go -package=mocks
	mockgen -source=src/service/auth_service.go -destination=mocks/service/mock_auth_service.go -package=mocks
	mockgen -source=src/service/encrypt_service.go -destination=mocks/service/mock_encrypt_service.go -package=mocks

tests:
	gotests -w -all ./src/usecase/$(FILE)
	gotests -w -all ./src/infrastructure/repository_impl/$(FILE)
	gotests -w -all ./src/usecase/line/$(FILE)
	gotests -w -all ./src/infrastructure/repository_impl/line/$(FILE)

unit_test:
	go test ./... -v

local_test:
	open tools/local_integration_test/general.html

ini:
	chmod 755 .init.sh && ./.init.sh