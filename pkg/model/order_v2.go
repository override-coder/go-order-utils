package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type OrderSignatureV2 = []byte

type OrderHashV2 = common.Hash

type OrderDataV2 struct {
	// Maker of the order, i.e. the source of funds for the order.
	Maker string `json:"maker"`

	// TokenID of the CTF ERC1155 asset to be bought or sold.
	// If BUY, this is the tokenID of the asset to be bought, i.e. the makerAssetID.
	// If SELL, this is the tokenID of the asset to be sold, i.e. the takerAssetID.
	TokenID string `json:"tokenId"`

	// MakerAmount is the max amount of tokens to be sold.
	MakerAmount string `json:"makerAmount"`

	// TakerAmount is the minimum amount of tokens to be received.
	TakerAmount string `json:"takerAmount"`

	// Side of the order, BUY or SELL.
	Side Side `json:"side"`

	// Signer of the order. Optional; if empty, the signer is the maker of the order.
	Signer string `json:"signer,omitempty"`

	// SignatureType used by the order. Default value is EOA.
	SignatureType SignatureType `json:"signatureType,omitempty"`

	// Timestamp of the order.
	Timestamp string `json:"timestamp,omitempty"`

	// Metadata of the order.
	Metadata string `json:"metadata,omitempty"`

	// Builder of the order.
	Builder string `json:"builder,omitempty"`

	// Expiration timestamp of the order, unix seconds. "0" means no expiration.
	Expiration string `json:"expiration,omitempty"`
}

type OrderV2 struct {
	// Unique salt to ensure entropy.
	Salt *big.Int `json:"salt"`

	// Maker of the order, i.e. the source of funds for the order.
	Maker common.Address `json:"maker"`

	// Signer of the order.
	Signer common.Address `json:"signer"`

	// TokenID of the CTF ERC1155 asset to be bought or sold.
	// If BUY, this is the tokenID of the asset to be bought, i.e. the makerAssetID.
	// If SELL, this is the tokenID of the asset to be sold, i.e. the takerAssetID.
	TokenID *big.Int `json:"tokenId"`

	// MakerAmount is the max amount of tokens to be sold.
	MakerAmount *big.Int `json:"makerAmount"`

	// TakerAmount is the minimum amount of tokens to be received.
	TakerAmount *big.Int `json:"takerAmount"`

	// Side of the order, BUY or SELL.
	Side *big.Int `json:"side"`

	// SignatureType used by the order.
	SignatureType *big.Int `json:"signatureType"`

	// Timestamp of the order.
	Timestamp *big.Int `json:"timestamp"`

	// Metadata of the order.
	Metadata common.Hash `json:"metadata"`

	// Builder of the order.
	Builder common.Hash `json:"builder"`
}

type SignedOrderV2 struct {
	OrderV2

	// The order signature
	Signature OrderSignatureV2
}
