package server

import (
	"context"

	"exemplo.com/desafioGo/assets/rpc"
	"exemplo.com/desafioGo/internal/domain/entities"
	"exemplo.com/desafioGo/internal/repositories/database"
	"github.com/pkg/errors"
)

type Server struct {
	// empresas    []*rpc.Empresa
	EmpresaRepo *database.EmpresaRepo
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) AddEmpresa(ctx context.Context, emp *rpc.Empresa) (*rpc.Empresa, error) {

	e := entities.Empresa{
		// Id:        emp.GetId(),
		Nome:      emp.GetNome(),
		Sede:      emp.GetSede(),
		Atividade: emp.GetAtividade(),
		Emissoes:  emp.GetEmissoes(),
	}

	empresa, err := s.EmpresaRepo.AddEmpresa(ctx, &e)
	if err != nil {
		return nil, errors.Wrap(err, "error adding empresa")
	}
	empresaResp := EmpresaDBtoRpc(*empresa)
	return &empresaResp, nil

}

// Updates emissoes field in the database
func (s *Server) UpdateEmissoes(ctx context.Context, param *rpc.UpdateEmissoesParam) (*rpc.Empresa, error) {
	id := param.Id
	emissoes := param.Emissoes

	empresa, err := s.EmpresaRepo.UpdateEmissoes(ctx, id, emissoes)
	if err != nil {
		return nil, errors.Wrap(err, "error updating emissoes field on empresa")
	}
	empresaResp := EmpresaDBtoRpc(*empresa)
	return &empresaResp, nil

}

func (s *Server) GetEmpresas(ctx context.Context, _ *rpc.Empty) (*rpc.ListEmpresas, error) {
	empresas, err := s.EmpresaRepo.GetEmpresas(ctx)

	if err != nil {
		return nil, errors.Wrap(err, "error finding empresas")
	}

	emps := []*rpc.Empresa{}

	for _, empresa := range empresas {
		rpcEmpresa := EmpresaDBtoRpc(*empresa)
		emps = append(emps, &rpcEmpresa)
	}

	ListEmpresas := rpc.ListEmpresas{
		Empresas: emps,
	}

	return &ListEmpresas, nil

}

//Creates a new empresa on DB, then adds it to the EmpresaRepo and finally converts DB object to rpc object
// returns the new empresa

func (s *Server) GetEmpresa(ctx context.Context, empresaId *rpc.ParamId) (*rpc.Empresa, error) {
	id := empresaId.Id

	empresa, err := s.EmpresaRepo.GetEmpresa(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "error getting empresa by id")
	}
	empresaResp := EmpresaDBtoRpc(*empresa)
	return &empresaResp, nil
}

//Makes use of the method deleteEmpresaId from EmpresaRepo to delete
func (s *Server) DeleteEmpresaById(ctx context.Context, empresaId *rpc.ParamId) (*rpc.Empty, error) {
	id := empresaId.Id
	err := s.EmpresaRepo.DeleteEmpresaId(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "error deleting empresa by id")
	}
	return &rpc.Empty{}, nil
}

func (s *Server) DeleteEmpresaByNome(ctx context.Context, empresaNome *rpc.ParamNome) (*rpc.Empty, error) {
	nome := empresaNome.Nome
	err := s.EmpresaRepo.DeleteEmpresaNome(ctx, nome)
	if err != nil {
		return nil, errors.Wrap(err, "error deleting empresa by name")
	}
	return &rpc.Empty{}, nil
}

func EmpresaDBtoRpc(emp entities.Empresa) rpc.Empresa {

	return rpc.Empresa{
		// Id:        emp.Id,
		Nome:      emp.Nome,
		Sede:      emp.Sede,
		Atividade: emp.Atividade,
		Emissoes:  emp.Emissoes,
	}
}
