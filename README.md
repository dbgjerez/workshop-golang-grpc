// TODO: explain the objective of this workshop

# Installation

## protoc

**Download the binary to a temporal directory**

```bash
curl -L https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip -o /tmp/protoc.zip
```

> NOTE: We'll use the "v21.12" version of protoc, which is the latest today. You can change it for another version that you can find in the following link, even though it's not necessary: https://github.com/protocolbuffers/protobuf/releases

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

## protoc-gen-go and protoc-gen-go-grpc

```zsh
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

> NOTE: if you can execute ```protoc-gen-go``` you have to add it to your PATH: ```export GO_PATH=~/go``` and ```export PATH=$PATH:/$GO_PATH/bin```

# Generate protobufs

A significant advantage of using protobufs is the capacity to auto-generate the code in your preferred language, even in different languages. 

In this example, we're going to create a ```.proto``` file with the definition of the entity Comment. 

Once we've created the file, we'll generate the source code. In this case, I have chosen the golang plugin.

```zsh
protoc src/domain/*.proto \
    --go_out=. \
    --go_opt=paths=source_relative \
    --go-grpc_out=. \
    --go-grpc_opt=paths=source_relative
```

> NOTE: It's very important to install the preferred language plugin, in the step before I installed ```protoc-gen-go```, but you can install another.

After the plugin execution, we can inspect our code inside the ```src/domain``` folder.

# Create a project

As a good practice, we shouldn't modify the autogenerated code, but we can import and override it.

We are going to create a project, called comment, to use our gRPC server:

```bash
go mod init comment
```

It's important to download the dependencies that we need:

```bash
go mod tidy
go get google.golang.org/grpc
go get google.golang.org/grpc/reflection
```

Once, we have our project we're going to create a ```main.go``` file and implement our server. 

The main file contains some important points. The first one is the struct to override the auto-generated code. This struct has been called sever and looks like that: 

```go
type server struct {
    domain.UnimplementedCommentServiceServer
}
```

As we can see, the struct uses the unimplemented struct. Now, we are going to create the server in the main function, registering the server struct:

```go
func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    domain.RegisterCommentServiceServer(s, &server{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

Our main function implements some aspects. We see how to listen to a port, this port will be used by our application. Later, a new grpc server is started and the comment service is registered. 

The line with ```reflection.Register(s)``` enables the possibility to call for know the different functions and endpoints exposes our application. We will see it later. 

The last part starts the server and returns an error it isn't possible. 

# Play with the application

Now, we can start our application and test it:

```bash
go run main.go
```

To test the application, we can use the ```grpcurl``` tool. For example, I'm going to list the different endpoints with the following command:

```bash
grpcurl -plaintext localhost:50051 list  
CommentService
grpc.reflection.v1alpha.ServerReflection
```

> NOTE: change port 50051 for your application port. 

Our application responds in different endpoints, the first one is for our CommentService and the second one is the Reflection API.

If we continue calling we can see the different methods that contain our endpoint:

```bash
grpcurl -plaintext localhost:50051 list CommentService                             
CommentService.Insert
CommentService.Retrieve
```

We've received two methods, also the same that we defined in de ```comment.proto``` file. So, it looks good. Now, I'll call the insert method:

```bash
grpcurl -plaintext localhost:50051 CommentService.Insert 
ERROR:
  Code: Unimplemented
  Message: method Insert not implemented
```

Yeah, our application is responding, so it runs ok but we've received an error code. This error is because we have not implemented the different methods as we have used the auto-generated code, and we only have defined an empty struct for our server.

# Implement the server code
Now, we can implement the business logic application, which will depend on each case. For this example, I'm only going to return a list of comments, but you can use whatever you want in your application.

As we defined the server struct, we only have to define the Retrieve method:

```go
func (s *server) Retrieve(context.Context, *domain.RetrieveRequest) (*domain.Comments, error) {
    return &domain.Comments{
        Commets: []*domain.Comment{
            &domain.Comment{
                IdComment:  1,
                IdObject:   12,
                TypeObject: "film",
                IdUser:     20,
                Comment:    "",
            },
        },
    }, nil
}
```

Now, we run the application again and test it by calling the Retrieve endpoint:

```bash
grpcurl -plaintext localhost:50051 CommentService.Retrieve
{
  "comments": [
    {
      "idComment": "1",
      "idObject": 12,
      "typeObject": "film",
      "idUser": 20
    },
    {
      "idComment": "2",
      "idObject": 12,
      "typeObject": "film",
      "idUser": 20
    }
  ]
}
```

And now, we can see how the application responds to the comment that we return.

# Create a client

# Using container as runtime