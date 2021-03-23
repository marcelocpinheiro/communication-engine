## Setting UP

Inicie os containeres da aplicação o comando abaixo
```bash
make containers
```

Rode as migrações
```bash
make migrate_up_local
```

Execute a API
```bash
go run api/main.go
```

## Opções

Para desfazer todas as migrações, você pode rodar o comando
```bash
make migrate_down_local
```