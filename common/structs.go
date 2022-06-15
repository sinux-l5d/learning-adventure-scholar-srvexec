package common

import "encoding/json"

// Une fonction de type Executor prend en paramètre un code à évaluer et in exercice, et renvoie un status et des logs (optionels).
type Executor func(j ToExecute) (s Status, log string)

type Exercice struct {
	Id          string          `json:"id"`
	Nom         string          `json:"nom"`
	Template    string          `json:"template"`
	Enonce      string          `json:"enonce"`
	Difficulte  uint8           `json:"difficulte"`
	Themes      []string        `json:"themes"`
	Langage     string          `json:"langage"`
	TempsMoyen  uint            `json:"tempsMoyen"`
	TempsMax    uint            `json:"tempsMaximum"`
	Contexte    json.RawMessage `json:"contexte"`
	Correction  string          `json:"correction"`
	Commentaire string          `json:"commentaire"`
	Aides       []string        `json:"aides"`
	Auteurs     []string        `json:"auteurs"`
}

type Environment struct {
	Name string
	Exec Executor
}

type ToExecute struct {
	Code     string   `json:"code"`
	Exercice Exercice `json:"exercice"`
}
