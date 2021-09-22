package incomeService

import (
	"bobee/database"
	"bobee/models"
	"log"
	"strings"
)

func GetAll(search string, orderBy string, orderDir string, skip int, limit int) ([]models.Finances, int, error) {
	data := []models.Finances{}
	res := database.Conn.Table("finances")
	var totalData int
	if search == "" {
		res = res.Where("LOWER(finances.asaltujuan) like ?", "%"+strings.ToLower(search)+"%")
	}
	res.Count(&totalData)
	res.Order(orderBy + " " + orderDir).Limit(limit).Offset(skip).Find(&data)
	err := res.Error
	if err != nil {
		log.Println("[Error] incomeService.GetAll : ", err)
		return []models.Finances{}, totalData, err
	}
	return data, totalData, nil
}

func GetDataByFilter(asaltujuan string, keyword string, orderBy, orderDir string, skip int, limit int) ([]models.Finances, int, error) {
	data := []models.Finances{}
	totalData := []models.Finances{}
	var totalDatas int

	if asaltujuan == "" {
		res := database.Conn.Raw("select id, type, asaltujuan, jumlah, keterangan, tanggal from finances where type = 'Pengeluaran'")
		res.Scan(&totalData)
		res.Order(orderBy + " " + orderDir).Limit(limit).Offset(skip).Scan(&data)
		err := res.Error
		if err != nil {
			log.Println("[Error] financesServices.GetDataByFilter : ", err)
			return []models.Finances{}, totalDatas, err
		}
	} else if asaltujuan != "" {
		res := database.Conn.Table("select id, type, asaltujuan, jumlah, keterangan, tanggal from finances where type = 'Pengeluaran'")
		res.Scan(&totalData)
		res.Order(orderBy + " " + orderDir).Limit(limit).Offset(skip).Scan(&data)
		err := res.Error
		if err != nil {
			log.Println("[Error] financesServices.GetDataByFilter : ", err)
			return []models.Finances{}, totalDatas, err
		}
	}

	totalDatas = len(totalData)
	return data, totalDatas, nil
}
