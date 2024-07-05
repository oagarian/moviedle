[![en](https://img.shields.io/badge/lang-en-blue.svg)](https://github.com/oagarian/moviedle/blob/master/README.md)

# Moviedle

## O que é 
Moviedle é um projeto inspirado nos jogos "dle", como [LoLdle](https://loldle.net/) e [Pokédle](https://pokedle.net/), com foco em filmes. O objetivo é proporcionar uma experiência divertida e interativa para os amantes do cinema.

## Dependências
Para rodar o projeto, é necessário ter instalado:
1. [GoLang](https://golang.org/)
2. [Docker](https://www.docker.com/)
3. [Pacote migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4)

## Como rodar o projeto
1. Clone o repositório:
    ```sh
    git clone https://github.com/oagarian/moviedle
    ```
2. Baixe as dependências do projeto:
    ```sh
    go mod tidy
    ```
3. Inicie o banco de dados utilizando o script `execute.sh` na raiz do projeto:
    ```sh
    chmod +x ./execute.sh
    ./execute.sh -environment -database
    ```
4. Execute os testes:
    ```sh
    go test ./...
    ```
5. Inicie a aplicação:
    ```sh
    go run src/apps/api/main.go
    ```

Para testar as rotas, acesse os endpoints descritos na seção de Rotas.

## Rotas

### Documentação da API
`/api/docs/index.html`  
Acesse para ver a documentação completa da API.

### Endpoints de Filmes
- **Criar/Adicionar Filme**  
  `POST /api/movies`  
  Adiciona um novo filme ao banco de dados.

- **Listar Todos os Filmes**  
  `GET /api/movies`  
  Lista todos os filmes cadastrados no banco de dados.

- **Detalhes de um Filme**  
  `GET /api/movies/{id}`  
  Exibe os detalhes de um filme específico pelo ID.

- **Listar Filmes por Filtros**  
  `GET /api/movies?{filters}`  
  Lista filmes de acordo com os filtros especificados.

- **Editar um Filme**  
  `PUT /api/movies/{id}`  
  Edita as informações de um filme específico pelo ID.

- **Apagar um Filme**  
  `DELETE /api/movies/{id}`  
  Remove um filme específico pelo ID do banco de dados.

## Contribuição
Contribuições são bem-vindas! Sinta-se à vontade para abrir issues e pull requests no repositório.

## Licença
Este projeto está licenciado sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---
