package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Flexible loan rate history API Endpoint
const (
	getFlexibleLoanRateHistoryEndpoint = "/bapi/margin/v1/public/flexibleLoan/isolated/new/loanData/annualInterestRates"
)

type FlexibleLoanRateHistoryService struct {
	c         *Client
	coin      *string
	size      *int64
	startTime *int64
	endTime   *int64
	current   *int64
}

func (s *FlexibleLoanRateHistoryService) Coin(coin string) *FlexibleLoanRateHistoryService {
	s.coin = &coin
	return s
}

func (s *FlexibleLoanRateHistoryService) Size(size int64) *FlexibleLoanRateHistoryService {
	s.size = &size
	return s
}

func (s *FlexibleLoanRateHistoryService) StartTime(startTime int64) *FlexibleLoanRateHistoryService {
	s.startTime = &startTime
	return s
}

func (s *FlexibleLoanRateHistoryService) EndTime(endTime int64) *FlexibleLoanRateHistoryService {
	s.endTime = &endTime
	return s
}

func (s *FlexibleLoanRateHistoryService) Current(current int64) *FlexibleLoanRateHistoryService {
	s.current = &current
	return s
}

type RateHistoryRecord struct {
	Total int64                        `json:"total"`
	Rows  []FlexibleLoanRateHistoryRow `json:"rows,omitempty"`
}

type FlexibleLoanRateHistoryResponse struct {
	Data RateHistoryRecord `json:"data"`
}

type FlexibleLoanRateHistoryRow struct {
	Asset              string `json:"coin"`
	Time               string `json:"dateTimestamp"`
	AnnualInterestRate string `json:"annualInterestRate"`
}

// Do send request
func (s *FlexibleLoanRateHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *FlexibleLoanRateHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getFlexibleLoanRateHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.coin != nil {
		r.setParam("coin", *s.coin)
	}
	if s.size != nil {
		r.setParam("size", *s.size)
	}
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.current != nil {
		r.setParam("current", *s.current)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(FlexibleLoanRateHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
