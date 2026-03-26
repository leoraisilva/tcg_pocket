package model

import "fmt"

type Tipo string

const (
	TipoFogo     Tipo = "Fogo"
	TipoAgua     Tipo = "Água"
	TipoPlanta   Tipo = "Planta"
	TipoEletrico Tipo = "Elétrico"
	TipoPsiquico Tipo = "Psíquico"
	TipoLutador  Tipo = "Lutador"
	TipoNoturno  Tipo = "Noturno"
	TipoMetal    Tipo = "Metal"
	TipoComum    Tipo = "Comum"
	TipoDragon   Tipo = "Dragão"
)

func (t Tipo) IsValid() bool {
	switch t {
	case TipoFogo, TipoAgua, TipoPlanta, TipoEletrico, TipoPsiquico, TipoLutador, TipoNoturno, TipoMetal, TipoComum, TipoDragon:
		return true
	}
	return false
}

func (t Tipo) String() string {
	return string(t)
}

func ParseTipo(s string) (Tipo, error) {
	tipo := Tipo(s)
	if !tipo.IsValid() {
		return "", fmt.Errorf("tipo inválido: %s", s)
	}
	return tipo, nil
}

func (t Tipo) GetTipo() string {
	switch t {
	case TipoFogo:
		return "Fogo"
	case TipoAgua:
		return "Água"
	case TipoPlanta:
		return "Planta"
	case TipoEletrico:
		return "Elétrico"
	case TipoPsiquico:
		return "Psíquico"
	case TipoLutador:
		return "Lutador"
	case TipoNoturno:
		return "Noturno"
	case TipoMetal:
		return "Metal"
	case TipoComum:
		return "Comum"
	case TipoDragon:
		return "Dragão"
	default:
		return "Desconecido"
	}
}
