docker-run:
	docker rmi main-service-img || true
	docker build -t main-service-img .

	docker-compose up

run:
	export POSTGRES_HOST=c-c9qt1ll800dqgld2lds5.rw.mdb.yandexcloud.net && \
	export POSTGRES_PORT=6432 && \
	export POSTGRES_USER=bolat && \
	export POSTGRES_PASSWORD=Voro123456 && \
	export POSTGRES_DB=voronezhack && \
	go run cmd/service.go