package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/J-Hendy/domain-insights/properties"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// func x(){

// }

func main() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres sslmode=disable dbname=domain_insights password=jupiter")
	// defer db.Close()
	if err != nil {
		logrus.Fatalf("could not establish connection to the database %v", err.Error())
	}
	
	// x()

	// Migrate the schema
	db.AutoMigrate(&properties.PropertyDetails{})
	
	handler := &PropertyHandler{
		DB: db,
	}
	
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/properties", handler.ServeProperties)
	r.Run() // listen and serve on 0.0.0.0:8080
}

// PropertyHandler server api endpoints related to properties
type PropertyHandler struct {
	DB *gorm.DB
}

// ServeProperties return the list of properties
func (handler *PropertyHandler)ServeProperties(c *gin.Context) {
	file, err := os.Open("./data/sales-results.json")
	if err != nil {
		logrus.Errorf("fail to load properties %v", err.Error())
	}

	b, _ := ioutil.ReadAll(file)
	var properties []*properties.PropertyDetails
	if err = json.Unmarshal(b, &properties); err != nil {
		logrus.Errorf("fail to unmarshal bytes to property list %v", err)
		c.JSON(500, gin.H{"err": "some error"})
		return
	}
	for _, item := range properties {
	
		handler.DB.Create(item)
	}

	c.JSON(200, &properties)
}
