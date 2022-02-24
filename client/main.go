package main

import (
	"context"
	"fmt"
	"net/http"

	"exemplo.com/desafioGo/assets/rpc"
	"github.com/golang/protobuf/jsonpb"
)

func main() {

	//1st Part of the Project
	// r := routes.NewRouter()

	// http.ListenAndServe(":6666", r)

	// Twirp

	client := rpc.NewEmpresaServiceProtobufClient("localhost:4000", &http.Client{})

	ctx := context.Background()

	client.AddEmpresa(context.Background(), &rpc.Empresa{})

	client.DeleteEmpresaById(ctx, &rpc.ParamId{})

	data, _ := client.GetEmpresas(ctx, &rpc.Empty{})

	client.UpdateEmissoes(ctx, &rpc.UpdateEmissoesParam{})
	
	m := jsonpb.Marshaler{}
	result, _ := m.MarshalToString(data)
	fmt.Println("Client is running ...")
	fmt.Println(result)

}
