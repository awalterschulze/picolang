
build:
	docker build -t pico-docker .

delete:
	docker rm pico-add-container || true
	docker rm pico-inc-container || true
	docker rm pico-map-container || true
	docker rm pico-reduce-container || true
	docker ps

run:
	make delete
	docker run -d -t -p 8080:8080 --name pico-add-container pico-docker picoservice add
	docker run -d -t -p 8081:8080 --name pico-inc-container pico-docker picoservice inc
	docker run -d -t -p 8082:8080 --name pico-map-container pico-docker picoservice map
	docker run -d -t -p 8083:8080 --name pico-reduce-container pico-docker picoservice reduce
	docker ps
	go run main.go
	make stop

stop:
	docker kill pico-add-container || true	
	docker kill pico-inc-container || true	
	docker kill pico-map-container || true	
	docker kill pico-reduce-container || true	
	docker ps

boot2docker-start:
	boot2docker init
	boot2docker start
	$(boot2docker shellinit)

