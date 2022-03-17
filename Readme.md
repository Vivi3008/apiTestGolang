# Api HTTP Rest de transferencias entre contas de um banco digital

## Rotas

### Criação de conta - Request

- Path: `/accounts`
- Method: `POST`
- Content-Type: `application/json`
- Body :

```json
{
  "name": "Wonder Woman",
  "cpf": 3333,
  "secret": "wonder",
  "balance": 900
}
```

Sendo que o balance da conta a ser criada não é obrigatório, sendo inicializada assim com 0.

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
{
  "id": "91dae4c2-97f4-4e19-9156-2551a7bf21a0",
  "name": "Wonder Woman",
  "cpf": 3333,
  "balance": 900,
  "createdAt": "2021-10-13T14:47:32.447513422-03:00"
}
```

##### Failure

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "This cpf already exists"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):
  ```json
  {
    "reason": "Cpf must have 11 caracters"
  }
  ```

### Listar todas as contas - Request

- Path: `/accounts`
- Method: `GET`

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
[
  {
    "id": "91dae4c2-97f4-4e19-9156-2551a7bf21a0",
    "name": "Spider Man",
    "cpf": 1111,
    "balance": 0,
    "createdAt": "2021-10-13T14:47:32.447513422-03:00"
  },
  {
    "id": "381443cb-c52c-429b-bd1a-990fcbd9d2fc",
    "name": "Wonder Woman",
    "cpf": 3333,
    "balance": 900,
    "createdAt": "2021-10-13T14:42:41.647594446-03:00"
  }
]
```

### Listar o saldo de uma conta pelo Id

- Path: `/accounts/{account_id}/balance`
- Method: `GET`

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
{
  "balance": 0
}
```

##### Failure

- Status code: `400`
- Content-Type: `application/json`
- Body (example):
  ```json
  {
    "reason": "Id does not exist"
  }
  ```

### Login

A rota de login retorna um token valido para ser usado nas rotas `/transfers`

- Path: `/login`
- Method: `POST`
- Body :

```json
{
  "cpf": 66568899564,
  "secret": "123"
}
```

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijc4Mjg2NmVmLTBiZmUtNDRhNi04MTk4LWZlYTk3YjIzYjg0MyJ9.x8rSh2h-Lm_P-zFTYHB-CmzDHYGmXf-KtCRM_YyISQg"
}
```

##### Failure

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Cpf does not exist"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Password invalid"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):
  ```json
  {
    "reason": "Invalid token"
  }
  ```

### Fazer uma transferência entre contas

Para acessar essa rota o usuario precisa se autenticar definindo o Auth no header com o token gerado no login.

- Path: `/transfers`
- Method: `POST`
- Header: `Authorization: token`
- Body :

```json
{
  "account_destination_id": "de7cb18e-5799-4f08-be4e-f69c2288e3ea",
  "amount": 10.5
}
```

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
{
  "id": "3422fdcc-37c4-480e-9954-a38502a2cb9b",
  "account_origin_id": "3d092a28-2c5e-4af2-bf12-b90010cc45fa",
  "account_destination_id": "de7cb18e-5799-4f08-be4e-f69c2288e3ea",
  "amount": 10.5,
  "createdAt": "2021-10-18T12:13:51.288333982-03:00"
}
```

##### Failure

- Status code: `401`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Auth required"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Account destiny id can't be the same account origin id"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Account origin id doesn't exists"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Account destiny id doesn't exists"
  }
  ```

- Status code: `400`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Insufficient Limit"
  }
  ```

### Listar todas as transferencias do usuario autenticado

Para acessar essa rota o usuario precisa se autenticar definindo o Auth no header com o token gerado no login.

- Path: `/transfers`
- Method: `GET`
- Header: `Authorization: token`

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
[
  {
    "id": "3422fdcc-37c4-480e-9954-a38502a2cb9b",
    "account_origin_id": "3d092a28-2c5e-4af2-bf12-b90010cc45fa",
    "account_destination_id": "de7cb18e-5799-4f08-be4e-f69c2288e3ea",
    "amount": 10,
    "createdAt": "2021-10-18T12:13:51.288333982-03:00"
  }
]
```

##### Failure

- Status code: `401`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Auth required"
  }
  ```

### Adicionar um pagamento

- Path: `/bills`
- Method: `POST`
- Header: `Authorization: token`

- Body :

```json
{
	"description": "Conta aleatoria",
	"value": 300,
	"due_date": "2022-01-05T17:35:18Z"
}
```

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
 {
	"id": "62451496-8b10-46cc-88e6-f924d2b409f0",
	"account_id": "02938f42-b7b5-44c5-90b2-51501c3c58d9",
	"description": "Conta aleatoria",
	"value": 300,
	"due_date": "2022-01-05T17:35:18Z",
	"scheduled_date": "2022-03-03T12:07:24.792614951-03:00",
	"status": "Pago",
	"created_at": "2022-03-03T12:07:24Z"
}
```
##### Failure

- Status code: `401`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Auth required"
  }
  ```

 - Status code: `500`
- Content-Type: `application/json`
- Body (example):

  ```json
  {
    "reason": "Insufficient Limit"
  }
  ```

### Listar pagamentos

- Path: `/bills`
- Method: `GET`
- Header: `Authorization: token`

#### Response Suscess

- Status code: `200`
- Content-Type: `application/json`
- Body :

```json
[
	{
		"id": "62451496-8b10-46cc-88e6-f924d2b409f0",
		"account_id": "02938f42-b7b5-44c5-90b2-51501c3c58d9",
		"description": "Conta aleatoria",
		"value": 300,
		"due_date": "2022-01-05T17:35:18Z",
		"scheduled_date": "2022-03-03T12:07:24.792614951-03:00",
		"status": "Pago",
		"created_at": "2022-03-03T12:07:24Z"
	}
]
```

## Usage

Faça o clone deste repositorio, entre na pasta apiTestGolang, crie o arquivo .env na raiz do projeto com as variaveis do .env.example e defina a porta que a aplicação ira rodar e o Access_secret que pode ser uma string de valor qualquer que serve para gerar o token.

Para criar a imagem da aplicação com docker digite no terminal.

`docker image build -t apitestegolang .`

Para rodar o container na porta 3000:

`docker run -it -p 3000:3000 apitestegolang`

Se nao tiver o docker instalado, baixe as dependencias com `go mod tidy` entre na pasta cmd `cd cmd` e rode a aplicação com `go run main.go`

Para rodar os testes rode o comando `go test -v ./...`

- Utilizar o storage local: adicionar a variável de ambiente LOCAL_STORAGE=true

### Dependências utilizadas

Para criaçao e manipulação de tokens:

- JWT-GO: github.com/dgrijalva/jwt-go v3.2.0

Para criação e geração de ids automatico.

- UUID: github.com/google/uuid v1.3.0

Para manipulação de rotas HTTP

- Gorilla Mux: github.com/gorilla/mux v1.8.0

Para criptografar senhas como hash e outros formatos.

- BCrypt: golang.org/x/crypto v0.0.0-20210921155107-089bfa567519

Para maniupular variáveis ambiente:

- GoDotEnv: github.com/joho/godotenv v1.4.0

#### Melhorias definidas

- Adicionar testes a camada HTTP
