# Weather API - Consulta de Temperatura por CEP

Uma API simples em Go que permite consultar informações de temperatura com base no CEP (Código de Endereçamento Postal) brasileiro.

## Funcionalidades

- Consulta de temperatura por CEP
- Servidor HTTP simples e eficiente
- Validação de formato de CEP

## Pré-requisitos

- Go 1.19 ou superior instalado
- Conexão com a internet para consultas de temperatura

## Como executar

Para iniciar o servidor, execute o seguinte comando no terminal:

```bash
go run main.go
```

O servidor será iniciado na porta **8080** e ficará aguardando requisições.

## Endpoints

### GET `/weather_from/{cep}`

Retorna as informações de temperatura para o CEP fornecido.

**Parâmetros:**
- `{cep}`: Código de Endereçamento Postal com exatamente 8 caracteres numéricos

**Exemplo de uso:**
```
GET http://localhost:8080/weather_from/01310100
```

**Formato do CEP:**
- Deve conter exatamente 8 dígitos numéricos
- Exemplos válidos: `01310100, 01310-100`
- Exemplos inválidos: `0131010`, `0131010a`

