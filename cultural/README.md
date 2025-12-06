# Arquitetura Frontend (React + Vite)

Este documento descreve a arquitetura da aplicação frontend localizada na pasta `cultural`.

## Visão Geral da Arquitetura

A aplicação é construída utilizando **React** com **Vite** como ferramenta de build. Ela segue uma arquitetura baseada em componentes com uma clara separação entre "Páginas" (Views) e "Componentes" (Blocos de UI), juntamente com a Context API para gerenciamento de estado.

```
cultural/
├── src/
│   ├── assets/          # Assets estáticos (imagens, ícones)
│   ├── components/      # Componentes de UI reutilizáveis
│   │   ├── Layout/      # Componentes estruturais (Header, Footer)
│   │   ├── [Modals]/    # Componentes modais de uso repetido no código
│   │   └── ProtectedRoute.jsx # Componente de guarda de rota
│   ├── contexts/        # React Contexts (Estado Global)
│   │   └── UserContext.jsx
│   ├── hooks/           # Hooks Customizados do React
│   │   └── useNotifications.js
│   ├── pages/           # Componentes de Página (Alvos de Rota)
│   │   ├── styles/      # Arquivos CSS específicos da página
│   │   └── [Page].jsx   # Lógica individual da página
│   ├── styles/          # Estilos globais
│   │   ├── App.css
│   │   ├── index.css
│   │   └── modal.css    # Estilos unificados de Modal
│   ├── App.jsx          # Configuração principal do Roteador
│   └── main.jsx         # Ponto de Entrada da Aplicação
├── public/              # Arquivos estáticos públicos
├── index.html           # Ponto de Entrada HTML
└── vite.config.js       # Configuração do Vite
```

## Responsabilidades das Camadas

### 1. Páginas (`src/pages/`)
- **Responsabilidade**: Representar visualizações/telas completas acessíveis via Rotas.
- **Conteúdo**: Buscar dados, gerenciar estado específico da página e compor Componentes.
- **Nomeação**: Sufixo `Page.jsx` (ex: `HomePage.jsx`, `LoginPage.jsx`).
- **Estilização**: Estilos específicos localizados em `src/pages/styles/`.

### 2. Componentes (`src/components/`)
- **Responsabilidade**: Elementos de UI reutilizáveis usados em múltiplas páginas.
- **Tipos**:
    - **Layout**: `Header`, `Footer`.
    - **Modals**: `NotificationModal`, `ConfirmModal`, etc.
    - **Guards**: `ProtectedRoute` para verificação de autenticação.
- **Isolamento**: Devem receber dados via `props` sempre que possível.

### 3. Contextos (`src/contexts/`)
- **Responsabilidade**: Gerenciamento de estado global.
- **Principais Contextos**:
    - `UserContext`: Gerencia o estado de autenticação do usuário e dados do perfil.

### 4. Hooks (`src/hooks/`)
- **Responsabilidade**: Encapsular lógica reutilizável.
- **Exemplo**: `useNotifications` lida com a busca e gerenciamento de notificações do usuário.

### 5. Estilos (`src/styles/`)
- **Responsabilidade**: Estilização global e tokens de design compartilhados.
- **Principais Arquivos**:
    - `modal.css`: Estilos centralizados para todas as janelas modais para garantir consistência.
    - `index.css`: Resets base e variáveis globais.

## Principais Decisões Arquiteturais

### Páginas vs Componentes
Separamos estritamente **Páginas** (que têm conhecimento do Roteamento e contextos de negócio específicos) de **Componentes** (que são genéricos e reutilizáveis). Isso torna a base de código mais fácil de navegar e manter.

### Estilização Centralizada de Modais
Para evitar duplicação de CSS e inconsistência, todos os Modais compartilham uma folha de estilo comum (`src/styles/modal.css`). Isso garante que `NotificationModal`, `ConfirmModal` e outros tenham aparência e comportamento idênticos.

### Roteamento
- **Biblioteca**: `react-router-dom`
- **Estrutura**: Definida em `App.jsx`.
- **Proteção**: Rotas privadas são envolvidas em `<ProtectedRoute />` que verifica o `UserContext`.

## Começando

1. **Instalar Dependências**:
   ```bash
   npm install
   ```

2. **Rodar Servidor de Desenvolvimento**:
   ```bash
   npm run dev
   ```

3. **Build para Produção**:
   ```bash
   npm run build
   ```
