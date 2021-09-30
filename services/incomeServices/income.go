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

func GetDataByFilter(dari_tanggal, sampe_tanggal, keyword string, orderBy, orderDir string, skip int, limit int) ([]models.Finances, int, error) {
	data := []models.Finances{}
	totalData := []models.Finances{}
	var totalDatas int

	if dari_tanggal == "" && sampe_tanggal == "" {
		res := database.Conn.
			Table("finances").
			Where("type = 'Pemasukan'").
			Where("is_deleted = false")
		res.Scan(&totalData)
		res.Order(orderBy + " " + orderDir).Limit(limit).Offset(skip).Scan(&data)
		err := res.Error
		if err != nil {
			log.Println("[Error] financesServices.GetDataByFilter : ", err)
			return []models.Finances{}, totalDatas, err
		}
	} else if dari_tanggal != "" && sampe_tanggal != "" {
		res := database.Conn.
			Table("finances").
			Where("type = 'Pemasukan'").
			Where("is_deleted = false").
			Where("tanggal BETWEEN ? AND ?", dari_tanggal, sampe_tanggal)
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

// Create add database
func Create(request models.Finances) (models.Finances, error) {
	res := database.Conn.Table("finances").Create(&request).Scan(&request)
	err := res.Error

	if err != nil {
		log.Println("[Error] infoService.Create : ", err)
		return models.Finances{}, err
	}

	return request, nil
}
