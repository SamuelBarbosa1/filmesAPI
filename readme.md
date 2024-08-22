# REST API de Filmes

Esta é uma API REST para gerenciar filmes. A API utiliza Go para o backend principal e Node.js com Express como um intermediário para algumas das rotas. Você pode usar o Postman para testar os endpoints.

## Tecnologias Utilizadas

- **Go**: Para o backend principal.
- **Node.js**: Como servidor intermediário com Express.
- **Express**: Para criar as rotas e gerenciar requisições.
- **Postman**: Para testar os endpoints da API.

## URLs Disponíveis

### Node.js (Express)

Essas URLs são roteadas através do servidor Node.js, que atua como um intermediário entre o cliente (como Postman) e o backend Go.

#### Criar um Filme (POST)

- **URL**: `http://localhost:3000/filmes`
- **Descrição**: Adiciona um novo filme.
- **Exemplo de Corpo da Requisição**:
    ```json
    {
      "nome": "The Matrix",
      "ator": "Keanu Reeves",
      "descricao": "A hacker discovers the reality is a simulation.",
      "ano": 1999
    }
    ```

#### Obter Todos os Filmes (GET)

- **URL**: `http://localhost:3000/filmes`
- **Descrição**: Retorna uma lista de todos os filmes cadastrados.

#### Obter um Filme por ID (GET)

- **URL**: `http://localhost:3000/filmes/:id`
- **Descrição**: Retorna os detalhes de um filme específico pelo ID.
- **Exemplo**: `http://localhost:3000/filmes/1`

#### Atualizar um Filme por ID (PUT)

- **URL**: `http://localhost:3000/filmes/:id`
- **Descrição**: Atualiza os dados de um filme específico.
- **Exemplo de Corpo da Requisição**:
    ```json
    {
      "nome": "The Matrix Reloaded",
      "ator": "Keanu Reeves",
      "descricao": "The second installment in The Matrix trilogy.",
      "ano": 2003
    }
    ```

#### Deletar um Filme por ID (DELETE)

- **URL**: `http://localhost:3000/filmes/:id`
- **Descrição**: Remove um filme específico do cadastro.
- **Exemplo**: `http://localhost:3000/filmes/1`

### Backend Go

Essas URLs são acessadas diretamente no backend Go, sem passar pelo servidor Node.js.

#### Criar um Filme (POST)

- **URL**: `http://localhost:8080/filme`
- **Descrição**: Adiciona um novo filme (similar à rota do Node.js).

#### Obter Todos os Filmes (GET)

- **URL**: `http://localhost:8080/filmes`
- **Descrição**: Retorna uma lista de todos os filmes cadastrados.

#### Obter um Filme por ID (GET)

- **URL**: `http://localhost:8080/filme/:id`
- **Descrição**: Retorna os detalhes de um filme específico pelo ID.
- **Exemplo**: `http://localhost:8080/filme/1`

#### Atualizar um Filme por ID (PUT)

- **URL**: `http://localhost:8080/filme/:id`
- **Descrição**: Atualiza os dados de um filme específico.

#### Deletar um Filme por ID (DELETE)

- **URL**: `http://localhost:8080/filme/:id`
- **Descrição**: Remove um filme específico do cadastro.

## Testando com o Postman

Para testar os endpoints, importe o arquivo de coleção do Postman incluído neste repositório.

## Como Rodar

1. Clone o repositório.
2. Instale as dependências para o Node.js:
    ```bash
    cd nodejs
    npm install
    ```
3. Inicie o servidor Node.js:
    ```bash
    npm start
    ```
4. Inicie o servidor Go:
    ```bash
    cd backend
    go run main.go
    ```

Certifique-se de que ambos os servidores estão rodando antes de testar os endpoints.

## Contribuições

Sinta-se à vontade para contribuir com melhorias, abrir issues ou enviar pull requests.

## Licença

Este projeto é licenciado sob a [MIT License](LICENSE).
