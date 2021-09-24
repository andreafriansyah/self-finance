package saldoServices

import (
	"bobee/database"
	"bobee/models"
	"log"
)

func GetSaldo() ([]models.Saldo, error) {
	data := []models.Saldo{}
	res := database.Conn.Table("saldos").First(&data)
	log.Println("data itu :", data)
	err := res.Error
	if err != nil {
		log.Println("[Error] saldoServices.GetAll : ", err)
		return []models.Saldo{}, err
	}
	return data, nil
}

func AddIncome(request int) (models.Saldo, error) {
	data := models.Saldo{}
	var id = 1
	res := database.Conn.Table("saldos").Where("saldos.id = ?", id).First(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] saldoServices.AddIncome : ", err)
		return models.Saldo{}, err
	}
	var saldo_baru = data.JumlahSaldo + request
	UpdSaldo(saldo_baru)
	return data, nil
}

func AddOutcome(request int) (models.Saldo, error) {
	data := models.Saldo{}
	var id = 1
	res := database.Conn.Table("saldos").Where("saldos.id = ?", id).First(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] saldoServices.AddOutcome : ", err)
		return models.Saldo{}, err
	}
	var saldo_baru = data.JumlahSaldo - request
	UpdSaldo(saldo_baru)
	return data, nil
}

func UpdSaldo(saldo_baru int) (models.Saldo, error) {
	data := models.Saldo{}
	var id = 1
	res := database.Conn.Model(&data).Where("saldos.id = ?", id).Update("jumlah_saldo", saldo_baru)
	log.Println("saldo baru : ", data)
	err := res.Error
	if err != nil {
		log.Println("[Error] saldoServices.UpdSaldo : ", err)
		return models.Saldo{}, err
	}

	return data, nil
}
