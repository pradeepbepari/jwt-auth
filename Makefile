up:
	docker-compose -f docker-compose.yaml up --build

run:
	docker-compose build && docker-compose up

down:
	docker-compose -f docker-compose.yaml down

show:
	aws --endpoint-url=http://localhost:4566 s3 ls s3://bucket-demo


.PHONY: up run down