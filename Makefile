containers:
	docker-compose up -d

migrate_up_local:
	migrate -path infrastructure/db/migration -database "mysql://clean_architecture_go_v2:clean_architecture_go_v2@tcp(127.0.0.1:3306)/clean_architecture_go_v2?parseTime=true" -verbose up

migrate_down_local:
	migrate -path infrastructure/db/migration -database "mysql://clean_architecture_go_v2:clean_architecture_go_v2@tcp(127.0.0.1:3306)/clean_architecture_go_v2?parseTime=true" -verbose up

.PHONY: containers migrate_up_local migrate_down_local