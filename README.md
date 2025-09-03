# entrega-eda

## Como rodar o projeto

Este projeto utiliza Docker e Docker Compose para orquestrar todos os serviços necessários (Go apps, MySQL, Kafka, Zookeeper, Control Center).

### Pré-requisitos
- Docker instalado
- Docker Compose instalado

### Passos para rodar

1. Clone o repositório:
   ```bash
   git clone <url-do-repositorio>
   cd entrega-eda
   ```

2. Suba todos os serviços:
   ```bash
   docker-compose up -d --build
   ```
   O parâmetro `--build` garante que as imagens estejam atualizadas.

3. Para derrubar todos os serviços e remover volumes:
   ```bash
   docker-compose down --volumes --remove-orphans
   ```

### Serviços disponíveis
- **walletcore**: API principal de transações
- **balanceservice**: Serviço de atualização de saldo via Kafka
- **MySQL**: Banco de dados
- **Kafka/Zookeeper**: Mensageria
- **Control Center**: UI para monitorar Kafka

### Observações
- Os bancos são inicializados automaticamente via scripts SQL.
- O fluxo de eventos entre walletcore e balanceservice é feito via Kafka.

## Como utilizar os arquivos .http (API REST)

Os arquivos `.http` na pasta `api/` permitem testar e interagir com as APIs do projeto diretamente pelo VS Code (usando a extensão REST Client) ou outras ferramentas que suportem requisições HTTP.

### Passo a passo:

1. **Consultar clientes e contas**
   - Use o arquivo `client.http` para consultar, criar ou listar clientes.
   - Use o arquivo `balanceservice.http` para consultar o saldo de uma conta.

2. **Realizar uma transação**
   - Use o arquivo `client.http` ou `balanceservice.http` para criar uma transação entre contas.
   - Exemplo de requisição:
     ```http
     POST http://localhost:3000/transactions
     Content-Type: application/json

     {
       "account_id_from": "<id_origem>",
       "account_id_to": "<id_destino>",
       "amount": 100
     }
     ```
   - Isso irá disparar um evento no Kafka e atualizar os saldos.

3. **Consultar saldo após transação**
   - Use o arquivo `balanceservice.http` para consultar o saldo das contas envolvidas:
     ```http
     GET http://localhost:3003/balances/<id_conta>
     ```
   - O saldo será atualizado automaticamente pelo serviço de eventos.