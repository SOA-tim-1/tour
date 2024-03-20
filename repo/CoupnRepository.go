package repo

import (
	"database-example/model"
	"errors"
	"gorm.io/gorm"
)

type CouponRepository struct {
	DatabaseConnection *gorm.DB
}

func (repo *CouponRepository) FindByHash(hash string) (*model.Coupon, error) {
	var coupon model.Coupon
	dbResult := repo.DatabaseConnection.First(&coupon, "coupon_hash = ?", hash)
	if dbResult.Error != nil {
		return &model.Coupon{}, dbResult.Error
	}
	return &coupon, nil
}

func (repo *CouponRepository) Create(coupon *model.Coupon) error {
	if result := repo.DatabaseConnection.Create(coupon); result.Error != nil {
		return errors.New("unable to save coupon: " + result.Error.Error())
	}
	return nil
}

func (repo *CouponRepository) Remove(coupon *model.Coupon) error {
	if result := repo.DatabaseConnection.Delete(coupon); result.Error != nil {
		return errors.New("unable to delete coupon: " + result.Error.Error())
	}
	return nil
}

func (repo *CouponRepository) Update(coupon *model.Coupon) error {
	if result := repo.DatabaseConnection.Save(coupon); result.Error != nil {
		return errors.New("unable to update coupon: " + result.Error.Error())
	}
	return nil
}
