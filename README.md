# Task-Manager-Backend

# Install Go For Ubutnu

```sudo apt update```

```sudo apt upgrade```

```wget https://dl.google.com/go/go1.21.0.linux-amd64.tar.gz```

```sudo tar -C /usr/local/ -xzf go1.21.0.linux-amd64.tar.gz```

# Set Path

```sudo nano $HOME/.profile```

```export PATH=$PATH:/usr/local/go/bin```

```source ~/.profile```

```go version```

# Initialize the Module

Initialize the Go module for your project using the following command:

```go mod init github.com/shumaimhaider/task_manager_api```

```go mod tidy```

# Installing and Setting up Gin

``` go get -u github.com/gin-gonic/gin```

# Run

```go run main.go```