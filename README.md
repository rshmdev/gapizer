
# GAPIzer

**Um gerador de APIs automatizado, simples e poderoso.**

O **GAPIzer** é uma ferramenta CLI que permite gerar APIs completas a partir de um arquivo de configuração YAML. Ele cria a estrutura do projeto, incluindo rotas, handlers, documentação (Swagger) e muito mais. Ideal para desenvolvedores que desejam agilidade no desenvolvimento de APIs.

---

## 🚀 Instalação

### Pré-requisitos
1. **Golang**: Certifique-se de que o Go está instalado.  
   - Para instalar: [https://go.dev/dl/](https://go.dev/dl/)
   - Para verificar:
     ```bash
     go version
     ```

2. **Node.js (opcional)**: Caso você queira disponibilizar a CLI pelo `npm`.

---

### Instalação Local
1. Clone o repositório:
   ```bash
   git clone https://github.com/seu-usuario/gapizer.git
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
   endpoints:
     - name: /users
       method: GET
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
   - **GET /users**
   - **POST /products**

---

## 📄 Configuração do YAML

O arquivo YAML define os endpoints e outras configurações da API. Aqui está um exemplo completo:

```yaml
app_name: MyAPI           # Nome da aplicação
port: 8080                # Porta do servidor
endpoints:                # Lista de endpoints
  - name: /users          # Nome do endpoint
    method: GET           # Método HTTP
    response:             # Resposta da API
      id: int
      name: string
  - name: /products        # Endpoint para criar produtos
    method: POST
    request:              # Dados de entrada
      name: string
      price: float
    response:             # Resposta da API
      id: int
      name: string
      price: float
```

### Configurações Suportadas
- **`app_name`**: Nome da aplicação.
- **`port`**: Porta na qual o servidor será executado.
- **`endpoints`**: Lista de endpoints, com:
  - **`name`**: Caminho do endpoint (ex.: `/users`).
  - **`method`**: Método HTTP (GET, POST, PUT, DELETE).
  - **`request`** (opcional): Estrutura do payload esperado na requisição.
  - **`response`**: Estrutura da resposta enviada pela API.

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
- Suporte nativo ao Swagger UI.
- Geração de clientes em diferentes linguagens (TypeScript, Python, etc.).
- Integração com bancos de dados para geração dinâmica de APIs.

---

## 📝 Licença

Este projeto é licenciado sob a licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.
