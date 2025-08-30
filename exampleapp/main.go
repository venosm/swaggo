package main

import (
	"net/http"
	"os"
	"strconv"
	"time"

	_ "exampleapp/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var uzivatele = []Uzivatel{
	{
		ID:        1,
		Jmeno:     "Jan Novák",
		Email:     "jan.novak@example.com",
		DatumReg:  time.Now().AddDate(0, -6, 0),
		IsActivni: true,
		Vek:       25,
	},
	{
		ID:        2,
		Jmeno:     "Marie Svobodová",
		Email:     "marie.svoboda@example.com",
		DatumReg:  time.Now().AddDate(0, -3, 0),
		IsActivni: true,
		Vek:       30,
	},
}

// @title           API pro správu uživatelů
// @version         1.0
// @description     Jednoduché REST API pro správu uživatelů s OpenAPI 3.0 dokumentací
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

	v1 := r.Group("/api/v1")
	{
		v1.GET("/uzivatele", ZiskejUzivatele)
		v1.GET("/uzivatele/:id", ZiskejUzivatele)
		v1.POST("/uzivatele", VytvorUzivatele)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}

// ZiskejUzivatele vrací seznam všech uživatelů nebo konkrétního uživatele podle ID
// @Summary      Získej uživatele
// @Description  Vrátí seznam všech uživatelů nebo konkrétního uživatele podle ID
// @Tags         uzivatele
// @Accept       json
// @Produce      json
// @Param        id   path      int  false  "ID uživatele"
// @Success      200  {array}   Uzivatel    "Seznam uživatelů"
// @Success      200  {object}  Uzivatel    "Konkrétní uživatel"
// @Failure      404  {object}  ChybovaOdpoved  "Uživatel nenalezen"
// @Router       /uzivatele [get]
// @Router       /uzivatele/{id} [get]
func ZiskejUzivatele(c *gin.Context) {
	idParam := c.Param("id")

	if idParam == "" {
		c.JSON(http.StatusOK, uzivatele)
		return
	}

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Neplatné ID uživatele",
		})
		return
	}

	for _, uzivatel := range uzivatele {
		if uzivatel.ID == id {
			c.JSON(http.StatusOK, uzivatel)
			return
		}
	}

	c.JSON(http.StatusNotFound, ChybovaOdpoved{
		Kod:    http.StatusNotFound,
		Zprava: "Uživatel nenalezen",
	})
}

// VytvorUzivatele vytvoří nového uživatele
// @Summary      Vytvoř nového uživatele
// @Description  Vytvoří nového uživatele v systému
// @Tags         uzivatele
// @Accept       json
// @Produce      json
// @Param        uzivatel body     UzivatelRequest  true  "Data nového uživatele"
// @Success      201      {object} UspesnaOdpoved{data=Uzivatel}  "Uživatel byl úspěšně vytvořen"
// @Failure      400      {object} ChybovaOdpoved  "Špatný požadavek"
// @Router       /uzivatele [post]
// @description @file ./docs/swagger.md
func VytvorUzivatele(c *gin.Context) {
	var req UzivatelRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ChybovaOdpoved{
			Kod:    http.StatusBadRequest,
			Zprava: "Špatný formát dat: " + err.Error(),
		})
		return
	}

	novyUzivatel := Uzivatel{
		ID:        len(uzivatele) + 1,
		Jmeno:     req.Jmeno,
		Email:     req.Email,
		Vek:       req.Vek,
		DatumReg:  time.Now(),
		IsActivni: true,
	}

	uzivatele = append(uzivatele, novyUzivatel)

	c.JSON(http.StatusCreated, UspesnaOdpoved{
		Zprava: "Uživatel byl úspěšně vytvořen",
		Data:   novyUzivatel,
	})
}
