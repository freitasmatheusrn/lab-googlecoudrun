# Weather API - Consulta de Temperatura por CEP

Uma API simples em Go que permite consultar informações de temperatura com base no CEP (Código de Endereçamento Postal) brasileiro.

## Funcionalidades

- Consulta de temperatura por CEP
- Servidor HTTP simples e eficiente
- Validação de formato de CEP


## Como executar

Acesse o servidor na url: 
  ### https://googlecloud-3wgw77geaq-uc.a.run.app/


## Endpoints

### GET `https://googlecloud-3wgw77geaq-uc.a.run.app/weather_from/{cep}`

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

