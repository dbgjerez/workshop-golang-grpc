
// TODO: explain the objective of this workshop

# Installation

## protoc

**Download the binary to a temporal directory**

```bash
curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip -o /tmp/protoc.zip
```

> NOTE: This example uses the v21.12 version, which is the latest version today. You can change for another version in the following link, even though it's not necessary: https://github.com/protocolbuffers/protobuf/releases

**Unzip it**
```bash
unzip /tmp/protoc.zip -d /tmp/protoc/
```

**Install the binary**

```bash
mv /tmp/protoc/bin/protoc /usr/local/bin/protoc
```

**Verify the installation**

```bash
protoc --version
libprotoc 3.21.12
```

## protoc-gen-go

```zsh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
```

> NOTE: if you can execute ```protoc-gen-go``` you have to add it to your PATH: ```export GO_PATH=~/go``` and ```export PATH=$PATH:/$GO_PATH/bin```