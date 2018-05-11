package finance

import "context"

type (
	// QuoteType alias for asset classification.
	QuoteType string
	// MarketState alias for market state.
	MarketState string
)

const (
	// QuoteTypeEquity the returned quote should be an equity.
	QuoteTypeEquity QuoteType = "EQUITY"
	// QuoteTypeIndex the returned quote should be an index.
	QuoteTypeIndex QuoteType = "INDEX"
	// QuoteTypeOption the returned quote should be an option contract.
	QuoteTypeOption QuoteType = "OPTION"
	// QuoteTypeForexPair the returned quote should be a forex pair.
	QuoteTypeForexPair QuoteType = "CURRENCY"
	// QuoteTypeFuture the returned quote should be a futures contract.
	QuoteTypeFuture QuoteType = "FUTURE"
	// QuoteTypeETF the returned quote should be an etf.
	QuoteTypeETF QuoteType = "ETF"
	// QuoteTypeMutualFund the returned quote should be an mutual fund.
	QuoteTypeMutualFund QuoteType = "MUTUALFUND"

	// MarketStatePrePre pre-pre market state.
	MarketStatePrePre MarketState = "PREPRE"
	// MarketStatePre pre market state.
	MarketStatePre MarketState = "PRE"
	// MarketStateRegular regular market state.
	MarketStateRegular MarketState = "REGULAR"
	// MarketStatePost post market state.
	MarketStatePost MarketState = "POST"
	// MarketStatePostPost post-post market state.
	MarketStatePostPost MarketState = "POSTPOST"
	// MarketStateClosed closed market state.
	MarketStateClosed MarketState = "CLOSED"
)

// QuoteParams determines which quote to
// retrieve and provides a context.
type QuoteParams struct {
	Symbol string `form:"symbols"`
	// Context used for request. It may carry deadlines, cancelation signals,
	// and other request-scoped values across API boundaries and between
	// processes.
	//
	// Note that a cancelled or timed out context does not provide any
	// guarantee whether the operation was or was not completed.
	Context *context.Context `form:"-"`
}

// Quote is the basic quote structure.
type Quote struct {
	// Quote classifying fields.
	Symbol      string      `json:"symbol"`
	MarketState MarketState `json:"marketState"`
	QuoteType   QuoteType   `json:"quoteType"`
	ShortName   string      `json:"shortName"`
	LongName    string      `json:"longName"`

	// Regular session quote data.
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
	RegularMarketPreviousClose float64 `json:"regularMarketPreviousClose"`
	RegularMarketPrice         float64 `json:"regularMarketPrice"`
	RegularMarketTime          int     `json:"regularMarketTime"`
	RegularMarketChange        float64 `json:"regularMarketChange"`
	RegularMarketOpen          float64 `json:"regularMarketOpen"`
	RegularMarketDayHigh       float64 `json:"regularMarketDayHigh"`
	RegularMarketDayLow        float64 `json:"regularMarketDayLow"`
	RegularMarketVolume        int     `json:"regularMarketVolume"`

	// Quote depth.
	Bid     float64 `json:"bid"`
	Ask     float64 `json:"ask"`
	BidSize int     `json:"bidSize"`
	AskSize int     `json:"askSize"`

	// Pre-market quote data.
	PreMarketPrice         float64 `json:"preMarketPrice"`
	PreMarketChange        float64 `json:"preMarketChange"`
	PreMarketChangePercent float64 `json:"preMarketChangePercent"`
	PreMarketTime          int     `json:"preMarketTime"`

	// Post-market quote data.
	PostMarketPrice         float64 `json:"postMarketPrice"`
	PostMarketChange        float64 `json:"postMarketChange"`
	PostMarketChangePercent float64 `json:"postMarketChangePercent"`
	PostMarketTime          int     `json:"postMarketTime"`

	// 52wk ranges.
	FiftyTwoWeekLowChange         float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent  float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekHighChange        float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow               float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh              float64 `json:"fiftyTwoWeekHigh"`

	// Averages.
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`

	// Volume metrics.
	AverageDailyVolume3Month int `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day  int `json:"averageDailyVolume10Day"`

	// Quote meta-data.
	QuoteSource               string `json:"quoteSourceName"`
	CurrencyID                string `json:"currency"`
	IsTradeable               bool   `json:"tradeable"`
	QuoteDelay                int    `json:"exchangeDataDelayedBy"`
	FullExchangeName          string `json:"fullExchangeName"`
	SourceInterval            int    `json:"sourceInterval"`
	ExchangeTimezoneName      string `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName string `json:"exchangeTimezoneShortName"`
	GMTOffSetMilliseconds     int    `json:"gmtOffSetMilliseconds"`
	MarketID                  string `json:"market"`
	ExchangeID                string `json:"exchange"`

	// Equity-only.
	EpsTrailingTwelveMonths     float64 `json:"epsTrailingTwelveMonths"`
	EpsForward                  float64 `json:"epsForward"`
	EarningsTimestamp           int     `json:"earningsTimestamp"`
	EarningsTimestampStart      int     `json:"earningsTimestampStart"`
	EarningsTimestampEnd        int     `json:"earningsTimestampEnd"`
	TrailingAnnualDividendRate  float64 `json:"trailingAnnualDividendRate"`
	DividendDate                int     `json:"dividendDate"`
	TrailingAnnualDividendYield float64 `json:"trailingAnnualDividendYield"`
	TrailingPE                  float64 `json:"trailingPE"`
	ForwardPE                   float64 `json:"forwardPE"`
	BookValue                   float64 `json:"bookValue"`
	PriceToBook                 float64 `json:"priceToBook"`
	SharesOutstanding           int     `json:"sharesOutstanding"`
	MarketCap                   int64   `json:"marketCap"`

	// MutualFund/ETF-only.
	YTDReturn                    float64 `json:"ytdReturn"`
	TrailingThreeMonthReturns    float64 `json:"trailingThreeMonthReturns"`
	TrailingThreeMonthNavReturns float64 `json:"trailingThreeMonthNavReturns"`

	// Options/Futures-only.
	UnderlyingSymbol         string  `json:"underlyingSymbol"`
	OpenInterest             int     `json:"openInterest"`
	ExpireDate               int     `json:"expireDate"`
	Strike                   float64 `json:"strike"`
	UnderlyingExchangeSymbol string  `json:"underlyingExchangeSymbol"`
	HeadSymbolAsString       string  `json:"headSymbolAsString"`
	IsContractSymbol         bool    `json:"contractSymbol"`
}