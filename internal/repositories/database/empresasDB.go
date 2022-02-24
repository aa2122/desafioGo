package database

import (
	"context"
	"errors"
	"log"

	"exemplo.com/desafioGo/internal/domain/entities"
	"gorm.io/gorm"
)

type Empresas interface {
	GetEmpresas(ctx context.Context) ([]*entities.Empresa, error)
	AddEmpresa(ctx context.Context, emp *entities.Empresa) (*entities.Empresa, error)
	InitEmpresaDB()
}

type EmpresaRepo struct {
	Empresas
	DB *gorm.DB
}

func NewEmpresaRepo(db *gorm.DB) *EmpresaRepo {
	return &EmpresaRepo{
		DB: db,
	}
}

//Needs update : Id must be auto-generated
func (r *EmpresaRepo) AddEmpresa(ctx context.Context, emp *entities.Empresa) (*entities.Empresa, error) {

	log.Printf("Adding empresa to the DB ...")
	empresa := entities.Empresa{
		// Id:        emp.Id,
		Nome:      emp.Nome,
		Sede:      emp.Sede,
		Atividade: emp.Atividade,
		Emissoes:  emp.Emissoes,
	}

	result := r.DB.Create(&empresa)

	if result == nil {
		return nil, errors.New("internal error creating movie")
	}

	if result.Error != nil {
		return nil, errors.New("error creating movie")
	}

	log.Printf("Empresa %s was sucessfully created !\n", emp.Nome)
	return &empresa, nil
}

//Retrieves all empresas from DB
func (r *EmpresaRepo) GetEmpresas(ctx context.Context) ([]*entities.Empresa, error) {
	log.Printf("Querying db for empresas...")
	var empresas []*entities.Empresa
	result := r.DB.Model(&empresas).Find(&empresas)
	if result.Error != nil {
		return nil, errors.New("error fetching all empresas")
	}
	return empresas, nil
}

func (r *EmpresaRepo) GetEmpresa(ctx context.Context, id uint64) (*entities.Empresa, error) {
	log.Printf("Querying db for empresa %d", id)
	var empresa *entities.Empresa
	result := r.DB.Model(empresa).Where("id = ?", id).Find(&empresa)
	if result.Error != nil {
		return nil, errors.New("error fetching empresa ")
	}
	return empresa, nil
}

//Deletes an empresa given its id
func (r *EmpresaRepo) DeleteEmpresaId(ctx context.Context, id uint64) error {
	log.Printf("Deleting empresa %v from DB", id)
	result := r.DB.Unscoped().Delete(&entities.Empresa{}, id)
	if result == nil || result.Error != nil {
		return errors.New("error deleting empresa")
	}
	return nil
}

//Deletes an empresa given its name
func (r *EmpresaRepo) DeleteEmpresaNome(ctx context.Context, nome string) error {
	log.Printf("Deleting empresa %s from DB", nome)
	var empresa *entities.Empresa
	result := r.DB.Where("nome = ?", nome).First(&empresa).Delete(&empresa)
	// result := r.DB.Unscoped().Model(empresa).Where("nome = ?", nome).Delete(&empresa)
	if result == nil || result.Error != nil {
		return errors.New("error deleting empresa")
	}
	return nil
}

// Updates emissoes field in DB
func (r *EmpresaRepo) UpdateEmissoes(ctx context.Context, id uint64, emissoes string) (*entities.Empresa, error) {
	log.Printf("Updating emissoes on empresa %d", id)
	var empresa *entities.Empresa
	result := r.DB.Where("id = ?", id).First(&empresa).Update("emissoes", emissoes)
	if result == nil {
		return nil, errors.New("internal error updating movie")
	}

	if result.Error != nil {
		return nil, errors.New("error updating empresa")
	}

	return empresa, nil
}

func (r *EmpresaRepo) InitEmpresaDB() {
	log.Printf("Migrating DB tables...")
	r.DB.AutoMigrate(&entities.Empresa{})
	log.Printf("DB Migrated !")
}
