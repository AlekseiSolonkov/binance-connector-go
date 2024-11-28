package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetFlexibleLoanOngoingOrders //

// Flexible loan repay API Endpoint
const (
	getFlexibleLoanAssetsDataEndpoint = "/sapi/v2/loan/flexible/loanable/data"
)

type GetFlexibleLoanAssetsData struct {
	c        *Client
	loanCoin *string
}

func (s *GetFlexibleLoanAssetsData) LoanCoin(loanCoin string) *GetFlexibleLoanAssetsData {
	s.loanCoin = &loanCoin
	return s
}

type GetFlexibleLoanAssetsDataResponse struct {
	Total int                 `json:"total"`
	Rows  []LoanAssetsDataRow `json:"rows"`
}

type LoanAssetsDataRow struct {
	LoanCoin             string `json:"loanCoin"`
	FlexibleInterestRate string `json:"flexibleInterestRate"`
	FlexibleMinLimit     string `json:"flexibleMinLimit"`
	FlexibleMaxLimit     string `json:"flexibleMaxLimit"`
}

// Do send request
func (s *GetFlexibleLoanAssetsData) Do(ctx context.Context, opts ...RequestOption) (res *GetFlexibleLoanAssetsDataResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFlexibleLoanAssetsDataEndpoint,
		secType:  secTypeSigned,
	}
	if s.loanCoin != nil {
		r.setParam("loanCoin", *s.loanCoin)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetFlexibleLoanAssetsDataResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
