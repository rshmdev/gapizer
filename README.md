
# GAPIzer

**Um gerador de APIs automatizado, simples e poderoso.**

O **GAPIzer** é uma ferramenta CLI que permite gerar APIs completas a partir de um arquivo de configuração YAML. Ele cria a estrutura do projeto, incluindo rotas, handlers, middlewares (como autenticação e logging), documentação (Swagger) e muito mais. Ideal para desenvolvedores que desejam agilidade no desenvolvimento de APIs.

---

## 🚀 Instalação

### Pré-requisitos

1. **Golang**: Certifique-se de que o Go está instalado.  
   - Para instalar: [https://go.dev/dl/](https://go.dev/dl/)
   - Para verificar:
     ```bash
     go version
     ```

---

### 1. Instalação via npm

Instale o GAPIzer globalmente usando o npm:

```bash
npm install -g gapizer
```

### 2. Usando o Instalador para Windows

1. Acesse a página de [Releases](https://github.com/rshmdev/gapizer/releases).
2. Baixe o instalador mais recente::
 - **Windows**: `GAPIzer-Installer.exe`

#### **Instalar no Windows**

  1. Execute o instalador baixado `(GAPIzer-Installer.exe)`.
  2. Siga as instruções exibidas na tela.
  3. Após a instalação, o comando gapizer será configurado automaticamente no PATH do sistema.

---

### 3. Instalação Local (Compilando o Código)

1. Clone o repositório:
   ```bash
   git clone https://github.com/rshmdev/gapizer.git
   cd gapizer
   ```

2. Instale a CLI globalmente:
   ```bash
   go install
   ```

3. Verifique se a CLI foi instalada com sucesso:
   ```bash
   gapizer --help
   ```

---

## 🛠️ Como Usar

### Comando Básico

```bash
gapizer --config <arquivo_yaml> --output <diretorio_destino>
```

### Exemplo Rápido

1. Crie um arquivo `example.yml` com a seguinte configuração:
   ```yaml
   app_name: MyAwesomeAPI
   port: 8080
   database:
     type: sqlite
     name: ./data.db

   logging:
     enabled: true
     output: console # Pode ser "console" ou "file"
     file_path: logs/server.log

   authentication:
     type: jwt
     secret: mysupersecretkey
     token_expiration_minutes: 60

   endpoints:
     - name: /users
       method: GET
       protected: true
       response:
         id: int
         name: string
     - name: /products
       method: POST
       request:
         name: string
         price: float
       response:
         id: int
         name: string
         price: float
   ```

2. Gere a API:
   ```bash
   gapizer --config example.yml --output ./generated
   ```

3. Estrutura gerada:
   ```plaintext
   generated/
   ├── handlers/
   │   ├── users.go
   │   └── products.go
   ├── routes/
   │   ├── users.go
   │   └── products.go
   ├── middleware/
   │   ├── auth.go
   │   └── logging.go
   ├── database/
   │   └── database.go
   ├── docs/
   │   └── swagger.yaml
   ├── main.go
   ├── go.mod
   └── swagger-ui/
   ```

4. Rode o servidor:
   ```bash
   cd generated
   go run main.go
   ```

5. Teste os endpoints gerados:
   - **GET /users** (protegido, requer JWT)
   - **POST /products**

---

## 📄 Configuração do YAML

O arquivo YAML define os endpoints e outras configurações da API. Aqui está um exemplo completo:

```yaml
app_name: MyAwesomeAPI           # Nome da aplicação
port: 8080                       # Porta do servidor
database:                        # Configuração do banco de dados
  type: sqlite                   # Tipos suportados: sqlite, mysql, postgresql
  name: ./data.db

logging:                         # Configuração de logs
  enabled: true                  # Habilitar ou desabilitar logs
  output: console                # "console" ou "file"
  file_path: logs/server.log     # Caminho para arquivo de logs (se output = file)

authentication:                  # Configuração de autenticação
  type: jwt                      # Tipo de autenticação (atualmente suporta apenas JWT)
  secret: mysupersecretkey       # Chave secreta para assinar tokens JWT
  token_expiration_minutes: 60   # Expiração do token em minutos

endpoints:                       # Lista de endpoints
  - name: /users                 # Caminho do endpoint
    method: GET                  # Método HTTP
    protected: true              # Define se o endpoint é protegido (requer autenticação)
    response:                    # Estrutura da resposta
      id: int
      name: string
  - name: /products
    method: POST
    request:                     # Estrutura da requisição
      name: string
      price: float
    response:                    # Estrutura da resposta
      id: int
      name: string
      price: float
```

---

## 📜 Documentação Automática (Swagger)

O **GAPIzer** gera automaticamente a documentação Swagger no arquivo `swagger.yaml`. Para visualizar:

1. Acesse o [Swagger Editor](https://editor.swagger.io/).
2. Faça upload do arquivo `swagger.yaml` gerado na pasta `docs`.

---

## 📋 Comandos Disponíveis

### Ajuda

Exibe os comandos disponíveis:
```bash
gapizer --help
```

### Gerar API

Gera a estrutura da API:
```bash
gapizer --config <arquivo_yaml> --output <diretorio_destino>
```
- **`--config`**: Caminho para o arquivo YAML de configuração.
- **`--output`**: Diretório onde a API será gerada.

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Para contribuir, siga estas etapas:

1. Faça um fork do repositório.
2. Crie um branch para a sua feature ou correção:
   ```bash
   git checkout -b minha-feature
   ```
3. Faça commit das suas alterações:
   ```bash
   git commit -m "Descrição clara da alteração"
   ```
4. Envie o branch:
   ```bash
   git push origin minha-feature
   ```
5. Abra um Pull Request.

---

## 🛣️ Roadmap

Funcionalidades planejadas para o futuro:
- Suporte a Swagger UI diretamente no servidor gerado.
- Geração de clientes em diferentes linguagens (TypeScript, Python, etc.).
- Suporte a banco de dados dinamicamente mapeado.
- Logs estruturados em formato JSON para integração com ferramentas de monitoramento.

---

## 📝 Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.
