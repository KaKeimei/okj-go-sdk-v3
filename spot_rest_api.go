package okcoin

import "strings"

type AccountInfo struct {
	Hold      string `json:"hold"`
	Currency  string `json:"currency"`
	Balance   string `json:"balance"`
	Available string `json:"available"`
}

/*
GetSpotAccounts
Account Information
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/accounts
*/
func (client *Client) GetSpotAccounts() ([]AccountInfo, error) {
	var r []AccountInfo
	if _, err := client.Request(GET, SPOT_ACCOUNTS, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

/*
GetSpotAccountsByCurrency
Get Currency
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/accounts/<currency>
*/
func (client *Client) GetSpotAccountsByCurrency(currency string) (*AccountInfo, error) {
	r := AccountInfo{}
	uri := GetCurrencyUri(SPOT_ACCOUNTS_CURRENCY, currency)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type OrderInfo struct {
	ClientOid      string `json:"client_oid"`
	CreatedAt      string `json:"created_at"`
	FilledNotional string `json:"filled_notional"`
	FilledSize     string `json:"filled_size"`
	Funds          string `json:"funds"`
	InstrumentID   string `json:"instrument_id"`
	Notional       string `json:"notional"`
	OrderID        string `json:"order_id"`
	OrderType      string `json:"order_type"`
	Price          string `json:"price"`
	PriceAvg       string `json:"price_avg"`
	ProductID      string `json:"product_id"`
	Side           string `json:"side"`
	Size           string `json:"size"`
	Status         string `json:"status"`
	State          string `json:"state"`
	Timestamp      string `json:"timestamp"`
	Type           string `json:"type"`
}

/*
GetSpotOrders
Order List
Limit: 10 requests per 2 seconds
GET Request: /api/spot/v3/orders
*/
func (client *Client) GetSpotOrders(instrument_id, state string, options *map[string]string) ([]OrderInfo, error) {
	var r []OrderInfo

	fullOptions := NewParams()
	fullOptions["instrument_id"] = instrument_id
	fullOptions["state"] = state
	if options != nil && len(*options) > 0 {
		fullOptions["before"] = (*options)["before"]
		fullOptions["after"] = (*options)["after"]
		fullOptions["limit"] = (*options)["limit"]
	}

	uri := BuildParams(SPOT_ORDERS, fullOptions)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

/*
GetSpotOrdersPending
Open Orders
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/orders_pending
*/
func (client *Client) GetSpotOrdersPending(instrumentId string, options *map[string]string) ([]OrderInfo, error) {
	var r []OrderInfo

	fullOptions := NewParams()
	fullOptions["instrument_id"] = instrumentId

	if options != nil && len(*options) > 0 {

		for k, v := range *options {
			if v != "" && len(v) > 0 {
				fullOptions[k] = v
			}
		}
	}

	uri := BuildParams(SPOT_ORDERS_PENDING, fullOptions)
	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}

	return r, nil
}

/*
GetSpotOrdersById
Order Details
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/orders/<order_id> or <client_oid>
*/
func (client *Client) GetSpotOrdersById(instrumentId, orderOrClientId string) (*OrderInfo, error) {
	r := OrderInfo{}
	uri := strings.Replace(SPOT_ORDERS_BY_ID, "{order_client_id}", orderOrClientId, -1)
	options := NewParams()
	options["instrument_id"] = instrumentId
	uri = BuildParams(uri, options)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type FilledOrder struct {
	CreatedAt    string `json:"created_at"`
	TradeID      string `json:"trade_id"`
	Currency     string `json:"currency"`
	ExecType     string `json:"exec_type"`
	Fee          string `json:"fee"`
	InstrumentID string `json:"instrument_id"`
	LedgerID     string `json:"ledger_id"`
	Liquidity    string `json:"liquidity"`
	OrderID      string `json:"order_id"`
	Price        string `json:"price"`
	ProductID    string `json:"product_id"`
	Side         string `json:"side"`
	Size         string `json:"size"`
	Timestamp    string `json:"timestamp"`
}

/*
GetSpotFills
Transaction Details
Limit: 10 requests per 2 seconds
GET Request: /api/spot/v3/fills
*/
func (client *Client) GetSpotFills(instrument_id string, options *map[string]string) ([]FilledOrder, error) {
	var r []FilledOrder

	fullOptions := NewParams()
	fullOptions["instrument_id"] = instrument_id
	if options != nil && len(*options) > 0 {
		fullOptions["order_id"] = (*options)["order_id"]
		fullOptions["before"] = (*options)["before"]
		fullOptions["after"] = (*options)["after"]
		fullOptions["limit"] = (*options)["limit"]
	}

	uri := BuildParams(SPOT_FILLS, fullOptions)

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type TradeFee struct {
	Maker     string `json:"maker"`
	Taker     string `json:"taker"`
	Timestamp string `json:"timestamp"`
}

/*
GetSpotTradeFee
Limit: 1 requests per 10 seconds
GET Request: /api/spot/v3/trade_fee
*/
func (client *Client) GetSpotTradeFee() ([]TradeFee, error) {
	var r []TradeFee

	if _, err := client.Request(GET, SPOT_TRADE_FEE, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type Instrument struct {
	BaseCurrency  string `json:"base_currency"`
	InstrumentID  string `json:"instrument_id"`
	MinSize       string `json:"min_size"`
	QuoteCurrency string `json:"quote_currency"`
	SizeIncrement string `json:"size_increment"`
	TickSize      string `json:"tick_size"`
}

/*
GetSpotInstruments
Public - Trading Pairs
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments
*/
func (client *Client) GetSpotInstruments() ([]Instrument, error) {
	var r []Instrument

	if _, err := client.Request(GET, SPOT_INSTRUMENTS, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type OrderBook struct {
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
	Timestamp string     `json:"timestamp"`
}

/*
GetSpotInstrumentBook
Public - Order Book
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/book
*/
func (client *Client) GetSpotInstrumentBook(instrumentId string, optionalParams *map[string]string) (*OrderBook, error) {
	r := OrderBook{}
	uri := GetInstrumentIdUri(SPOT_INSTRUMENT_BOOK, instrumentId)
	if optionalParams != nil && len(*optionalParams) > 0 {
		optionals := NewParams()
		optionals["size"] = (*optionalParams)["size"]
		optionals["depth"] = (*optionalParams)["depth"]
		uri = BuildParams(uri, optionals)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type Ticker struct {
	BestAsk        string `json:"best_ask"`
	BestBid        string `json:"best_bid"`
	InstrumentID   string `json:"instrument_id"`
	ProductID      string `json:"product_id"`
	Last           string `json:"last"`
	LastQty        string `json:"last_qty"`
	Ask            string `json:"ask"`
	BestAskSize    string `json:"best_ask_size"`
	Bid            string `json:"bid"`
	BestBidSize    string `json:"best_bid_size"`
	Open24h        string `json:"open_24h"`
	High24h        string `json:"high_24h"`
	Low24h         string `json:"low_24h"`
	BaseVolume24h  string `json:"base_volume_24h"`
	Timestamp      string `json:"timestamp"`
	QuoteVolume24h string `json:"quote_volume_24h"`
}

/*
GetSpotInstrumentsTicker
Public - Ticker
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/ticker
*/
func (client *Client) GetSpotInstrumentsTicker() ([]Ticker, error) {
	var r []Ticker

	if _, err := client.Request(GET, SPOT_INSTRUMENTS_TICKER, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

/*
GetSpotInstrumentTicker
Public - Trading Pair Information
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/ticker
*/
func (client *Client) GetSpotInstrumentTicker(instrument_id string) (*Ticker, error) {
	r := Ticker{}

	uri := GetInstrumentIdUri(SPOT_INSTRUMENT_TICKER, instrument_id)
	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

type Trade struct {
	Time      string `json:"time"`
	Timestamp string `json:"timestamp"`
	TradeID   string `json:"trade_id"`
	Price     string `json:"price"`
	Size      string `json:"size"`
	Side      string `json:"side"`
}

/*
GetSpotInstrumentTrade
Public - Filled Orders
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/trades
*/
func (client *Client) GetSpotInstrumentTrade(instrument_id string, options *map[string]string) ([]Trade, error) {
	var r []Trade

	uri := GetInstrumentIdUri(SPOT_INSTRUMENT_TRADES, instrument_id)
	fullOptions := NewParams()
	if options != nil && len(*options) > 0 {
		fullOptions["limit"] = (*options)["limit"]
		uri = BuildParams(uri, fullOptions)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

/*
GetSpotInstrumentCandles
Public - Market Data
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/candles
*/
func (client *Client) GetSpotInstrumentCandles(instrument_id string, options *map[string]string) ([][]string, error) {
	var r [][]string

	uri := GetInstrumentIdUri(SPOT_INSTRUMENT_CANDLES, instrument_id)
	fullOptions := NewParams()
	if options != nil && len(*options) > 0 {
		fullOptions["start"] = (*options)["start"]
		fullOptions["end"] = (*options)["end"]
		fullOptions["granularity"] = (*options)["granularity"]
		uri = BuildParams(uri, fullOptions)
	}

	if _, err := client.Request(GET, uri, nil, &r); err != nil {
		return nil, err
	}
	return r, nil
}

type OrderPostResp struct {
	ClientOID    string `json:"client_oid"`
	ErrorMessage string `json:"error_message"`
	ErrorCode    string `json:"error_code"`
	OrderID      string `json:"order_id"`
	Result       bool   `json:"result"`
}

/*
PostSpotOrders
Place Order
Limit: 100 requests per 2 seconds
POST Request: /api/spot/v3/orders
*/
func (client *Client) PostSpotOrders(instrument_id, side string, optionalOrderInfo *map[string]string) (result *OrderPostResp, err error) {

	r := OrderPostResp{}
	postParams := NewParams()
	postParams["side"] = side
	postParams["instrument_id"] = instrument_id

	if optionalOrderInfo != nil && len(*optionalOrderInfo) > 0 {

		for k, v := range *optionalOrderInfo {
			postParams[k] = v
		}

		if postParams["type"] == "limit" {
			postParams["price"] = (*optionalOrderInfo)["price"]
			postParams["size"] = (*optionalOrderInfo)["size"]

		} else if postParams["type"] == "market" {
			postParams["size"] = (*optionalOrderInfo)["size"]
			postParams["notional"] = (*optionalOrderInfo)["notional"]

		}
	}

	if _, err := client.Request(POST, SPOT_ORDERS, postParams, &r); err != nil {
		return nil, err
	}

	return &r, nil
}

/*
PostSpotBatchOrders
Batch Orders
Limit: 50 requests per 2 seconds
POST Request: /api/spot/v3/batch_orders
*/
func (client *Client) PostSpotBatchOrders(orderInfos *[]map[string]string) (*map[string][]OrderPostResp, error) {
	r := map[string][]OrderPostResp{}
	if _, err := client.Request(POST, SPOT_BATCH_ORDERS, orderInfos, &r); err != nil {
		return nil, err
	}
	return &r, nil
}

/*
PostSpotCancelOrders
Cancel Order
Limit: 100 requests per 2 seconds
POST Request: /api/spot/v3/cancel_orders/<order_id> or <client_oid>
*/
func (client *Client) PostSpotCancelOrders(instrumentId, orderOrClientId string) (*OrderPostResp, error) {
	r := OrderPostResp{}

	uri := strings.Replace(SPOT_CANCEL_ORDERS_BY_ID, "{order_client_id}", orderOrClientId, -1)
	options := NewParams()
	options["instrument_id"] = instrumentId

	if _, err := client.Request(POST, uri, options, &r); err != nil {
		return nil, err
	}
	return &r, nil

}

/*
PostSpotCancelBatchOrders
Cancel Multiple Orders
Limit: 20 requests per 2 seconds
POST Request: /api/spot/v3/cancel_batch_orders
*/
func (client *Client) PostSpotCancelBatchOrders(orderInfos *[]map[string]interface{}) (*map[string][]OrderPostResp, error) {
	r := map[string][]OrderPostResp{}
	if _, err := client.Request(POST, SPOT_CANCEL_BATCH_ORDERS, orderInfos, &r); err != nil {
		return nil, err
	}
	return &r, nil
}
