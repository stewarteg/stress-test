# Stress Test CLI

Sistema CLI em Go para realizar testes de carga.

## Parâmetros

- `--url`: URL do serviço a ser testado.
- `--requests`: Número total de requests.
- `--concurrency`: Número de chamadas simultâneas.

### Local
rode o build:
   go build -o stress-test

e depois esse aq:
   ./stress-test --url=http://google.com --requests=1000 --concurrency=10



