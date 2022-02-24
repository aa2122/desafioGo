package main

import (
	"fmt"
	"log"
	"net/http"

	"exemplo.com/desafioGo/assets/rpc"
	"exemplo.com/desafioGo/config"
	server "exemplo.com/desafioGo/internal/empresaServer"
	handlers "exemplo.com/desafioGo/internal/handler"
	"exemplo.com/desafioGo/internal/middlewares"
	"exemplo.com/desafioGo/internal/repositories/database"
	"github.com/go-chi/chi"
)

func main() {

	db := config.SetupDbConnection()

	repo := database.NewEmpresaRepo(db)

	repo.InitEmpresaDB()

	empresaServer := &server.Server{
		EmpresaRepo: repo,
	}

	twirphandler := rpc.NewEmpresaServiceServer(empresaServer, nil)
	twirpServer := server.NewServer()

	// ctx := context.Background()
	// falta tratar do config
	// a := application.New(ctx, cfg, &twirphandler)

	// handler := routes.Routes()

	r := chi.NewRouter()

	// group multi handlers (twirp, middleware)
	r.Group(func(group chi.Router) {
		//middleware
		group.Use(middlewares.MyAuth)
		group.Mount(rpc.EmpresaServicePathPrefix, rpc.NewEmpresaServiceServer(empresaServer, nil))
	})

	eh := handlers.NewHandler(twirpServer)

	r.HandleFunc("/GetEmpresas", eh.GetEmpresas)

	// rpcServer := http.Server{
	// 	Addr:         ":4000",
	// 	Handler:      router,
	// 	IdleTimeout:  20 * time.Second,
	// 	ReadTimeout:  20 * time.Second,
	// 	WriteTimeout: 20 * time.Second,
	// }
	// log.Fatal(rpcServer.ListenAndServe())

	log.Fatal(http.ListenAndServe(":4000", twirphandler))
	fmt.Println("Server started on :4000")

}
