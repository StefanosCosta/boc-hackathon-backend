package data

import (
	"OpenBankingAPI/database"
	"OpenBankingAPI/models"
)

func ByAuditorAndBusinessId(access *[]models.AuditorAccess, auditorId uint, businessId uint) error {
	return database.DB.Where("auditor_id = ? and business_id = ?", auditorId, businessId).Find(access).Error
}

func ByAuditorAndBusinessIdRequest(access *[]models.AuditorRequests, auditorId uint, businessId uint) error {
	return database.DB.Where("auditor_id = ? and business_id = ?", auditorId, businessId).Find(access).Error
}

func ByBusinessIdRequest(access *[]models.AuditorRequests, businessId uint) error {
	return database.DB.Where("business_id = ?", businessId).Find(access).Error
}

func ByRequestId(access *[]models.AuditorRequests, requestId uint) error {
	return database.DB.Where("id = ?", requestId).Find(access).Error
}

func ByAuditorAndBusinessID(users *models.AuditorRequests, auditorId uint, businessId string) error {
	return database.DB.Where("auditor_id = ? AND business_id = ?", auditorId, businessId).Find(users).Error
}
