package service

import (
	"database-example/model"
	"database-example/repo"
	"fmt"
	"math/rand"
)

type CouponService struct {
	CouponRepo *repo.CouponRepository
}

func (service *CouponService) FindCouponByHash(hash string) (*model.Coupon, error) {
	coupon, err := service.CouponRepo.FindByHash(hash)
	if err != nil {
		return nil, fmt.Errorf("coupon with hash %s not found: %v", hash, err)
	}
	return coupon, nil
}

func (service *CouponService) CreateCoupon(coupon *model.Coupon) error {
	coupon.CouponHash = CreateRandomCouponHash()
	fmt.Println("napraviohash")
	err := service.CouponRepo.Create(coupon)
	if err != nil {
		return fmt.Errorf("failed to create coupon: %v", err)
	}
	return nil
}

func (service *CouponService) RemoveCoupon(hash string, userId int) error {
	coupon, _ := service.FindCouponByHash(hash)
	if coupon.CouponIssuerId != userId {
		return fmt.Errorf("User not allowed")
	}
	err := service.CouponRepo.Remove(coupon)
	if err != nil {
		return fmt.Errorf("failed to delete coupon: %v", err)
	}
	return nil
}

func (service *CouponService) UpdateCoupon(coupon *model.Coupon) error {
	err := service.CouponRepo.Update(coupon)
	if err != nil {
		return fmt.Errorf("failed to update coupon: %v", err)
	}
	return nil
}

func CreateRandomCouponHash() string {
	const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 8)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
