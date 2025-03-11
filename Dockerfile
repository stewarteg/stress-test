# Usar uma imagem base do Go
FROM golang:1.20

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos do projeto para o container
COPY . .

# Baixar as dependências e compilar o binário
RUN go mod tidy
RUN go build -o stress-test .

# Comando padrão para executar a aplicação
CMD ["./stress-test"]