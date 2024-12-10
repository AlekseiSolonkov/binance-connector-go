package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetFlexibleLoanOngoingOrders //

// Flexible loan repay API Endpoint
const (
	getFlexibleLoanAdjustEndpoint = "/sapi/v2/loan/flexible/adjust/ltv"
)

// FlexibleLoanAdjustService adjust flexible loan LTV
type FlexibleLoanAdjustService struct {
	c                *Client
	loanCoin         *string
	collateralCoin   *string
	adjustmentAmount *float64
	direction        *string
}

func (s *FlexibleLoanAdjustService) LoanCoin(loanCoin string) *FlexibleLoanAdjustService {
	s.loanCoin = &loanCoin
	return s
}

func (s *FlexibleLoanAdjustService) CollateralCoin(collateralCoin string) *FlexibleLoanAdjustService {
	s.collateralCoin = &collateralCoin
	return s
}

func (s *FlexibleLoanAdjustService) AdjustmentAmount(adjustmentAmount float64) *FlexibleLoanAdjustService {
	s.adjustmentAmount = &adjustmentAmount
	return s
}

func (s *FlexibleLoanAdjustService) Direction(direction string) *FlexibleLoanAdjustService {
	s.direction = &direction
	return s
}

type FlexibleLoanAdjustResponse struct {
	LoanCoin       string `json:"loanCoin"`
	CollateralCoin string `json:"collateralCoin"`
	Direction      string `json:"direction"`
	Amount         string `json:"amount"`
	CurrentLTV     string `json:"currentLTV"`
	Status         string `json:"status"`
}

// Do send request
func (s *FlexibleLoanAdjustService) Do(ctx context.Context, opts ...RequestOption) (res *FlexibleLoanAdjustResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: getFlexibleLoanAdjustEndpoint,
		secType:  secTypeSigned,
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	if s.collateralCoin != nil {
		r.setParam("collateralCoin", *s.collateralCoin)
	}
	if s.adjustmentAmount != nil {
		r.setParam("adjustmentAmount", *s.adjustmentAmount)
	}
	if s.direction != nil {
		r.setParam("direction", *s.direction)
	}

	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FlexibleLoanAdjustResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
