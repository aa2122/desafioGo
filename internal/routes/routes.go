package routes

import (
	"exemplo.com/desafioGo/assets/rpc"
	server "exemplo.com/desafioGo/internal/empresaServer"
	handlers "exemplo.com/desafioGo/internal/handler"
	middlewares "exemplo.com/desafioGo/internal/middlewares"

	"github.com/go-chi/chi"
)

var authHandler handlers.AuthHandler

func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// group multi handlers (twirp, middleware)
	r.Group(func(group chi.Router) {
		//middleware
		r.Use(middlewares.MyAuth)
		r.Mount(rpc.EmpresaServicePathPrefix, rpc.NewEmpresaServiceServer(server.NewServer(), nil))
	})

	eh := handlers.NewHandler(server.NewServer())

	r.HandleFunc("/GetEmpresas", eh.GetEmpresas)

	//Public API
	r.Group(func(r chi.Router) {
		r.Get("/empresa/{nome}", authHandler.GetEmpresaByName)
		r.Get("/sede/{sede}", authHandler.GetEmpresaBySede)
		r.Post("/empresa/sportzone", authHandler.AddEmpresa)
		r.Post("/empresa/wells", authHandler.AddEmpresa)
		// r.Delete("/empresa/wells", authHandler.DeleteEmpresa)
	})

	// Protected API
	r.Group(func(r chi.Router) {
		// r.Use(middlewares.MyAuth)
		r.Get("/empresaP/{nome}", authHandler.GetEmpresaByNameP)
		r.Put("/empresaP/{nome}", authHandler.UpdateEmpresa)
	})

	return r
}
