# Cultural - Aplicação React

Esta é a versão React da aplicação Cultural, convertida das páginas HTML originais seguindo as melhores práticas e padrões do React.

## 🏗️ Estrutura do Projeto

```
cultural/
├── src/
│   ├── components/          # Componentes reutilizáveis
│   │   ├── Header.jsx      # Cabeçalho da aplicação
│   │   ├── SearchBar.jsx   # Barra de busca
│   │   ├── Card.jsx        # Card para eventos/atrações
│   │   ├── UserTypeSelection.jsx # Seleção de tipo de usuário
│   │   └── AppRouter.jsx   # Roteamento principal
│   ├── pages/              # Páginas da aplicação
│   │   ├── Home.jsx        # Página inicial
│   │   ├── Login.jsx       # Página de login
│   │   ├── Register.jsx    # Página de registro
│   │   ├── Profile.jsx     # Perfil do usuário
│   │   ├── Favorites.jsx   # Favoritos do usuário
│   │   ├── ChangePassword.jsx # Alterar senha
│   │   └── PasswordRecovery.jsx # Recuperação de senha
│   ├── contexts/           # Contextos React
│   │   └── UserContext.jsx # Contexto do usuário
│   ├── hooks/              # Hooks personalizados
│   │   ├── useFavorites.js # Hook para favoritos
│   │   └── useSearch.js    # Hook para busca
│   ├── utils/              # Utilitários
│   │   └── constants.js    # Constantes da aplicação
│   ├── App.jsx             # Componente principal
│   └── main.jsx            # Ponto de entrada
├── public/                 # Arquivos estáticos
│   ├── images/            # Imagens da aplicação
│   └── logo6.png          # Logo principal
└── package.json           # Dependências e scripts
```

## 🚀 Funcionalidades Implementadas

### ✅ Páginas Convertidas
- **Página Inicial**: Seleção de tipo de usuário
- **Home**: Dashboard com eventos e atrações
- **Login**: Autenticação de usuários
- **Registro**: Criação de novas contas
- **Perfil**: Visualização e edição de dados do usuário
- **Favoritos**: Lista de itens favoritados
- **Alterar Senha**: Mudança de senha
- **Recuperação de Senha**: Recuperação via email

### ✅ Componentes Reutilizáveis
- **Header**: Cabeçalho com navegação e ações
- **SearchBar**: Barra de busca funcional
- **Card**: Card para exibir eventos e atrações
- **UserTypeSelection**: Seleção de tipo de usuário

### ✅ Sistema de Autenticação
- Contexto global para estado do usuário
- Proteção de rotas
- Persistência no localStorage
- Diferentes tipos de usuário (comum/organizador)

### ✅ Funcionalidades Avançadas
- Sistema de favoritos com localStorage
- Busca e filtros em tempo real
- Navegação entre páginas
- Estados de carregamento
- Tratamento de erros
- Validação de formulários

## 🛠️ Tecnologias Utilizadas

- **React 19.1.1** - Biblioteca principal
- **React Router DOM** - Roteamento
- **Vite** - Build tool e dev server
- **ESLint** - Linting de código
- **CSS Modules** - Estilização

## 📱 Padrões React Implementados

### 1. **Componentes Funcionais**
- Uso de hooks (useState, useEffect, useContext)
- Props para comunicação entre componentes
- Componentes puros quando possível

### 2. **Gerenciamento de Estado**
- Context API para estado global
- useState para estado local
- localStorage para persistência

### 3. **Roteamento**
- React Router para navegação
- Rotas protegidas
- Redirecionamentos condicionais

### 4. **Hooks Personalizados**
- `useFavorites`: Gerenciamento de favoritos
- `useSearch`: Funcionalidade de busca e filtros
- Reutilização de lógica entre componentes

### 5. **Estrutura de Pastas**
- Separação clara de responsabilidades
- Componentes reutilizáveis
- Páginas específicas
- Utilitários e constantes

## 🎯 Como Usar

### Instalação
```bash
cd cultural
npm install
```

### Desenvolvimento
```bash
npm run dev
```

### Build
```bash
npm run build
```

### Preview
```bash
npm run preview
```

## 🔐 Credenciais de Teste

Para testar a aplicação:

**Login:**
- Email: `teste@teste.com`
- Senha: `123456`

## 🚧 Próximos Passos

### Funcionalidades a Implementar
- [ ] Páginas detalhadas de eventos
- [ ] Páginas detalhadas de atrações
- [ ] Sistema de notificações
- [ ] Upload de imagens
- [ ] Integração com API backend
- [ ] Testes unitários
- [ ] Testes de integração

### Melhorias Técnicas
- [ ] Implementar TypeScript
- [ ] Adicionar Storybook para componentes
- [ ] Implementar PWA
- [ ] Otimização de performance
- [ ] Internacionalização (i18n)

## 📝 Notas de Desenvolvimento

### Conversão HTML → React
- Formulários HTML convertidos para componentes controlados
- Eventos inline convertidos para handlers React
- Classes CSS convertidas para className
- Estrutura de dados mockada para demonstração

### Responsividade
- Design mobile-first mantido
- Componentes adaptáveis
- CSS flexível para diferentes tamanhos de tela

### Acessibilidade
- Labels apropriados para formulários
- Alt text para imagens
- Navegação por teclado
- Estrutura semântica HTML

## 🤝 Contribuição

Para contribuir com o projeto:

1. Fork o repositório
2. Crie uma branch para sua feature
3. Implemente as mudanças
4. Adicione testes se aplicável
5. Faça commit das mudanças
6. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.

