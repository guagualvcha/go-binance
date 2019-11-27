package binance

import (
	"context"
	"encoding/json"
)

// GetCoinService get account info
type GetCoinService struct {
	c *Client
}

// Do send request
func (s *GetCoinService) Do(ctx context.Context, opts ...RequestOption) (res []Coin, err error) {
	r := &request{
		method:   "GET",
		endpoint: "/sapi/v1/capital/config/getall",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = make([]Coin, 0, 200)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Coin define Coin info
type Coin struct {
	Coin        string    `json:"coin"`
	NetworkList []Network `json:"networkList"`
}

type Network struct {
	Network        string `json:"network"`
	WithdrawEnable bool   `json:"withdrawEnable"`
	DepositEnable  bool   `json:"depositEnable"`
}
