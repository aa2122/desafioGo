package entities

import "fmt"

type Empresa struct {
	Id uint64 `gorm:"primarykey, autoIncrement" json:"id"`
	// Id        string `json:"id"`
	Nome      string `gorm:"index:idx_empresa_nome,unique" json:"nome"`
	Sede      string `gorm:"index:idx_empresa_sede" json:"sede"`
	Atividade string `gorm:"index:idx_empresa_atividade" json:"atividade"`
	Emissoes  string `gorm:"index:idx_empresa_emissoes" json:"emissoes"`
}

// Retrieves public info from the empresa
func (e *Empresa) GetInfoPublica() []string {

	id := fmt.Sprintf("%v", e.Id)

	infoP := []string{id, e.Nome, e.Atividade, e.Sede}
	return infoP
}

// Retrieves public + private info(emissoes) from the empresa
func (e *Empresa) GetInfoPrivada() []string {
	return append(e.GetInfoPublica(), e.Emissoes)
}
