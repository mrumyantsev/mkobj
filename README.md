# mkdir

**mkdir** is a lightweight CLI utility, which creates packages in Go. This is useful for developing general purpose types of applications: (micro/macro)servers, utilities or scripts. The goal is - to save developer's forces on routinly making of the same structure Go files.

## Installing

### Linux

```
sudo apt-get update
sudo apt-get install -y make
git clone https://github.com/mrumyantsev/mkobj.git
cd mkobj
make
sudo cp ./build/mkobj /usr/local/bin
```

## Usage

Create a new package directory and specify its full path. Example: `mkobj internal/Http-Server`. The last part of path will always tell the application to create a Go struct, named `HttpServer` (no hyphens), and the internal package name will be `httpserver` (no hyphens, lowercases), but the package directory and the Go file in it will be named accordingly: `http-server` and `http-server.go` (hyphens, lowercased).

```
mkobj internal/App
mkobj internal/Handler
mkobj internal/Service
mkobj internal/Repository
mkobj internal/Models
mkobj internal/Config
mkobj internal/Http-Server
mkobj internal/Http-Client
mkobj internal/Database
```
