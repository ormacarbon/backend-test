# 🛠️ Bvio Mono

Monorepo da aplicação **Bvio**, uma plataforma de educação gamificada com foco em competições, rankings e recompensas via sistema de indicações.

Este repositório contém tanto o **backend** (Go) quanto o **frontend** (Next.js), com orquestração via Docker e arquitetura orientada por **DDD (Domain-Driven Design)** no backend.

---

## ✨ Funcionalidades

- Cadastro de usuários com código de convite único
- Autenticação via JWT
- Atribuição de pontos por convites
- Ranking de usuários por pontuação
- Finalização da competição:
  - Envio de e-mails para os 10 primeiros colocados
  - Reset da pontuação de todos os usuários

---

## 🧠 Arquitetura e DDD

A arquitetura segue os princípios de Domain-Driven Design:

- **Domain Layer**: contém entidades, regras de negócio e interfaces dos repositórios
- **Application Layer**: orquestra os casos de uso
- **Infrastructure Layer**: implementações concretas de repositórios (DB, e-mail)
- **Interface Layer**: entrega da aplicação (ex: HTTP handlers com Gin)

### Benefícios

- Alta testabilidade
- Baixo acoplamento
- Clareza na separação de responsabilidades
- Facilita manutenção e extensão de funcionalidades

---

## 🐳 Docker

A aplicação é inteiramente dockerizada. Para subir o ambiente localmente:

```bash
docker-compose up --build
```

Serviços disponíveis:

- `backend`: aplicação Go na porta `${APP_PORT}`
- `frontend`: aplicação Next.js na porta `${FRONTEND_PORT}`
- `db`: banco PostgreSQL na porta `5432`

---

## 🔐 Autenticação

- Feita via **JWT**
- Após login, o token é retornado no corpo da resposta
- Rotas privadas requerem o cabeçalho:

```
Authorization: <token>
```

---

## 🥪 Testes

Os testes unitários são escritos com `testify` e `testify/mock`.

### Estrutura dos testes:

1. Setup de mocks
2. Execução do caso de uso
3. Verificações com `assert` e `AssertExpectations`

```bash
cd backend
go test ./...
```

---

## 🌐 Variáveis de Ambiente

As variáveis estão centralizadas em um único arquivo `.env`:

```env
# Backend
APP_PORT=8080

# Frontend
FRONTEND_PORT=3000
NEXT_PUBLIC_API_URL=http://localhost:8080

# Banco de dados
DB_NAME=bvio
DB_USER=postgres
DB_PASSWORD=secret
```

O Docker Compose injeta esse `.env` automaticamente nos serviços.

---

## 🚀 Deploy

### Backend

- Hospedado na [Render](https://render.com)

---

## 🤝 Contribuindo

Contribuições são bem-vindas! Sinta-se à vontade para abrir:

- Issues com dúvidas, bugs ou sugestões
- Pull Requests com melhorias ou novas funcionalidades

---

## 📬 Contato

Feito com 💙 por Cássius Queiroz Bessa.

Se tiver dúvidas ou quiser trocar ideias, abra uma issue ou entre em contato!

