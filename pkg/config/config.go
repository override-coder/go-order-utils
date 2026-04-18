package config

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type Contracts struct {
	Exchange         common.Address
	FeeModule        common.Address
	NegRiskExchange  common.Address
	NegRiskFeeModule common.Address
	NegRiskAdapter   common.Address
	Collateral       common.Address
	Conditional      common.Address

	ExchangeV2        common.Address
	NegRiskExchangeV2 common.Address
}

var (
	_AMOY_CONTRACTS = &Contracts{
		Exchange:          common.HexToAddress("0xdFE02Eb6733538f8Ea35D585af8DE5958AD99E40"),
		FeeModule:         common.HexToAddress("0xE3f18aCc55091e2c48d883fc8C8413319d4Ab7b0"),
		NegRiskExchange:   common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
		NegRiskFeeModule:  common.HexToAddress("0xB768891e3130F6dF18214Ac804d4DB76c2C37730"),
		NegRiskAdapter:    common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
		Collateral:        common.HexToAddress("0x9c4e1703476e875070ee25b56a58b008cfb8fa78"),
		Conditional:       common.HexToAddress("0x69308FB512518e39F9b16112fA8d994F4e2Bf8bB"),
		ExchangeV2:        common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B"),
		NegRiskExchangeV2: common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59"),
	}

	_MATIC_CONTRACTS = &Contracts{
		Exchange:          common.HexToAddress("0x4bFb41d5B3570DeFd03C39a9A4D8dE6Bd8B8982E"),
		FeeModule:         common.HexToAddress("0xE3f18aCc55091e2c48d883fc8C8413319d4Ab7b0"),
		NegRiskExchange:   common.HexToAddress("0xC5d563A36AE78145C45a50134d48A1215220f80a"),
		NegRiskFeeModule:  common.HexToAddress("0xB768891e3130F6dF18214Ac804d4DB76c2C37730"),
		NegRiskAdapter:    common.HexToAddress("0xd91E80cF2E7be2e162c6513ceD06f1dD0dA35296"),
		Collateral:        common.HexToAddress("0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"),
		Conditional:       common.HexToAddress("0x4D97DCd97eC945f40cF65F87097ACe5EA0476045"),
		ExchangeV2:        common.HexToAddress("0xE111180000d2663C0091e4f400237545B87B996B"),
		NegRiskExchangeV2: common.HexToAddress("0xe2222d279d744050d28e00520010520000310F59"),
	}
)

func GetContracts(chainId int64) (*Contracts, error) {
	switch chainId {
	case 137:
		return _MATIC_CONTRACTS, nil
	case 80002:
		return _AMOY_CONTRACTS, nil
	default:
		return nil, fmt.Errorf("invalid chain id")
	}
}
