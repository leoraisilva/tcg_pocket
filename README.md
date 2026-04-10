# 🃏 TCG Pocket API

API desenvolvida em Go para gerenciamento de cartas do TCG Pocket, implementando um CRUD completo seguindo boas práticas de arquitetura em camadas.

# 🚀 Sobre o projeto

O TCG Pocket API é um serviço backend responsável por gerenciar cartas do jogo, permitindo operações de criação, leitura, atualização e remoção (CRUD).

O projeto foi estruturado com foco em organização, escalabilidade e separação de responsabilidades.

# 🧱 Arquitetura

O projeto segue uma abordagem inspirada em Clean Architecture, separando responsabilidades em diferentes camadas:

    ├── cmd/         # Ponto de entrada da aplicação
    ├── controller/  # Camada de entrada (HTTP handlers)
    ├── usecase/     # Regras de negócio
    ├── repository/  # Acesso a dados (DB ou memória)
    ├── model/       # Estruturas de dados (entidades)
    ├── resource/    # DTOs / responses / requests
    └── helper/      # Funções utilitárias

# 📌 Descrição das camadas

<ol>
    <li>cmd → inicializa o servidor e configura dependências</li>
    <li>controller → recebe requisições HTTP e chama os use cases</li>
    <li>repository → abstração de persistência de dados</li>
    <li>model → definição das entidades principais</li>
    <li>resource → objetos de entrada/saída da API</li>
    <li>helper → funções auxiliares conexão ao banco</li>
</ol>

# ✨ Funcionalidades
### ➕ Criar carta
### 📄 Listar cartas
### 🔍 Buscar carta por ID
### ✏️ Atualizar carta
### ❌ Remover carta

# Tecnologias utilizadas
<ol>
    <ul>Go (Golang)</ul>
    <ul>net/http (ou framework, se aplicável)</ul>
    <ul>JSON para comunicação de dados</ul>
</ol>
