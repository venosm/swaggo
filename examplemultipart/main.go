package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "examplemultipart/docs"
)

var nahraneSoubory = []SouborInfo{}
var nextID = 1

// @title           API pro upload souborů
// @version         1.0
// @description     REST API pro upload souborů s multipart/form-data a OpenAPI 3.0 dokumentací
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI Specifikace
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := gin.Default()

	// Vytvoření složky pro uploads
	os.MkdirAll("uploads", 0755)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/upload", NahrajSoubor)
		v1.GET("/soubory", ZiskejSoubory)
		v1.GET("/soubory/:id", ZiskejSoubor)
		v1.GET("/soubory/:id/download", StahniSoubor)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

// NahrajSoubor zpracuje upload souboru s metadaty
// @Summary      Nahraj soubor
// @Description  Nahraje soubor spolu s JSON metadaty pomocí multipart/form-data
// @Tags         soubory
// @Accept       multipart/form-data
// @Produce      json
// @Param        file      formData  file    true  "Soubor k nahrání"
// @Param        metadata  formData  string  true  "JSON metadata souboru"  SchemaExample({"nazev":"dokument.pdf","popis":"Důležitý dokument","kategorie":"dokumenty","je_verejny":false})
// @Success      201       {object} UploadOdpoved  "Soubor byl úspěšně nahrán"
// @Failure      400       {object} ChybovaOdpoved  "Špatný požadavek"
// @Failure      500       {object} ChybovaOdpoved  "Chyba serveru"
// @Router       /upload [post]
func NahrajSoubor(c *gin.Context) {
	// Získání souboru z multipart formu
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Soubor nebyl nalezen: " + err.Error(),
		})
		return
	}
	defer file.Close()

	// Získání JSON metadat z form data
	metadataJSON := c.PostForm("metadata")
	if metadataJSON == "" {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Metadata jsou povinná",
		})
		return
	}

	var metadata SouborMetadata
	if err := json.Unmarshal([]byte(metadataJSON), &metadata); err != nil {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Neplatný formát metadat: " + err.Error(),
		})
		return
	}

	// Validace povinných polí
	if metadata.Nazev == "" || metadata.Kategorie == "" {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Název a kategorie jsou povinné",
		})
		return
	}

	// Generování jedinečného názvu souboru
	ext := filepath.Ext(header.Filename)
	timestamp := time.Now().Unix()
	novoJmeno := fmt.Sprintf("%s_%d%s", 
		metadata.Nazev[:min(len(metadata.Nazev), 50)], 
		timestamp, 
		ext)
	
	// Cesta pro uložení
	cestaSouboru := filepath.Join("uploads", novoJmeno)

	// Vytvoření souboru na disku
	out, err := os.Create(cestaSouboru)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ChybovaOdpoved{
			Kod:    http.StatusInternalServerError,
			Zprava: "Nepodařilo se vytvořit soubor: " + err.Error(),
		})
		return
	}
	defer out.Close()

	// Kopírování dat
	size, err := io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ChybovaOdpoved{
			Kod:    http.StatusInternalServerError,
			Zprava: "Nepodařilo se uložit soubor: " + err.Error(),
		})
		return
	}

	// Nastavení dodatečných informací
	metadata.Velikost = size
	metadata.MimeType = header.Header.Get("Content-Type")

	// Vytvoření záznamu o souboru
	souborInfo := SouborInfo{
		ID:           nextID,
		Metadata:     metadata,
		NazevSouboru: novoJmeno,
		CestaSouboru: cestaSouboru,
		DatumUpload:  time.Now(),
	}
	
	nahraneSoubory = append(nahraneSoubory, souborInfo)
	nextID++

	// Odpověď
	odpoved := UploadOdpoved{
		ID:           souborInfo.ID,
		Metadata:     metadata,
		NazevSouboru: novoJmeno,
		CestaSouboru: cestaSouboru,
		DatumUpload:  souborInfo.DatumUpload,
		Zprava:       "Soubor byl úspěšně nahrán",
	}

	c.JSON(http.StatusCreated, odpoved)
}

// ZiskejSoubory vrací seznam všech nahraných souborů
// @Summary      Získej seznam souborů
// @Description  Vrátí seznam všech nahraných souborů
// @Tags         soubory
// @Accept       json
// @Produce      json
// @Success      200  {array}   SouborInfo  "Seznam souborů"
// @Router       /soubory [get]
func ZiskejSoubory(c *gin.Context) {
	c.JSON(http.StatusOK, nahraneSoubory)
}

// ZiskejSoubor vrací informace o konkrétním souboru
// @Summary      Získej informace o souboru
// @Description  Vrátí informace o konkrétním souboru podle ID
// @Tags         soubory
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "ID souboru"
// @Success      200  {object}  SouborInfo      "Informace o souboru"
// @Failure      404  {object}  ChybovaOdpoved  "Soubor nenalezen"
// @Router       /soubory/{id} [get]
func ZiskejSoubor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Neplatné ID souboru",
		})
		return
	}

	for _, soubor := range nahraneSoubory {
		if soubor.ID == id {
			c.JSON(http.StatusOK, soubor)
			return
		}
	}

	c.JSON(http.StatusNotFound, ChybovaOdpoved{
		Kod:    http.StatusNotFound,
		Zprava: "Soubor nenalezen",
	})
}

// StahniSoubor umožní stáhnout soubor
// @Summary      Stáhni soubor
// @Description  Stáhne soubor podle ID
// @Tags         soubory
// @Param        id   path  int  true  "ID souboru"
// @Success      200  {file}  file        "Soubor"
// @Failure      404  {object}  ChybovaOdpoved  "Soubor nenalezen"
// @Router       /soubory/{id}/download [get]
func StahniSoubor(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Neplatné ID souboru",
		})
		return
	}

	for _, soubor := range nahraneSoubory {
		if soubor.ID == id {
			c.File(soubor.CestaSouboru)
			return
		}
	}

	c.JSON(http.StatusNotFound, ChybovaOdpoved{
		Kod:    http.StatusNotFound,
		Zprava: "Soubor nenalezen",
	})
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}