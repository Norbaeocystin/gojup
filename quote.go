package gojup

import (
	"encoding/json"
	"github.com/gagliardetto/solana-go"
	"math/big"
	"net/url"
)

var JUP_API = "https://quote-api.jup.ag/v4/"

// default slippage 1 %
func GetQuote(inputMint, outputMint, user solana.PublicKey, amount *big.Int) (QuoteGenerated, error) {
	urlParsed, _ := url.Parse(JUP_API + "quote")
	values := urlParsed.Query()
	values.Add("inputMint", inputMint.String())
	values.Add("outputMint", outputMint.String())
	values.Add("amount", amount.String())
	values.Add("swapMode", "ExactIn") // ExactIn or ExactOut
	values.Add("slippageBps", "1")
	values.Add("onlyDirectRoutes", "true")
	values.Add("asLegacyTransaction", "true")
	values.Add("userPublicKey", user.String())
	urlParsed.RawQuery = values.Encode()
	b, err := Get(urlParsed.String())
	var quote QuoteGenerated
	json.Unmarshal(b, &quote)
	return quote, err
}

type QuoteParams struct {
	InputMint           solana.PublicKey
	OutputMint          solana.PublicKey
	Amount              *big.Int
	SwapMode            string // ExactIn or ExactOut
	SlippageBps         string
	OnlyDirectRoutes    bool
	AsLegacyTransaction bool
	User                solana.PublicKey
}

func GetQuoteWithParams(params QuoteParams) (QuoteGenerated, error) {
	urlParsed, _ := url.Parse(JUP_API + "quote")
	values := urlParsed.Query()
	values.Add("inputMint", params.InputMint.String())
	values.Add("outputMint", params.OutputMint.String())
	values.Add("amount", params.Amount.String())
	values.Add("swapMode", params.SwapMode) // ExactIn or ExactOut
	values.Add("slippageBps", params.SlippageBps)
	if params.OnlyDirectRoutes {
		values.Add("onlyDirectRoutes", "true")
	} else {
		values.Add("onlyDirectRoutes", "false")
	}
	if params.AsLegacyTransaction {
		values.Add("asLegacyTransaction", "true")
	} else {
		values.Add("asLegacyTransaction", "false")
	}
	values.Add("userPublicKey", params.User.String())
	urlParsed.RawQuery = values.Encode()
	b, err := Get(urlParsed.String())
	var quote QuoteGenerated
	json.Unmarshal(b, &quote)
	return quote, err
}
