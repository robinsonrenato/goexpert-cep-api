# Desafio Go Expert - Módulo Multithreading

## Proposta

Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/" + cep + ".json

http://viacep.com.br/ws/" + cep + "/json/

Os requisitos para este desafio são:

- Aceitar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.

## Como rodar

Clone o repositório e entre dentro do diretório, após, rode o seguinte comando:

> CEP é o argumento passado para a aplicação, deve conter 8 dígitos no seguinte formato: 03947000.

```
go run cmd/main.go -cep=03947000
```

### Possíveis retornos

Sucesso:

```json
{
    "api_name": "Via Cep Service",
    "bairro": "Jardim Nove de Julho",
    "logradouro": "Rua Professor Antônio Sampaio Dória",
    "cidade": "São Paulo",
    "uf": "SP"
}
```

Erro no retorno da API:

```json
{
    "api_name": "Api Cep Service",
    "message": "Blocked by flood cdn"
}
```

Tempo limite de execução esgotado

```json
timeout exceeded
```

CEP inválido

```json
{
    "api_name": "Via Cep Service",
    "message": "CEP incorreto"
}
```
