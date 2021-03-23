## Setting UP

Inicie os containeres da aplicação
```bash
docker-compose up
```
ou
```bash
make containers
```

Rode as migrações
```bash
migrate -path db/migration -database "[USUARIO]:[SENHA]@tcp([HOST]:3306)/[DATABASE]?parseTime=true" -verbose up
```
ou
```bash
make migrate_up_local
```