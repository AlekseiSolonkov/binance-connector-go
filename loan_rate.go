package binance_connector

import (
	"context"
	"encoding/json"
	"net/http"
)

// Flexible loan rate history API Endpoint
const (
	getSimpleEarnFlexibleRateHistoryEndpoint = "/sapi/v1/simple-earn/flexible/history/rateHistory"
)

type SimpleEarnFlexibleRateHistoryService struct {
	c         *Client
	productId *string
	startTime *int64
	endTime   *int64
	current   *int64
}

func (s *SimpleEarnFlexibleRateHistoryService) ProductId(productId string) *SimpleEarnFlexibleRateHistoryService {
	s.productId = &productId
	return s
}

func (s *SimpleEarnFlexibleRateHistoryService) StartTime(startTime int64) *SimpleEarnFlexibleRateHistoryService {
	s.startTime = &startTime
	return s
}

func (s *SimpleEarnFlexibleRateHistoryService) EndTime(endTime int64) *SimpleEarnFlexibleRateHistoryService {
	s.endTime = &endTime
	return s
}

func (s *SimpleEarnFlexibleRateHistoryService) Current(current int64) *SimpleEarnFlexibleRateHistoryService {
	s.current = &current
	return s
}

type SimpleEarnFlexibleRateHistoryResponse struct {
	Total int64                              `json:"total"`
	Rows  []SimpleEarnFlexibleRateHistoryRow `json:"rows,omitempty"`
}

type SimpleEarnFlexibleRateHistoryRow struct {
	Asset                string `json:"asset"`
	AnnualPercentageRate string `json:"annualPercentageRate"`
	ProductId            string `json:"productId"`
	Time                 int64  `json:"time"`
}

// Do send request
func (s *SimpleEarnFlexibleRateHistoryService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnFlexibleRateHistoryResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: getSimpleEarnFlexibleRateHistoryEndpoint,
		secType:  secTypeSigned,
	}
	if s.productId != nil {
		r.setParam("productId", *s.productId)
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
	res = new(SimpleEarnFlexibleRateHistoryResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
