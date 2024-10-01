package models

import "gorm.io/gorm"

type Quizdua struct {
	gorm.Model
	Question    string `json:"question"`
	Questionurl string `json:"questionurl"`
	Answ        string `json:"answer"`
	Optiona     string `json:"optiona"`
	Optionb     string `json:"optionb"`
	Optionc     string `json:"optionc"`
	Optiond     string `json:"optiond"`
	Optione     string `json:"optione"`
	Result      int    `json:"result"`
	Tipe        string `json:"tipe"`
	Val         string `json:"value"`
	Vala        int    `json:"valuea"`
	Valb        int    `json:"valueb"`
	Valc        int    `json:"valuec"`
	Vald        int    `json:"valued"`
	Vale        int    `json:"valuee"`
}