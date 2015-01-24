
build:
	docker build -t pico-docker .

run:
	docker rm pico-double-container || true
	docker run -d -t -p 8080:8080 --name pico-double-container pico-docker picoservice double
	docker rm pico-inc-container || true
	docker run -d -t -p 8081:8080 --name pico-inc-container pico-docker picoservice inc
	docker rm pico-log-container || true
	docker run -d -t -p 8082:8080 --name pico-log-container pico-docker picoservice log
	docker rm pico-map-container || true
	docker run -d -t -p 8083:8080 --name pico-map-container pico-docker picoservice map
	docker ps

stop:
	docker kill pico-double-container || true	
	docker kill pico-inc-container || true	
	docker kill pico-log-container || true	
	docker kill pico-map-container || true	
	docker ps

boot2docker-start:
	boot2docker init
	boot2docker start
	$(boot2docker shellinit)

