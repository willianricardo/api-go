# API de Teste em Golang

Esta é uma API de teste em Golang que oferece funcionalidades básicas para gerenciar clientes (customers), produtos (products) e pedidos (orders).

## Pré-requisitos

Antes de começar, certifique-se de ter o Go instalado em sua máquina. Você pode fazer o download do Go em [https://golang.org/dl/](https://golang.org/dl/).

## Configuração

1. Clone este repositório para a sua máquina:

   ```bash
   git clone https://github.com/willianricardo/api-go.git
   ```

2. Navegue até o diretório do projeto:

   ```bash
   cd api-go
   ```

3. Instale as dependências:

   ```bash
   go mod tidy
   ```

4. Inicie o servidor:

   ```bash
   go run main.go
   ```

A API estará disponível em [http://localhost:8080](http://localhost:8080).

## Rotas

A API oferece as seguintes rotas:

### Clientes (Customers)

- **GET /customers**: Obtém a lista de clientes.
- **GET /customers/{id}**: Obtém detalhes de um cliente específico por ID.
- **POST /customers**: Cria um novo cliente.
- **PUT /customers/{id}**: Atualiza os detalhes de um cliente por ID.
- **DELETE /customers/{id}**: Exclui um cliente por ID.

### Produtos (Products)

- **GET /products**: Obtém a lista de produtos.
- **GET /products/{id}**: Obtém detalhes de um produto específico por ID.
- **POST /products**: Cria um novo produto.
- **PUT /products/{id}**: Atualiza os detalhes de um produto por ID.
- **DELETE /products/{id}**: Exclui um produto por ID.

### Pedidos (Orders)

- **GET /orders**: Obtém a lista de pedidos.
- **GET /orders/{id}**: Obtém detalhes de um pedido específico por ID.
- **POST /orders**: Cria um novo pedido.
- **PUT /orders/{id}**: Atualiza os detalhes de um pedido por ID.
- **DELETE /orders/{id}**: Exclui um pedido por ID.
