package mysql

import "github.com/jinzhu/gorm"

type Note struct {
	gorm.Model
	Title string
	Text  string
}
