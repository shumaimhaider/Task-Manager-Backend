# Task-Manager-Backend

# Install Go For Ubutnu

```sudo apt update```

```sudo apt upgrade```

```wget https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz```

```sudo tar -C /usr/local/ -xzf go1.21.0.linux-amd64.tar.gz```

# Set Path

```sudo nano $HOME/.profile```

Place the following path in file and save it

```export PATH=$PATH:/usr/local/go/bin```

Execute this command

```source ~/.profile```

To check go version

```go version```

# Initialize the Module

Initialize the Go module for your project using the following command:

```go mod init github.com/shumaimhaider/task_manager_api```

# Install Depedencies

FrontEnd
```make install-requirements```

Backend
```npm install```

# Configure Database credentials

Configure username, password and db name in [migration main](migrations/main.go) and [connect](connect/connect.go)

# DataBase Migrations 

Intialize the table for migrations in db

```make migrate-init```

Runs all available migrations

```make migrate-up```

Reverts last migration

```make migrate-down```

# Run
 
 For BackEnd
```make dev-local```

For FrontEnd
```npm start```
