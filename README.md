# api-store
## Uma API simples usando clean arch e DDD

### Dependência:
* redis (porta = 6379)
### Variáveis de ambiente:
* `SENDER` (esta variável deverá conter o email do remetente)
* `PASS` (esta variável deverá conter a senha do remetente)
* `SECRET_KEY` (está variável será responsável por guardar a chave secreta do JWT)

### Execução:
```
go mod tidy
cd cmd
go run main.go
```
