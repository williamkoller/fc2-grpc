# fc2-grpc

- required `go`

- install protobuf-compile
  - `sudo apt install protobuf-compile`

- install protoprotobuf
  - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
  - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`

- add in your bash
  - `export PATH="$PATH:$(go env GOPATH)/bin"`
  - `source ~/.bashrc` or `source ~/.zshrc`