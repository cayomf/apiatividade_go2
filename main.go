package main

import (
	"apiatividade_go2/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "apiatividade_go2/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Correção de Atividades API
// @version 1.0
// @description This is an example API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8084
func main() {
	router := gin.Default()
	router.POST("/updateEntrega", updateEntrega)
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run("localhost:8084")
}

// updateEntrega responds with the list of all books as JSON.
// updateEntrega             godoc
// @Summary      Atualiza a nota na entrega do aluno
// @Description  Atualiza a nota na entrega do aluno
// @Tags         nota
// @Produce      json
// @Param        body     body     models.EntregaRequest     true        "EntregaRequest"
// @Success      200  {array}  models.EntregaRequest
// @Router       /updateEntrega [post]
func updateEntrega(c *gin.Context) {
	// Create a new HTTP PUT request
	req, err := http.NewRequest("PUT", "http://localhost:8083/updateEntrega", nil)
	if err != nil {
		log.Fatal(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Create a new Data object and encode it as JSON
	var request models.EntregaRequest
	c.BindJSON(&request)

	jsonData, err := json.Marshal(request)
	if err != nil {
		log.Fatal(err)
	}

	// Set the JSON data as the request body
	req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonData))

	log.Println(string(fmt.Sprintf("%v", req.Body)))

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print the response body
	log.Println(string(body))
}
