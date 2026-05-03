package builder

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/polymarket/go-order-utils/pkg/eip712"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/signer"
	"github.com/polymarket/go-order-utils/pkg/utils"
)

type ExchangeOrderBuilderImplV2 struct {
	chainId       *big.Int
	saltGenerator func() int64
}

var _ ExchangeOrderBuilderV2 = (*ExchangeOrderBuilderImplV2)(nil)

func NewExchangeOrderBuilderImplV2(chainId *big.Int, saltGenerator func() int64) *ExchangeOrderBuilderImplV2 {
	if saltGenerator == nil {
		saltGenerator = utils.GenerateRandomSalt
	}
	return &ExchangeOrderBuilderImplV2{
		chainId:       chainId,
		saltGenerator: saltGenerator,
	}
}

// build an order object including the signature.
//
// @param private key
//
// @param orderData
//
// @returns a SignedOrder object (order + signature)
func (e *ExchangeOrderBuilderImplV2) BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderDataV2, contract model.VerifyingContract) (*model.SignedOrderV2, error) {
	order, err := e.BuildOrder(orderData)
	if err != nil {
		return nil, err
	}

	orderHash, err := e.BuildOrderHash(order, contract)
	if err != nil {
		return nil, err
	}

	signature, err := e.BuildOrderSignatureByType(privateKey, order, contract, orderHash)
	if err != nil {
		return nil, err
	}

	if order.SignatureType.Uint64() != uint64(model.POLY_1271) {
		ok, err := signer.ValidateSignature(order.Signer, orderHash, signature)
		if err != nil {
			return nil, err
		}
		if !ok {
			return nil, fmt.Errorf("signature error")
		}
	}

	return &model.SignedOrderV2{
		OrderV2:   *order,
		Signature: signature,
	}, nil
}

// Creates an Order object from order data.
//
// @param orderData
//
// @returns a Order object (not signed)
func (e *ExchangeOrderBuilderImplV2) BuildOrder(orderData *model.OrderDataV2) (*model.OrderV2, error) {
	var signer common.Address
	if orderData.Signer == "" {
		signer = common.HexToAddress(orderData.Maker)
	} else {
		signer = common.HexToAddress(orderData.Signer)
	}

	var tokenId *big.Int
	var ok bool
	if tokenId, ok = new(big.Int).SetString(orderData.TokenID, 10); !ok {
		return nil, fmt.Errorf("can't parse TokenId: %s as valid *big.Int", orderData.TokenID)
	}

	var expiration *big.Int
	if orderData.Expiration == "" {
		orderData.Expiration = "0"
	}
	if expiration, ok = new(big.Int).SetString(orderData.Expiration, 10); !ok {
		return nil, fmt.Errorf("can't parse Expiration: %s as valid *big.Int", orderData.Expiration)
	}

	var makerAmount *big.Int
	if makerAmount, ok = new(big.Int).SetString(orderData.MakerAmount, 10); !ok {
		return nil, fmt.Errorf("can't parse MakerAmount: %s as valid *big.Int", orderData.MakerAmount)
	}

	var takerAmount *big.Int
	if takerAmount, ok = new(big.Int).SetString(orderData.TakerAmount, 10); !ok {
		return nil, fmt.Errorf("can't parse TakerAmount: %s as valid *big.Int", orderData.TakerAmount)
	}

	var timestamp *big.Int
	if timestamp, ok = new(big.Int).SetString(orderData.Timestamp, 10); !ok {
		return nil, fmt.Errorf("can't parse Timestamp: %s as valid *big.Int", orderData.Timestamp)
	}

	return &model.OrderV2{
		Salt:          new(big.Int).SetInt64(e.saltGenerator()),
		Maker:         common.HexToAddress(orderData.Maker),
		Signer:        signer,
		TokenID:       tokenId,
		MakerAmount:   makerAmount,
		TakerAmount:   takerAmount,
		Side:          new(big.Int).SetInt64(int64(orderData.Side)),
		SignatureType: new(big.Int).SetInt64(int64(orderData.SignatureType)),
		Timestamp:     timestamp,
		Metadata:      common.HexToHash(orderData.Metadata),
		Builder:       common.HexToHash(orderData.Builder),
		Expiration:    expiration,
	}, nil
}

