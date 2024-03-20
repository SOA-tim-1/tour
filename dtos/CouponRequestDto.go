package dtos

import (
	"database-example/model"
	"time"
)

// CouponRequest represents the JSON structure for a coupon creation or update request.
type CouponRequest struct {
	ExpirationDate     time.Time `json:"expirationDate"`
	DiscountPercentage float64   `json:"discountPercentage"`
	TourID             *int      `json:"tourId,omitempty"` // Use a pointer to allow for nullability
	UserID             int       `json:"userId"`
}

// ToCouponModel converts a CouponRequest into a Coupon model.
// This assumes you have a Coupon model that closely mirrors the structure of your database table for coupons.
func (req *CouponRequest) ToCouponModel() model.Coupon {
	return model.Coupon{
		DiscountExpiration: req.ExpirationDate,
		DiscountPercentage: req.DiscountPercentage,
		ApplicableTourId:   req.TourID,
		CouponIssuerId:     req.UserID,
		// Initialize other fields as needed, potentially with default values or derived values.
	}
}
