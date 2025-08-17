# Cultural :)

## Configurações iniciais:

> [!NOTE]
> É preciso fazer o clone do repositório primeiro e ter as seguintes instalaçãos em sua máquina local:
> 1. [go](https://go.dev/);
> 2. node.js


**Backend:**

```
go mod tidy
``` 

**Frontend:**

```
sudo apt-get node
npm install react-router-dom
``` 
> [!TIP]
> Confira se é a versão mais recente do node.

**Banco de dados:**

1. Instale o mysql

```
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=senha -p 3306:3306 -d mysql:latest
sudo apt install mysql-server
go get github.com/go-sql-driver/mysql
sudo apt install mysql-server 
```

> [!TIP]
> Se a instalação via docker não funcionar, faça direto do [site](https://dev.mysql.com/downloads/mysql/).

2. Inicie o _mysql_ como **root**
``` 
sudo mysql
``` 

3. Crie o banco de dados, o usuário e a senha:
```
CREATE DATABASE POCII;
CREATE USER 'claracosta86'@'localhost' IDENTIFIED BY 'sua_senha';
GRANT ALL PRIVILEGES ON meuprojeto.* TO 'claracosta86'@'localhost';
FLUSH PRIVILEGES;
```

> [!WARNING]
> Não troque nenhum dos dois, para não quebrar o código;
> Se fizer questão, em `main.go` altere a variável `dns` de acordo.

4. Entre em seu usuário:
```
mysql -u claracosta86 -p
```

5. Entre no banco de dados criado;
```
USE POCII;
```

6. Crie suas tabelas
```
POCII < back/schema/users_table.sql
POCII < back/schema/user_favorite_table.sql
POCII < back/schema/events_table.sql
POCII < back/schema/tourist_attractions_table.sql
```

## Rodando o projeto: 

- Em terminais diferentes rode os seguintes comandos:

```
cd back
go run cmd/api/main.go
```

```
cd cultural
npm run dev
```

> [!IMPORTANT]
> Garanta que não há nada rodando nas portas `5173` e `8080`. 
