package database

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/RTradeLtd/Temporal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

/*
	roll.Token = "POST_SERVER_ITEM_ACCESS_TOKEN"
	//roll.Environment = "production" // defaults to "development"

	r := gin.Default()
	r.Use(rollbar.Recovery(true))

	r.Run(":8080")
	func l(err error) {
	token := os.Getenv("ROLLBAR_TOKEN")
	rollbar.SetToken(token)
	rollbar.SetServerRoot("github.com/RTradeLtd/Temporal") // path of project (required for GitHub integration and non-project stacktrace collapsing)

	rollbar.Error(err)

	rollbar.Wait()
}
*/

var uploadObj *models.Upload
var userObj *models.User

type DatabaseManager struct {
	DB     *gorm.DB
	Upload *models.UploadManager
}

func Initialize() *DatabaseManager {
	dbm := DatabaseManager{}
	dbm.OpenDBConnection()
	dbm.RunMigrations()
	return &dbm
}

func (dbm *DatabaseManager) RunMigrations() {
	dbm.DB.AutoMigrate(uploadObj)
	dbm.DB.AutoMigrate(userObj)
}

// OpenDBConnection is used to create a database connection
func (dbm *DatabaseManager) OpenDBConnection() {
	db, err := gorm.Open("sqlite3", "./ipfs_database.db")
	if err != nil {
		log.Fatal(err)
	}
	dbm.DB = db
}

// CloseDBConnection is used to close a db
func CloseDBConnection(db *gorm.DB) {
	db.Close()
}

func (dbm *DatabaseManager) GetUpload(hash string, uploaderAddress string) []*models.Upload {
	var uploads []*models.Upload
	dbm.DB.Find(&uploads).Where("hash = ? AND uploader_address = ?", hash, uploaderAddress)
	return uploads
}

// GetUploads is used to retrieve all uploads
func GetUploads() []*models.Upload {
	var uploads []*models.Upload
	db = OpenDBConnection()
	db.Find(&uploads)
	return uploads
}

func GetUploadsForAddress(address string) []*models.Upload {
	var uploads []*models.Upload
	db = OpenDBConnection()
	db.Where("upload_address = ?", address).Find(&uploads)
	return uploads
}

// AddHash his used to add a hash to our database
func AddHash(c *gin.Context) error {
	var upload models.Upload
	hash := c.Param("hash")
	address, exists := c.GetPostForm("uploadAddress")
	if !exists {
		c.AbortWithError(http.StatusBadRequest, errors.New("uploadAddress param des not exist"))
		return errors.New("uploadAddress param des not exist")
	}
	holdTime, exists := c.GetPostForm("holdTime")
	if !exists {
		c.AbortWithError(http.StatusBadRequest, errors.New("holdTime param does not exist"))
		return errors.New("holdTime param does not exist")
	}
	holdTimeInt, err := strconv.ParseInt(holdTime, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return err
	}
	upload.Hash = fmt.Sprintf("%s", hash)
	upload.Type = "pin"
	upload.HoldTimeInMonths = holdTimeInt
	upload.UploadAddress = address
	db := OpenDBConnection()
	db.Create(&upload)
	db.Close()
	return nil
}

// AddFileHash is used to add the hash of a file to our database
func AddFileHash(c *gin.Context, hash string) {
	var upload models.Upload
	address := c.PostForm("uploadAddress")
	holdTimeInt, err := strconv.ParseInt(c.PostForm("holdTime"), 10, 64)
	if err != nil {
		c.Error(err)
	}
	upload.HoldTimeInMonths = holdTimeInt
	upload.UploadAddress = address
	upload.Hash = hash
	upload.Type = "file"
	db := OpenDBConnection()
	db.AutoMigrate(&upload)
	db.Create(&upload)
	db.Close()
}
