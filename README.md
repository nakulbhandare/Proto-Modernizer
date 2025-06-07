# 🛠️ Proto Modernizer

This project demonstrates how to migrate `.proto` files from legacy **GoGoProto** to modern **`protoc-gen-go`** format, with a working **gRPC server and client** implementation in Go.

---

## 📦 Project Structure

```
proto-modernizer/
├── proto/
│   ├── legacy/           # Legacy .proto files using gogo/protobuf extensions
│   └── modern/           # Updated .proto files using protoc-gen-go
│       ├── user/
│       ├── auth/
│       └── ...           # Separate folders per proto package
├── server/
│   ├── server.go         # gRPC server implementation
├── client/
│   └── client.go         # gRPC client implementation
```

---

## ✅ Migration Highlights

- Removed all usage of `gogo/protobuf` extensions.
- Cleaned and reorganized `.proto` files using standard `protoc-gen-go` style.
- Introduced separate Go packages per service for maintainability.
- Generated `.pb.go` and `_grpc.pb.go` using:
  ```bash
  protoc     --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    -I=. *.proto
  ```

---

## 🚀 Running the gRPC Example

### 1. Generate Go Code

```bash
cd proto/modern
protoc \
  --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  -I=. user.proto
```

### 2. Start the Server

```bash
cd server
go run server.go
```

### 3. Run the Client (in a separate terminal)

```bash
cd client
go run client.go
```

You should see:
```
User: ID=123, Name=John Doe, Description=Test user from server
```

---

## 💡 Use Case

This project can be used to:

- Showcase protobuf modernization and refactoring skills.
- Demonstrate knowledge of gRPC microservice architecture.
- Serve as a template for migrating deprecated proto stacks (e.g., Dropbox's GoGoProto removal).

---

## 🔗 Tech Stack

- Golang
- Protocol Buffers (`protoc-gen-go`, `protoc-gen-go-grpc`)
- gRPC

---

## ✍️ Author

Nakul Bhandare  
> Contact for freelance gRPC, DevOps, and cloud-native development assistance.
