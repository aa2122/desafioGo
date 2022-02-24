package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"exemplo.com/desafioGo/assets/rpc"
	empresas "exemplo.com/desafioGo/internal/domain/entities"
	server "exemplo.com/desafioGo/internal/empresaServer"
	"exemplo.com/desafioGo/internal/response"
	"github.com/go-chi/chi"
)

var e1 = empresas.Empresa{
	Id:        1,
	Nome:      "Worten",
	Sede:      "Carnaxide",
	Atividade: "Retalho de Eletrónica",
	Emissoes:  "10mil toneladas",
}

var e2 = empresas.Empresa{
	Id:        2,
	Nome:      "Continente",
	Sede:      "Matosinhos",
	Atividade: "Retalho",
	Emissoes:  "20mil toneladas",
}

var id int

type AuthHandler struct {
	server *server.Server
}

func NewHandler(s *server.Server) *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) GetEmpresas(w http.ResponseWriter, r *http.Request) {
	h.server.GetEmpresas(context.Background(), &rpc.Empty{})

}

func (handler *AuthHandler) GetEmpresaByName(w http.ResponseWriter, r *http.Request) {

	nome := chi.URLParam(r, "nome")

	switch nome {
	case "worten":
		empresa := e1.GetInfoPublica()
		response.JSONResponse(w, http.StatusOK, empresa)
	case "continente":
		empresa := e2.GetInfoPublica()
		response.JSONResponse(w, http.StatusOK, empresa)
	default:
		response.ErrorResponse(w, http.StatusInternalServerError, "Erro")

	}

}

func (handler *AuthHandler) GetEmpresaByNameP(w http.ResponseWriter, r *http.Request) {

	nome := chi.URLParam(r, "nome")

	switch nome {
	case "wortenP":
		empresa := e1.GetInfoPrivada()
		response.JSONResponse(w, http.StatusOK, empresa)
	case "continenteP":
		empresa := e2.GetInfoPrivada()
		response.JSONResponse(w, http.StatusOK, empresa)
	default:
		response.ErrorResponse(w, http.StatusInternalServerError, "Erro")

	}

}

func (handler *AuthHandler) GetEmpresaBySede(w http.ResponseWriter, r *http.Request) {
	sede := chi.URLParam(r, "sede")

	switch sede {
	case "Matosinhos":
		empresa := e1.GetInfoPublica()
		response.JSONResponse(w, http.StatusOK, empresa)
	case "Carnaxide":
		empresa := e2.GetInfoPublica()
		response.JSONResponse(w, http.StatusOK, empresa)
	default:
		response.ErrorResponse(w, http.StatusInternalServerError, "Erro")

	}

}

func (handler *AuthHandler) AddEmpresa(w http.ResponseWriter, r *http.Request) {

	var e empresas.Empresa

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)

	if err != nil {
		log.Print(err)
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}

	e.Id = 666
	response.JSONResponse(w, http.StatusOK, e)

	// Deveria ter um servico para criar a Empresa
	// newEmpresa, err := handler.AddEmpresa(empresa)

}

func (handler *AuthHandler) UpdateEmpresa(w http.ResponseWriter, r *http.Request) {

	var e empresas.Empresa

	//  nome := chi.URLParam(r, "nome")

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&e)

	if err != nil {
		log.Print(err)
		response.ErrorResponse(w, http.StatusInternalServerError, err.Error())
	}
	response.JSONResponse(w, http.StatusOK, e)
}
