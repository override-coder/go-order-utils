package builder

import (
	"crypto/ecdsa"

	"github.com/polymarket/go-order-utils/pkg/model"
)

//go:generate mockery --name ExchangeOrderBuilderV2
type ExchangeOrderBuilderV2 interface {
	// build an order object including the signature.
	//
	// @param private key
	//
	// @param orderData
	//
	// @returns a SignedOrder object (order + signature)
	BuildSignedOrder(privateKey *ecdsa.PrivateKey, orderData *model.OrderDataV2, contract model.VerifyingContract) (*model.SignedOrderV2, error)

	// Creates an Order object from order data.
	//
	// @param orderData
	//
	// @returns a Order object (not signed)
	BuildOrder(orderData *model.OrderDataV2) (*model.OrderV2, error)

	// Generates the hash of the order from a EIP712TypedData object.
	//
	// @param Order
	//
	// @returns a OrderHash that is a 'common.Hash'
	BuildOrderHash(order *model.OrderV2, contract model.VerifyingContract) (model.OrderHashV2, error)

	// Generates the wrapped hash for POLY_1271 signatures.
	//
	// @param Order
	//
	// @returns a wrapped hash that is a 'common.Hash'
	BuildPoly1271WrappedHash(order *model.OrderV2, contract model.VerifyingContract) (model.OrderHashV2, error)

	// Assembles the final POLY_1271 signature from the inner signature.
	//
	// @param Order
	//
	// @param inner signature
	//
	// @returns a OrderSignature that is []byte
	BuildPoly1271FinalSignature(order *model.OrderV2, contract model.VerifyingContract, innerSig []byte) (model.OrderSignatureV2, error)

	// signs an order
	//
	// @param private key
	//
	// @param order hash
	//
	// @returns a OrderSignature that is []byte
	BuildOrderSignature(privateKey *ecdsa.PrivateKey, orderHash model.OrderHashV2) (model.OrderSignatureV2, error)
}
