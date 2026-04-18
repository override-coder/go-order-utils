package model

type SignatureType = int

const (
	// ECDSA EIP712 signatures signed by EOAs
	EOA SignatureType = iota

	// EIP712 signatures signed by EOAs that own Polymarket Proxy wallets
	POLY_PROXY

	// EIP712 signatures signed by EOAs that own Polymarket Gnosis safes
	POLY_GNOSIS_SAFE

	/**
	 * EIP1271 signatures signed by smart contracts. To be used by smart contract wallets or vaults
	 */
	POLY_1271
)
