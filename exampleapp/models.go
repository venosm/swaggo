package main

import (
	"time"
)

// Uzivatel reprezentuje uživatele v systému
type Uzivatel struct {
	ID         int        `json:"id" example:"1" description:"Jedinečný identifikátor uživatele"`
	Jmeno      string     `json:"jmeno" example:"Jan Novák" description:"Celé jméno uživatele"`
	Email      string     `json:"email" example:"jan.novak@example.com" description:"Emailová adresa uživatele"`
	DatumReg   time.Time  `json:"datum_registrace" example:"2023-01-01T00:00:00Z" description:"Datum a čas registrace"`
	IsActivni  bool       `json:"je_aktivni" example:"true" description:"Indikátor aktivity uživatele"`
	Vek        int        `json:"vek" example:"25" minimum:"0" maximum:"150" description:"Věk uživatele v letech"`
	StavZpravy StavZpravy `json:"stav_zpravy,omitempty" validate:"omitempty,oneof=AKTIVNI ARCHIVOVANA SMAZANA" swaggertype:"string" example:"AKTIVNI" description:"Stav zprávy (defaultní hodnota je aktivní)."`
}

// UzivatelRequest struktura pro vytvoření nového uživatele
type UzivatelRequest struct {
	Jmeno string `json:"jmeno" binding:"required" example:"Jan Novák" description:"Jméno nového uživatele"`
	Email string `json:"email" binding:"required" example:"jan.novak@example.com" description:"Email nového uživatele"`
	Vek   int    `json:"vek" binding:"required" example:"25" minimum:"0" maximum:"150" description:"Věk nového uživatele"`
}

// ChybovaOdpoved struktura pro chybové zprávy
type ChybovaOdpoved struct {
	Kod    int    `json:"kod" example:"400" description:"HTTP status kód chyby"`
	Zprava string `json:"zprava" example:"Špatný požadavek" description:"Popis chyby pro uživatele"`
}

// UspesnaOdpoved struktura pro úspěšné odpovědi
type UspesnaOdpoved struct {
	Zprava string      `json:"zprava" example:"Operace proběhla úspěšně" description:"Zpráva o úspěchu operace"`
	Data   interface{} `json:"data,omitempty" description:"Volitelná data odpovědi"`
}

type StavZpravy string //	@name	StavZpravy

const (
	AKTIVNI     StavZpravy = "AKTIVNI"
	ARCHIVOVANA StavZpravy = "ARCHIVOVANA"
	SMAZANA     StavZpravy = "SMAZANA"
)
