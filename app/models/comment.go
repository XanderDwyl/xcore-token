package models

import (
	"encoding/json"
	"log"
	"time"

	"github.com/a-fis/xcore-token/app/helpers"
)

// Comment ...
type Comment struct {
	ID    int64  `json:"-"  gorm:"AUTO_INCREMENT"`
	Title string `json:"title" `
	Body  string `json:"body" `
	Utime int64  `json:"utime"`
	Hash  string `json:"-" `
}

// CreateComment ...
func CreateComment(t, b string) (Comment, error) {
	u := Comment{}
	u.Title = t
	u.Body = b
	err := u.Create()
	return u, err
}

// Create ...
func (u *Comment) Create() error {
	ethereumData := helpers.EthereumData{}

	u.Utime = time.Now().Unix()
	u.Hash = u.HashString()

	ethereumData.HashedData = u.Hash
	ethereumData.Timestamp = u.Utime
	log.Println("DATA!!!")
	log.Println(ethereumData.ToString())
	helpers.PublishToEthNetwork(ethereumData.ToString())
	return db.Debug().Model(&u).Create(&u).Error
}

// GetCommentByID ...
func GetCommentByID(id int64) (*Comment, error) {
	var item Comment
	err := db.Debug().Model(&item).Where("id=?", id).Scan(&item).Error
	log.Print(err)
	return &item, err
}

// JSONString ...
func (u *Comment) JSONString() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// HashString ...
func (u *Comment) HashString() string {
	j, e := u.JSONString()
	if e != nil {
		log.Printf("error:%s", e.Error())
		return ""
	}
	return helpers.Sha512(j)
}

// Update ...
func (u *Comment) Update() error {
	return db.Debug().Model(&u).Update(&u).Error
}

// Destroy ...
func (u *Comment) Destroy() error {
	err := db.Debug().Model(&u).Delete(&u, "id=?", u.ID).Error
	return err
}

// GetComments ..
func GetComments() ([]Comment, error) {

	log.Printf("GetComments")
	var list []Comment
	var err error
	err = db.Debug().Model(&Comment{}).Order("id desc").Scan(&list).Limit(100).Error
	return list, err
}
