run:
	docker-compose -f ./docker-compose.yaml up -d
	docker run -d -p 8082:8082 --name library-app library-app:1.0
stop:
	docker stop library-app
	docker-compose -f ./docker-compose.yaml down