package main

import "time"

// SouborMetadata reprezentuje metadata o uploadovaném souboru
type SouborMetadata struct {
	Nazev       string `json:"nazev" example:"dokument.pdf" binding:"required"`
	Popis       string `json:"popis" example:"Důležitý dokument"`
	Kategorie   string `json:"kategorie" example:"dokumenty" binding:"required"`
	IsVerejny   bool   `json:"je_verejny" example:"false"`
	Velikost    int64  `json:"velikost,omitempty" example:"1024"`
	MimeType    string `json:"mime_type,omitempty" example:"application/pdf"`
}

// UploadOdpoved reprezentuje odpověď po úspěšném uploadu
type UploadOdpoved struct {
	ID           int             `json:"id" example:"1"`
	Metadata     SouborMetadata  `json:"metadata"`
	NazevSouboru string          `json:"nazev_souboru" example:"dokument_123456.pdf"`
	CestaSouboru string          `json:"cesta_souboru" example:"/uploads/dokument_123456.pdf"`
	DatumUpload  time.Time       `json:"datum_upload" example:"2023-01-01T00:00:00Z"`
	Zprava       string          `json:"zprava" example:"Soubor byl úspěšně nahrán"`
}

// ChybovaOdpoved reprezentuje chybovou odpověď
type ChybovaOdpoved struct {
	Kod    int    `json:"kod" example:"400"`
	Zprava string `json:"zprava" example:"Špatný požadavek"`
}

// SouborInfo reprezentuje informace o uploadovaném souboru
type SouborInfo struct {
	ID           int             `json:"id" example:"1"`
	Metadata     SouborMetadata  `json:"metadata"`
	NazevSouboru string          `json:"nazev_souboru" example:"dokument_123456.pdf"`
	CestaSouboru string          `json:"cesta_souboru" example:"/uploads/dokument_123456.pdf"`
	DatumUpload  time.Time       `json:"datum_upload" example:"2023-01-01T00:00:00Z"`
}