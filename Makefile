build: generate-env
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o docker/main src/main.go

generate-env:
	@SECRET_KEY=$$(LC_ALL=C tr -dc 'A-Za-z0-9' </dev/urandom | head -c 12); \
	echo "JWT_SECRET_KEY=$$SECRET_KEY" > ./docker/.env

start: build
	cd docker && docker-compose up --build

deploy: build
	eval "$(ssh-agent -s)"
	ssh-add ~/.ssh/gcp_tamurakeito_key
	rsync -avz docker/ tamurakeito@xx.xxx.xx.xx:/home/tamurakeito/go-ddd-template
	rm -r go-ddd-template/main

ssh:
	ssh -i ~/.ssh/gcp_tamurakeito_key tamurakeito@xx.xxx.xx.xx

mock:
	mockgen -source=src/domain/repository/repository.go -destination=mocks/repository/mock_repository.go -package=mocks
	mockgen -source=src/service/auth_service.go -destination=mocks/service/mock_auth_service.go -package=mocks
	mockgen -source=src/service/encrypt_service.go -destination=mocks/service/mock_encrypt_service.go -package=mocks

tests:
	gotests -w -all ./src/usecase/$(FILE)
	gotests -w -all ./src/infrastructure/repository_impl/$(FILE)

unit_test:
	go test ./... -v

local_test:
	open tools/local_integration_test/test.html

ini:
	chmod 755 .init.sh && ./.init.sh