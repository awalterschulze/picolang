
build:
	docker build -t pico-docker .

run:
	docker rm pico-inc-container || true
	docker run -d -t -p 8080:8080 --name pico-inc-container pico-docker picoservice inc
	docker rm pico-map-container || true
	docker run -d -t -p 8081:8080 --name pico-map-container pico-docker picoservice map
	docker ps

stop:
	docker kill pico-inc-container || true	
	docker kill pico-map-container || true	
	docker ps

boot2docker-start:
	boot2docker init
	boot2docker start
	$(boot2docker shellinit)

