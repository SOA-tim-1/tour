package model

import (
	"time"
)

// Coupon represents a discount coupon that can be applied to tours.
type Coupon struct {
	ID                         int64     `json:"id" gorm:"primaryKey:autoIncrement"`
	CouponHash                 string    `json:"couponHash" gorm:"type:varchar(255)"`
	DiscountPercentage         float64   `json:"discountPercentage"`
	DiscountExpiration         time.Time `json:"discountExpiration"`
	ApplicableTourId           *int      `json:"applicableTourId" gorm:"type:int"`
	CouponIssuerId             int       `json:"couponIssuerId"`
	IsApplicableToAllUserTours *bool     `json:"isApplicableToAllUserTours" gorm:"type:boolean;default:false"`
	IsValid                    bool      `json:"isValid" gorm:"type:boolean;default:false"`
}

// Assuming other related types like Checkpoint and Equipment are defined elsewhere in your package.
