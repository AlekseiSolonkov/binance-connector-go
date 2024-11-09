package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetFlexibleLoanOngoingOrders //

// Flexible loan repay API Endpoint
const (
	getSimpleEarnFlexibleRedeemEndpoint = "/sapi/v1/simple-earn/flexible/redeem"
)

type SimpleEarnFlexibleRedeem struct {
	c         *Client
	productId *string
}

func (s *SimpleEarnFlexibleRedeem) ProductId(productId string) *SimpleEarnFlexibleRedeem {
	s.productId = &productId
	return s
}

type SimpleEarnFlexibleRedeemResponse struct {
	RedeemId int64 `json:"redeemId"`
	Success  bool  `json:"success"`
}

// Do send request
func (s *SimpleEarnFlexibleRedeem) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexibleRedeemResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: getSimpleEarnFlexibleRedeemEndpoint,
		secType:  secTypeSigned,
	}
	if s.productId != nil {
		r.setParam("productId", *s.productId)
	}
	r.setParam("redeemAll", true)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SimpleEarnFlexibleRedeemResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
