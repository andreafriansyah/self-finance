package saldoServices

import (
	"bobee/database"
	"bobee/models"
	"log"
)

func GetSaldo() ([]models.Saldo, error) {
	data := []models.Saldo{}
	res := database.Conn.Table("saldo").First(&data)
	log.Println("data itu :", data)
	err := res.Error
	if err != nil {
		log.Println("[Error] saldoServices.GetAll : ", err)
		return []models.Saldo{}, err
	}
	return data, nil
}
