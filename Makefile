build:
	docker build -t pico-docker .

run:
	docker rm pico-container | true
	docker run -d -t -p 80:8080 --name pico-container pico-docker 
	docker ps

stop:
	docker kill pico-container
	docker ps

boot2docker-more-memory:
	boot2docker stop
	VBoxManage modifyvm boot2docker-vm --memory 4096

boot2docker-start:
	boot2docker init
	boot2docker start
	$(boot2docker shellinit)

gofmt:
	gofmt -l -s -w .