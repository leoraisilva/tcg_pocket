package model

type Card struct {
	Id         int    `json:"id"`
	Nome       string `json:"nome"`
	Tipo       Tipo   `json:"tipo"`
	Estagio    int    `json:"estagio"`
	Habilidade string `json:"habilidade"`
	Ataque     string `json:"ataque"`
	PS         int    `json:"ps"`
	Recuo      int    `json:"recuo"`
	Fraqueza   Tipo   `json:"fraqueza"`
}
