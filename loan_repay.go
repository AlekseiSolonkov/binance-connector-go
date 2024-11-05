package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Flexible loan repay API Endpoint
const (
	flexibleLoanRepayEndpoint = "/sapi/v2/loan/flexible/repay"
)

// FlexibleLoanRepayService repay flexible loan
type FlexibleLoanRepayService struct {
	c                *Client
	loanCoin         *string
	collateralCoin   *string
	repayAmount      *float64
	collateralReturn *bool
	fullRepayment    *bool
}

// LoanCoin set loanCoin
func (s *FlexibleLoanRepayService) LoanCoin(loanCoin string) *FlexibleLoanRepayService {
	s.loanCoin = &loanCoin
	return s
}

func (s *FlexibleLoanRepayService) CollateralCoin(collateralCoin string) *FlexibleLoanRepayService {
	s.collateralCoin = &collateralCoin
	return s
}

func (s *FlexibleLoanRepayService) RepayAmount(repayAmount float64) *FlexibleLoanRepayService {
	s.repayAmount = &repayAmount
	return s
}

func (s *FlexibleLoanRepayService) FullRepayment(fullRepayment bool) *FlexibleLoanRepayService {
	s.fullRepayment = &fullRepayment
	return s
}

func (s *FlexibleLoanRepayService) CollateralReturn(collateralReturn bool) *FlexibleLoanRepayService {
	s.collateralReturn = &collateralReturn
	return s
}

type FlexibleLoanRepayResponse struct {
	LoanCoin            string `json:"loanCoin"`
	CollateralCoin      string `json:"collateralCoin"`
	RemainingDebt       string `json:"remainingDebt"`
	RemainingCollateral string `json:"remainingCollateral"`
	CurrentLTV          string `json:"currentLTV"`
	RepayStatus         string `json:"repayStatus"`
	FullRepayment       bool   `json:"fullRepayment"`
}

// Do send request
func (s *FlexibleLoanRepayService) Do(ctx context.Context, opts ...RequestOption) (res *FlexibleLoanRepayResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: flexibleLoanRepayEndpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"fullRepayment":    false,
		"collateralReturn": true,
	}
	if s.loanCoin != nil {
		m["loanCoin"] = *s.loanCoin
	}
	if s.collateralCoin != nil {
		m["collateralCoin"] = *s.collateralCoin
	}
	if s.repayAmount != nil {
		m["repayAmount"] = *s.repayAmount
	}
	if s.collateralReturn != nil {
		m["collateralReturn"] = *s.collateralReturn
	}
	if s.fullRepayment != nil {
		m["fullRepayment"] = *s.fullRepayment
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FlexibleLoanRepayResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
