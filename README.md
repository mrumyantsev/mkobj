# Mkdir

**Mkdir** is a lightweight CLI utility, which creates packages in Go. This is helpful for broad purpose application developing: (micro/macro)REST-servers, CLI utilities or scripts. The goal is - to save developer's strength on routinly making of the same structure Go files.

## Installing

### Linux

Run the commands in the Terminal one by one. Omit unsuitable.

```
sudo apt-get update
sudo apt-get install -y make
git clone https://github.com/mrumyantsev/mkobj.git
cd mkobj
make
sudo mv ./build/mkobj /usr/local/bin
```

## Usage

### Creating an Object

Create a new package directory and specify its full path. Example: `mkobj internal/Http-Server`. The last part of path will always tell the application to create a Go struct, named `HttpServer` (no hyphens), and the internal package name will be `httpserver` (no hyphens, lowercases), but the package directory and the Go file in it will be named accordingly: `http-server` and `http-server.go` (hyphens, lowercased).

**Note:** every object will always be created (or placed) by lowercase path, no matter how it were specified in CLI parameter.

```
mkobj internal/App
mkobj internal/Handler
mkobj internal/Service
mkobj internal/Repository
mkobj internal/Repository/PostgresRepository
mkobj internal/Repository/MySqlRepository
mkobj internal/Repository/MongoDbRepository
mkobj internal/Models
mkobj internal/Config
mkobj internal/Http-Server
mkobj internal/Http-Client
mkobj internal/Database
```
