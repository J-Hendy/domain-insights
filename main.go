package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"os"

	"github.com/J-Hendy/domain-insights/properties"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// func x(){

// }

func main() {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres sslmode=disable dbname=domain_insights password=jupiter")
	// defer db.Close()
	if err != nil {
		log.Fatalf("could not establish connection to the database %v", err.Error())
	}

	// x()

	// Migrate the schema
	db.AutoMigrate(&properties.PropertyDetails{})

	handler := &PropertyHandler{
		DB: db,
	}

	count, err := handler.PropertiesCount()

	if err != nil {
		log.Fatalf("couldn't get properties count from db %v", err.Error())
	}
	if count == 0 {
		log.Info("found empty table for properties, will import data from json")
		if err = handler.loadPropertiesFromJSON(); err != nil {
			log.Fatal(err.Error())
		}
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

// PropertiesCount shows the total count of properties in db
func (handler *PropertyHandler) PropertiesCount() (int, error) {
	var count int
	err := handler.DB.Table("property_details").Count(&count).Error
	return count, err
}

func (handler *PropertyHandler) loadPropertiesFromJSON() error {
	file, err := os.Open("./data/sales-results.json")
	if err != nil {
		return fmt.Errorf("fail to load properties %v", err.Error())
	}

	b, _ := ioutil.ReadAll(file)
	var properties []*properties.PropertyDetails
	if err = json.Unmarshal(b, &properties); err != nil {
		return fmt.Errorf("fail to unmarshal bytes to property list %v", err)
	}
	for _, item := range properties {
		if err = handler.DB.Create(item).Error; err != nil {
			return fmt.Errorf("fail to save property %s to db with error %s", item.ID, err.Error())
		}
	}

	return nil
}

// ServeProperties return the list of properties
func (handler *PropertyHandler) ServeProperties(c *gin.Context) {
	var properties []*properties.PropertyDetails
	postcode := c.DefaultQuery("postcode", "3000")
	if err := validateForm(c); err != nil {
		log.Errorf("got error when validate the request parameters %v", err)
		c.JSON(400, gin.H{"msg": "please provide valid postcode"})
		return
	}
	err := handler.DB.Where("postcode = ?", postcode).Find(&properties).Error
	if err != nil {
		log.Errorf("couldn't get properties from db %s ", err.Error())
		c.JSON(500, gin.H{"err": "some internal error, please try again later"})
		return
	}
	c.JSON(200, &properties)
}

func validateForm(c *gin.Context) error {
	validateError := fmt.Errorf("invalid postcode")
	//TODO: validate postcode
	// length must be 4
	// the string should be able to parse to integer
	postcode := c.DefaultQuery("postcode", "3000")
	if len(postcode) != 4 {
		return validateError
	}
	_, err := strconv.Atoi(postcode)
	if err != nil {
		return validateError
	}
	return nil
}
