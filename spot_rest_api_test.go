package okcoin

import (
	"testing"
)

/*
Account Information
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/accounts
*/
func TestGetSpotAccounts(t *testing.T) {
	NewTestClient().GetSpotAccounts()
}

/*
Get Currency
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/accounts/<currency>
*/
func TestGetSpotAccountsByCurrency(t *testing.T) {
	NewTestClient().GetSpotAccountsByCurrency("BTC")
}

/*
Place Order
Limit: 100 requests per 2 seconds
POST Request: /api/spot/v3/orders
*/
func TestPostSpotOrders(t *testing.T) {
	options := NewParams()
	options["client_oid"] = ""
	options["order_type"] = "0"
	options["type"] = "market"
	options["price"] = ""
	options["size"] = "0.001"
	options["notional"] = "0"
	NewTestClient().PostSpotOrders("BTC-JPY", "sell", &options)
}

/*
Batch Orders
Limit: 50 requests per 2 seconds
POST Request: /api/spot/v3/batch_orders
*/
func TestClient_PostSpotBatchOrders(t *testing.T) {
	orderInfos := []map[string]string{
		map[string]string{"client_oid": "", "instrument_id": "BTC-JPY", "side": "buy", "type": "limit", "size": "0.001", "price": "", "order_type": "0", "notional": "0"},
		map[string]string{"client_oid": "", "instrument_id": "BTC-JPY", "side": "buy", "type": "limit", "size": "0.001", "price": "", "order_type": "0", "notional": "0"},
	}
	NewTestClient().PostSpotBatchOrders(&orderInfos)
}

/*
Cancel Order
Limit: 100 requests per 2 seconds
POST Request: /api/spot/v3/cancel_orders/<order_id> or <client_oid>
*/
func TestPostSpotCancelOrders(t *testing.T) {
	NewTestClient().PostSpotCancelOrders("BTC-JPY", "")
}

/*
Cancel Multiple Orders
Limit: 20 requests per 2 seconds
POST Request: /api/spot/v3/cancel_batch_orders
*/
func TestClient_PostSpotCancelBatchOrders(t *testing.T) {
	orderInfos := []map[string]interface{}{
		map[string]interface{}{"instrument_id": "BTC-JPY", "client_oids": []string{"", ""}},
		map[string]interface{}{"instrument_id": "BTC-JPY", "order_ids": []string{"", ""}},
	}
	NewTestClient().PostSpotCancelBatchOrders(&orderInfos)
}

/*
Order List
Limit: 10 requests per 2 seconds
GET Request: /api/spot/v3/orders
*/
func TestGetSpotOrders(t *testing.T) {
	options := map[string]string{}
	options["after"] = ""
	options["before"] = ""
	options["limit"] = ""
	NewTestClient().GetSpotOrders("ETC-JPY", "2", &options)
}

/*
Open Orders
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/orders_pending
*/
func TestGetSpotOrdersPending(t *testing.T) {
	options := map[string]string{}
	options["after"] = ""
	options["before"] = ""
	options["limit"] = "1"
	NewTestClient().GetSpotOrdersPending("BTC-JPY", &options)
}

/*
Order Details
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/orders/<order_id> or <client_oid>
*/
func TestGetSpotOrdersById(t *testing.T) {
	NewTestClient().GetSpotOrdersById("BTC-JPY", "")
}

/*
Transaction Details
Limit: 10 requests per 2 seconds
GET Request: /api/spot/v3/fills
*/
func TestGetSpotFills(t *testing.T) {
	options := map[string]string{}
	options["order_id"] = ""
	options["after"] = ""
	options["before"] = ""
	options["limit"] = ""
	NewTestClient().GetSpotFills("BTC-JPY", &options)
}

/*
Trade Fee
Limit: 1 requests per 10 seconds
GET Request: /api/spot/v3/trade_fee
*/
func TestGetSpotTradeFee(t *testing.T) {
	NewTestClient().GetSpotTradeFee()
}

/*
Public - Trading Pairs
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments
*/
func TestGetSpotInstruments(t *testing.T) {
	NewTestClient().GetSpotInstruments()
}

/*
Public - Order Book
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/book
*/
func TestGetSpotInstrumentBook(t *testing.T) {
	NewTestClient().GetSpotInstrumentBook("BTC-JPY", nil)

	options := map[string]string{}
	options["size"] = "1"
	options["depth"] = ""
	NewTestClient().GetSpotInstrumentBook("BTC-JPY", &options)
}

/*
Public - Ticker
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/ticker
*/
func TestGetSpotInstrumentsTicker(t *testing.T) {
	NewTestClient().GetSpotInstrumentsTicker()
}

/*
Public - Trading Pair Information
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/ticker
*/
func TestGetSpotInstrumentTicker(t *testing.T) {
	NewTestClient().GetSpotInstrumentTicker("BTC-JPY")
}

/*
Public - Filled Orders
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/trades
*/
func TestGetSpotInstrumentTrade(t *testing.T) {
	c := NewTestClient()
	c.GetSpotInstrumentTrade("BTC-JPY", nil)

	options := map[string]string{}
	options["limit"] = "1"
	c.GetSpotInstrumentTrade("BTC-JPY", &options)
}

/*
Public - Market Data
Limit: 20 requests per 2 seconds
GET Request: /api/spot/v3/instruments/<instrument_id>/candles
*/
func TestGetSpotInstrumentCandles(t *testing.T) {
	options := map[string]string{}
	options["start"] = ""
	options["end"] = ""
	options["granularity"] = "300"
	NewTestClient().GetSpotInstrumentCandles("BTC-JPY", &options)
}
