package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// GetFlexibleLoanOngoingOrders //

// Flexible loan repay API Endpoint
const (
	getFlexibleLoanOngoingOrdersEndpoint = "/sapi/v2/loan/flexible/ongoing/orders"
)

type GetFlexibleLoanOngoingOrders struct {
	c              *Client
	loanCoin       *string
	collateralCoin *string
}

func (s *GetFlexibleLoanOngoingOrders) LoanCoin(loanCoin string) *GetFlexibleLoanOngoingOrders {
	s.loanCoin = &loanCoin
	return s
}

func (s *GetFlexibleLoanOngoingOrders) CollateralCoin(collateralCoin string) *GetFlexibleLoanOngoingOrders {
	s.collateralCoin = &collateralCoin
	return s
}

type GetFlexibleLoanOngoingOrdersResponse struct {
	Total int                    `json:"total"`
	Rows  []LoanOngoingOrdersRow `json:"rows"`
}

type LoanOngoingOrdersRow struct {
	LoanCoin         string `json:"loanCoin"`
	CollateralCoin   string `json:"collateralCoin"`
	TotalDebt        string `json:"totalDebt"`
	CollateralAmount string `json:"collateralAmount"`
	CurrentLTV       string `json:"currentLTV"`
}

// Do send request
func (s *GetFlexibleLoanOngoingOrders) Do(ctx context.Context, opts ...RequestOption) (res *GetFlexibleLoanOngoingOrdersResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFlexibleLoanOngoingOrdersEndpoint,
		secType:  secTypeSigned,
	}
	m := params{}
	if s.loanCoin != nil {
		m["loanCoin"] = *s.loanCoin
	}
	if s.collateralCoin != nil {
		m["collateralCoin"] = *s.collateralCoin
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(GetFlexibleLoanOngoingOrdersResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Flexible loan repay //

// Flexible loan repay API Endpoint
const (
	flexibleLoanBorrowEndpoint = "/sapi/v2/loan/flexible/borrow"
)

// FlexibleLoanBorrowService borrow flexible loan
type FlexibleLoanBorrowService struct {
	c                *Client
	loanCoin         *string
	loanAmount       *float64 // Mandatory when collateralAmount is empty
	collateralCoin   *string
	collateralAmount *float64 // Mandatory when loanAmount is empty
}

func (s *FlexibleLoanBorrowService) LoanCoin(loanCoin string) *FlexibleLoanBorrowService {
	s.loanCoin = &loanCoin
	return s
}

func (s *FlexibleLoanBorrowService) LoanAmount(loanAmount float64) *FlexibleLoanBorrowService {
	s.loanAmount = &loanAmount
	return s
}

func (s *FlexibleLoanBorrowService) CollateralCoin(collateralCoin string) *FlexibleLoanBorrowService {
	s.collateralCoin = &collateralCoin
	return s
}

func (s *FlexibleLoanBorrowService) CollateralAmount(collateralAmount float64) *FlexibleLoanBorrowService {
	s.collateralAmount = &collateralAmount
	return s
}

type FlexibleLoanBorrowResponse struct {
	CollateralCoin   string `json:"collateralCoin"`
	LoanCoin         string `json:"loanCoin"`
	LoanAmount       string `json:"loanAmount"`
	CollateralAmount string `json:"collateralAmount"`
	Status           string `json:"status"`
}

// Do send request
func (s *FlexibleLoanBorrowService) Do(ctx context.Context, opts ...RequestOption) (res *FlexibleLoanBorrowResponse, err error) {
	r := &request{
		method:   http.MethodPost,
		endpoint: flexibleLoanBorrowEndpoint,
		secType:  secTypeSigned,
	}
	m := params{}
	if s.loanCoin != nil {
		m["loanCoin"] = *s.loanCoin
	}
	if s.loanAmount != nil {
		m["loanAmount"] = *s.loanAmount
	}
	if s.collateralCoin != nil {
		m["collateralCoin"] = *s.collateralCoin
	}
	if s.collateralAmount != nil {
		m["collateralAmount"] = *s.collateralAmount
	}
	r.setParams(m)
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FlexibleLoanBorrowResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}