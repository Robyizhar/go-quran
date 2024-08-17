package models

import (
	"time"
)

type Surah struct {
	Id             int16     `gorm:"primarykey" json:"id"`                     //id
	IndonesianName string    `gorm:"type:varchar(300)" json:"indonesian_name"` //latin
	ArabicName     string    `gorm:"type:varchar(300)" json:"arabic_name"`     //arabic
	Meaning        string    `gorm:"type:varchar(300)" json:"meaning"`         // translation
	Location       string    `gorm:"type:varchar(300)" json:"location"`        // location
	TotalAyah      int       `json:"total_ayah"`                               // num_ayah
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
	Ayahs          []Ayah
}

type Ayah struct {
	Id             int16     `gorm:"primarykey" json:"id"`
	SurahId        uint      `gorm:"index"`
	NumberAyah     int       `json:"number_ayah"`
	Juz            int       `json:"juz"`
	Arabic         string    `gorm:"type:text" json:"arabic"`
	Kitabah        string    `gorm:"type:text" json:"kitabah"`
	Latin          string    `gorm:"type:text" json:"latin"`
	Translation    string    `gorm:"type:text" json:"translation"`
	Interpretation string    `gorm:"type:text" json:"interpretation"`
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
