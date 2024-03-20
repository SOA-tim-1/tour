package handler

import (
	"database-example/dtos"
	"database-example/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CouponHandler struct {
	CouponService *service.CouponService
}

func (h *CouponHandler) CreateCoupon(writer http.ResponseWriter, req *http.Request) {
	var couponReq dtos.CouponRequest

	if err := json.NewDecoder(req.Body).Decode(&couponReq); err != nil {
		http.Error(writer, "Error parsing request body", http.StatusBadRequest)
		return
	}

	coupon := couponReq.ToCouponModel()
	if err := h.CouponService.CreateCoupon(&coupon); err != nil {
		http.Error(writer, "Error creating coupon", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	json.NewEncoder(writer).Encode(coupon)
}

func (h *CouponHandler) DeleteCoupon(writer http.ResponseWriter, req *http.Request) {
	couponHash := req.URL.Query().Get("couponHash")
	userId := req.URL.Query().Get("userId")
	num, _ := strconv.Atoi(userId)

	fmt.Print(couponHash + userId) // Adjusted for clarity in logging.

	if err := h.CouponService.RemoveCoupon(couponHash, num); err != nil {
		http.Error(writer, "Error deleting coupon", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
}

/*// UpdateCoupon handles coupon updates.
func (h *CouponHandler) UpdateCoupon(writer http.ResponseWriter, req *http.Request) {
	couponHash := mux.Vars(req)["couponHash"] // Assuming you're using the couponHash as a URL parameter

	var couponReq model.CouponRequest
	if err := json.NewDecoder(req.Body).Decode(&couponReq); err != nil {
		http.Error(writer, "Error parsing request body", http.StatusBadRequest)
		return
	}

	coupon := couponReq.ToCouponModel() // Convert request to model
	coupon.CouponHash = couponHash      // Ensure we're updating the correct coupon

	if err := h.CouponService.UpdateCoupon(&coupon); err != nil {
		http.Error(writer, "Error updating coupon", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(coupon) // Sending back the updated coupon
}
*/
// Assuming there's a method to convert CouponRequest to Coupon model in your model definitions.
// You might need to adjust model and method names based on your actual application structure.
