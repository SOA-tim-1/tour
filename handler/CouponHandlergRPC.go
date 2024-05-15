package handler

import (
	"context"
	"database-example/model"
	"database-example/proto/coupon"
	"database-example/service"

	"github.com/golang/protobuf/ptypes"
)

type CouponHandlergRPC struct {
	CouponService *service.CouponService
	coupon.UnimplementedCouponServiceServer
}

func (handler *CouponHandlergRPC) CreateCoupon(ctx context.Context, in *coupon.Coupon) (*coupon.CreateCouponResponse, error) {

	createdCoupon := ConvertCouponResponseToCoupon(in)
	err := handler.CouponService.CreateCoupon(createdCoupon)
	if err != nil {
		return nil, err
	}

	return &coupon.CreateCouponResponse{}, nil
}

func (handler *CouponHandlergRPC) RemoveCoupon(ctx context.Context, in *coupon.RemoveCouponRequest) (*coupon.RemoveCouponResponse, error) {

	err := handler.CouponService.RemoveCoupon(in.GetHash(), int(in.GetUserId()))
	if err != nil {
		return nil, err
	}

	return &coupon.RemoveCouponResponse{}, nil
}

func ConvertCouponToCouponResponse(couponModel *model.Coupon) *coupon.Coupon {

	discountExpiration, _ := ptypes.TimestampProto(couponModel.DiscountExpiration)

	couponResponse := &coupon.Coupon{
		Id:                         couponModel.ID,
		CouponHash:                 couponModel.CouponHash,
		DiscountPercentage:         float32(couponModel.DiscountPercentage),
		DiscountExpiration:         discountExpiration,
		ApplicableTourId:           int32(*couponModel.ApplicableTourId),
		CouponIssuerId:             int32(couponModel.CouponIssuerId),
		IsApplicableToAllUserTours: *couponModel.IsApplicableToAllUserTours,
		IsValid:                    couponModel.IsValid,
	}

	return couponResponse
}

func ConvertCouponResponseToCoupon(couponResponse *coupon.Coupon) *model.Coupon {
	discountExpiration, _ := ptypes.Timestamp(couponResponse.DiscountExpiration)

	applicableTourId := int(couponResponse.ApplicableTourId)

	couponModel := &model.Coupon{
		ID:                         couponResponse.Id,
		CouponHash:                 couponResponse.CouponHash,
		DiscountPercentage:         float64(couponResponse.DiscountPercentage),
		DiscountExpiration:         discountExpiration,
		ApplicableTourId:           &applicableTourId,
		CouponIssuerId:             int(couponResponse.CouponIssuerId),
		IsApplicableToAllUserTours: &couponResponse.IsApplicableToAllUserTours,
		IsValid:                    couponResponse.IsValid,
	}

	return couponModel
}