// Generates the hash of the order from a EIP712TypedData object.
//
// @param Order
//
// @returns a OrderHash that is a 'common.Hash'
func (e *ExchangeOrderBuilderImplV2) BuildOrderHash(order *model.OrderV2, contract model.VerifyingContract) (model.OrderHashV2, error) {
	verifyingContract, err := utils.GetVerifyingContractAddressV2(e.chainId, contract)
	if err != nil {
		return model.OrderHash{}, err
	}

	domainSeparator, err := eip712.BuildEIP712DomainSeparator(_PROTOCOL_NAME, _PROTOCOL_VERSION_V2, e.chainId, verifyingContract)
	if err != nil {
		return model.OrderHash{}, err
	}

	values := []interface{}{
		_ORDER_STRUCTURE_HASH_V2,
		order.Salt,
		order.Maker,
		order.Signer,
		order.TokenID,
		order.MakerAmount,
		order.TakerAmount,
		uint8(order.Side.Uint64()),
		uint8(order.SignatureType.Uint64()),
		order.Timestamp,
		order.Metadata,
		order.Builder,
	}
	orderHash, err := eip712.HashTypedDataV4(domainSeparator, _ORDER_STRUCTURE_V2, values)
	if err != nil {
		return model.OrderHash{}, err
	}

	return orderHash, nil
}

// signs an order
//
// @param private key
//
// @param order hash
//
// @returns a OrderSignature that is []byte
func (e *ExchangeOrderBuilderImplV2) BuildOrderSignatureByType(privateKey *ecdsa.PrivateKey, order *model.OrderV2, contract model.VerifyingContract, orderHash model.OrderHashV2) (model.OrderSignatureV2, error) {
	if order.SignatureType.Uint64() != uint64(model.POLY_1271) {
		return e.BuildOrderSignature(privateKey, orderHash)
	}

	exchangeAddress, err := utils.GetVerifyingContractAddressV2(e.chainId, contract)
	if err != nil {
		return nil, err
	}

	appDomainSep, err := eip712.BuildEIP712DomainSeparator(_PROTOCOL_NAME, _PROTOCOL_VERSION_V2, e.chainId, exchangeAddress)
	if err != nil {
		return nil, err
	}

	orderValues := []interface{}{
		_ORDER_STRUCTURE_HASH_V2,
		order.Salt,
		order.Maker,
		order.Signer,
		order.TokenID,
		order.MakerAmount,
		order.TakerAmount,
		uint8(order.Side.Uint64()),
		uint8(order.SignatureType.Uint64()),
		order.Timestamp,
		order.Metadata,
		order.Builder,
	}
	orderEncoded, err := eip712.Encode(_ORDER_STRUCTURE_V2, orderValues)
	if err != nil {
		return nil, err
	}
	contentsHash := crypto.Keccak256Hash(orderEncoded)

	wrappedValues := []interface{}{
		_TYPED_DATA_SIGN_STRUCTURE_HASH_V2,
		contentsHash,
		_DEPOSIT_WALLET_NAME_HASH_V2,
		_DEPOSIT_WALLET_VERSION_HASH_V2,
		e.chainId,
		order.Signer,
		common.Hash{},
	}
	wrappedHash, err := eip712.HashTypedDataV4(appDomainSep, _TYPED_DATA_SIGN_STRUCTURE_V2, wrappedValues)
	if err != nil {
		return nil, err
	}

	innerSig, err := signer.Sign(privateKey, wrappedHash)
	if err != nil {
		return nil, err
	}

	orderTypeString := []byte("Order(uint256 salt,address maker,address signer,uint256 tokenId,uint256 makerAmount,uint256 takerAmount,uint8 side,uint8 signatureType,uint256 timestamp,bytes32 metadata,bytes32 builder)")
	orderTypeLen := len(orderTypeString)
	finalSig := make([]byte, 0, len(innerSig)+32+32+len(orderTypeString)+2)
	finalSig = append(finalSig, innerSig...)
	finalSig = append(finalSig, appDomainSep.Bytes()...)
	finalSig = append(finalSig, contentsHash.Bytes()...)
	finalSig = append(finalSig, orderTypeString...)
	finalSig = append(finalSig, byte(orderTypeLen>>8), byte(orderTypeLen))
	return finalSig, nil
}

func (e *ExchangeOrderBuilderImplV2) BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHashV2) (model.OrderSignatureV2, error) {
	return signer.Sign(privateKey, orderHash)
}
