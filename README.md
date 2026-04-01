# Weather API - Consulta de Temperatura por CEP

Uma API simples em Go que permite consultar informações de temperatura com base no CEP (Código de Endereçamento Postal) brasileiro.

## Funcionalidades

- Consulta de temperatura por CEP
- Servidor HTTP simples e eficiente
- Validação de formato de CEP


## Como executar

### Localmente via Docker

```bash
# Build da imagem
docker build -t weather-api .

# Executar o container
docker run -p 8080:8080 weather-api
```

A API ficará disponível em `http://localhost:8080`.

### Cloud Run

Faça um get request para: 
  ### https://googlecloud-3wgw77geaq-uc.a.run.app/wheater_from/{cep}


## Endpoints

### GET `/wheater_from/{cep}`

Retorna as informações de temperatura para o CEP fornecido.

**Parâmetros:**
- `{cep}`: Código de Endereçamento Postal com exatamente 8 caracteres numéricos

**Exemplo de uso:**
```
GET https://googlecloud-3wgw77geaq-uc.a.run.app/weather_from/01310100
```

**Formato do CEP:**
- Deve conter exatamente 8 dígitos numéricos
- Exemplos válidos: `01310100, 01310-100`
- Exemplos inválidos: `0131010`, `0131010a`

