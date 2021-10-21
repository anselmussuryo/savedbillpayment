mysql:
	docker run --name mysql8 -p 3307:3306 -e MYSQL_USER=adminDB -e MYSQL_PASSWORD=simaS123 -e MYSQL_ROOT_PASSWORD=simaS123 -e MYSQL_DATABASE=savedbillpaymentdb -d mysql:8.0.26

mysqlWithCopyData:
	docker run --name mysql8 -p 3307:3306 -e MYSQL_USER=adminDB -e MYSQL_PASSWORD=simaS123 -e MYSQL_ROOT_PASSWORD=simaS123 -e MYSQL_DATABASE=savedbillpaymentdb -v /Users/anselmussuryo/Database/MySQL:/var/lib/mysql -d mysql:8.0.26

mysqllogin:
	mysql -h localhost -P 3307 --protocol=tcp -u root -p

mysqlrm:
	docker rm mysql8

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migration -database "mysql://adminDB:simaS123@tcp(localhost:3307)/savedbillpaymentdb" -verbose up

migratedown:
	migrate -path db/migration -database "mysql://adminDB:simaS123@tcp(localhost:3307)/savedbillpaymentdb" -verbose down

test:
	go test -v -cover ./...

.PHONY: mysql mysqlWithCopyData mysqlrm migrateup migratedown sqlc