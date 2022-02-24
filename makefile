chi: 
go get -u github.com/go-chi/chi/v5

twirpBuf: 
	go get github.com/twitchtv/twirp/protoc-gen-twirp
	go get github.com/golang/protobuf/protoc-gen-go

protoc : 
	protoc --proto_path=/home/amiguela/go/src/exemplo.com/desafioGo/:. --twirp_out=../ --go_out=../ empresa.proto







