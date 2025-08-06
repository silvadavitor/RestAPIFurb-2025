# RestAPIFurb

API REST desenvolvida para a prova de suficiência da disciplina de Programação Web II (FURB - 2025/2).  


## Tecnologias utilizadas

- **Go** — linguagem utilizada
- **Gin** — framework web para rotas REST
- **GORM** — ORM para persistência em banco relacional
- **PostgreSQL** — banco de dados utilizado via container
- **Docker / Docker Compose** — orquestração da API + banco
- **Swagger** — documentação da API
- **JWT** — autenticação via token

## Como executar o projeto com Docker

1. Clone o repositório:
```bash
git clone https://github.com/silvadavitor/RestAPIFurb-2025.git
cd RestAPIFurb-2025
```


2. Suba os containers com Docker Compose:
```bash
docker-compose up --build
```


## API Reference

#### Login

```http
  POST /RestAPIFurb/login
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `Usuario`      | `string` | **Required**. Nome do usuario |
 `Senha`         | `string` | **Required**. Senha do usuario |

Retorna o token JWT

#### Get all comandas (resumo)

```http
  GET /RestAPIFurb/comandas
```
Não requer autenticação. Retorna apenas dados de usuário.


#### Get comanda by ID

```http
  GET /RestAPIFurb/comandas/{id}

```

Não requer autenticação. Retorna os dados e os produtos da comanda.


#### Criar nova comanda

```http
  POST /RestAPIFurb/comandas
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `idUsuario` | `int` | **Required**. ID do usuário |
| `nomeUsuario` | `string` | **Required**. Nome do usuário |
| `telefoneUsuario` | `string` | **Required**. Telefone do usuário |
| `produtos	` | `array` | **Required**. Lista de produtos |

Não requer autenticação.



#### Atualizar comanda (produtos)
```http
  PUT /RestAPIFurb/comandas/{id}

```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `valor que deseja alterar`    | `string` |

Não requer autenticação.


#### Deletar comanda

```http
  DELETE /RestAPIFurb/comandas/{id}

```
| Header | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `Authorization` | `string` | **Required**. Bearer token JWT |

Requer autenticação.


#### Doc Swagger

```http
  GET /RestAPIFurb/swagger/index.html
```

## Estrutura utilizada
<img width="1280" height="720" alt="Image" src="https://github.com/user-attachments/assets/eaa7c526-0ab4-46cf-9766-fc321c5ee7b8" />
