package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/config"
	"github.com/polymarket/go-order-utils/pkg/model"
)

func GetVerifyingContractAddress(chainId *big.Int, contract model.VerifyingContract) (common.Address, error) {
	contracts, err := config.GetContracts(chainId.Int64())
	if err != nil {
		return common.Address{}, err
	}

	switch contract {
	case model.CTFExchange:
		return contracts.Exchange, nil
	case model.NegRiskCTFExchange:
		return contracts.NegRiskExchange, nil
	}

	return common.Address{}, fmt.Errorf("invalid contract")
}

func GetVerifyingContractAddressV2(chainId *big.Int, contract model.VerifyingContract) (common.Address, error) {
	contracts, err := config.GetContracts(chainId.Int64())
	if err != nil {
		return common.Address{}, err
	}

	switch contract {
	case model.CTFExchange:
		return contracts.ExchangeV2, nil
	case model.NegRiskCTFExchange:
		return contracts.NegRiskExchangeV2, nil
	}

	return common.Address{}, fmt.Errorf("invalid contract")
}
