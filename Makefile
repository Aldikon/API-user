.PHONY: dockerup

dcup:
	docker-compose up --build -d

dcdown:
	docker-compose down

logapp:
	docker-compose logs -f app

logdb:
	docker-compose logs -f db
