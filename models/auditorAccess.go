package models

import (
	"gorm.io/gorm"
	"time"
)

type AuditorAccess struct {
	gorm.Model
	AuditorId  uint       `gorm:"many2one:user_id;" json:"auditor_id"`
	BusinessId uint       `json:"business_id"`
	Key        string     `json:"key"`
	ExpiresAt  *time.Time `json:"expires_at"`
}

type AuditorRequests struct {
	gorm.Model
	AuditorId      uint       `json:"auditor_id"`
	BusinessId     uint       `json:"business_id"`
	Accept         bool       `json:"accepted"`
	AcceptedAt     *time.Time `json:"accepted_at"`
	SubscriptionID string     `json:"subscription_id"`
}
