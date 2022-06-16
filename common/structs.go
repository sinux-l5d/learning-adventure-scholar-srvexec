package common

import "encoding/json"

// Une fonction de type Executor prend en paramètre un code à évaluer et un exercice, et renvoie un status et des logs (optionels).
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

// Créer un nouvel exercice à partir d'un json
func NewExercice(data string) (*Exercice, error) {
	var exo Exercice
	if err := json.Unmarshal([]byte(data), &exo); err != nil {
		return nil, err
	}

	return &exo, nil
}

// Unmarshal le contexte d'un exercice sous format json et l'écrit dans la variable donnée
// La variable ctx doit être un pointeur vert une structure
func (e *Exercice) UnmarshalContexte(ctx interface{}) error {
	if e.Contexte == nil {
		return nil
	}

	if err := json.Unmarshal(e.Contexte, ctx); err != nil {
		return err
	}

	return nil
	// e.Contexte.UnmarshalJSON(&ctx)
}

// Environnement pour exécuter du code
type Environment struct {
	Name string
	Exec Executor
}

// Code à exécuter en fonction d'un exercice
type ToExecute struct {
	Code     string   `json:"code"`
	Exercice Exercice `json:"exercice"`
}

// Étant donné une chaîne JSON, renvoie un pointeur vers une structure ToExecute ou une erreur.
func NewToExecute(data string) (*ToExecute, error) {
	var toExec ToExecute
	if err := json.Unmarshal([]byte(data), &toExec); err != nil {
		return nil, err
	}

	return &toExec, nil
}
