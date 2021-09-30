package grafikServices

import (
	"bobee/database"
	"bobee/models"
	"log"
)

func GetDataByFilter(tahun string) ([]models.Income, []models.Outcome, error) {
	income := models.Income{}
	outcome := models.Outcome{}
	if tahun == "" {
		res := database.Conn.Table("finances").
			Select("sum(jumlah)").
			Where("type = 'Pemasukan'").
			Where("extract(year from tanggal) = 2021").
			Where("extract(month from tanggal) = 2021").
			Scan(&income)

		database.Conn.Table("finances").
			Select("sum(jumlah)").
			Where("type = 'Pengeluaran'").
			Where("extract(year from tanggal) = 2021").
			Where("extract(month from tanggal) = 2021").
			Scan(&outcome)
		log.Println("sum outcome: ", outcome)
		err := res.Error
		if err != nil {
			log.Println("[Error] grafikServices.GetDataByFilter : ", err)
			return []models.Income{}, []models.Outcome{}, err
		}
	} else if tahun != "" {
		res := database.Conn.
			Table("finances").
			Select("jumlah").
			Where("extract(year from tanggal) = ?", 2021).
			Where("extract(month from tanggal) = ?", 9)

		err := res.Error
		if err != nil {
			log.Println("[Error] grafikServices.GetDataByFilter : ", err)
			return []models.Income{}, []models.Outcome{}, err
		}
	}
	return []models.Income{}, []models.Outcome{}, nil
}
