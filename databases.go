package main

import (
	"fmt"
	"os"

	_ "github.com/apache/calcite-avatica-go"
	"github.com/jinzhu/gorm"
)

// Db variable
var Db *gorm.DB
var err error

func initDb() (*gorm.DB, error) {
	Db, err = gorm.Open("avatica", os.Getenv("pqs_url"))

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// Db.AutoMigrate(&RequestData{}, &ResponseData{}, &CpaBriFitur{}, &CpaInvestasi{}, &CpaLoan{},
	// 	&CpaRcmdProduct{}, &CpaSaving{}, &CpaSegmenAccno{}, &CpaTxnetworkAgg{}, &CpaTransaksi{}, &CpaCust{})

	if err != nil {
		fmt.Println("Connection Database Error ", err.Error())
	} else {
		Db.DB().SetMaxIdleConns(2)
		Db.DB().SetMaxOpenConns(1000)
		Db.LogMode(true)
		fmt.Println("Connected")
	}

	return Db, err
}

func connect() (*gorm.DB, error) {
	return Db, err
}
