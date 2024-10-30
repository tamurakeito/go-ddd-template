build:
	GO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o docker/main src/main.go

start: build
	cd docker && docker-compose up --build

deploy: build
	eval "$(ssh-agent -s)"
	ssh-add ~/.ssh/gcp_tamurakeito_key
	rsync -avz docker/ tamurakeito@xx.xxx.xx.xx:/home/tamurakeito/go-ddd-template
	rm -r go-ddd-template/main

ssh:
	ssh -i ~/.ssh/gcp_tamurakeito_key tamurakeito@xx.xxx.xx.xx

ini:
	chmod 755 .init.sh && ./.init.sh