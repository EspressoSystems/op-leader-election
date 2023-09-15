package e2eutils

import (
	"crypto/ecdsa"
	"fmt"

	hdwallet "github.com/ethereum-optimism/go-ethereum-hdwallet"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// DefaultMnemonicConfig is the default mnemonic used in testing.
// We prefer a mnemonic rather than direct private keys to make it easier
// to export all testing keys in external tooling for use during debugging.
// If these values are changed, it is subject to breaking tests. They
// must be in sync with the values in the DeployConfig used to create the system.
var DefaultMnemonicConfig = &MnemonicConfig{
	Mnemonic:     "test test test test test test test test test test test junk",
	CliqueSigner: "m/44'/60'/0'/0/0",
	Proposer:     "m/44'/60'/0'/0/1",
	Batcher:      "m/44'/60'/0'/0/2",
	Deployer:     "m/44'/60'/0'/0/3",
	Alice:        "m/44'/60'/0'/0/4",
	SequencerP2P: "m/44'/60'/0'/0/5",
	Bob:          "m/44'/60'/0'/0/7",
	Mallory:      "m/44'/60'/0'/0/8",
	SysCfgOwner:  "m/44'/60'/0'/0/9",
	Batcher1:     "m/44'/60'/0'/0/10",
	Batcher2:     "m/44'/60'/0'/0/11",
	Batcher3:     "m/44'/60'/0'/0/12",
	Batcher4:     "m/44'/60'/0'/0/13",
}

// MnemonicConfig configures the private keys for the hive testnet.
// It's json-serializable, so we can ship it to e.g. the hardhat script client.
type MnemonicConfig struct {
	Mnemonic string

	CliqueSigner string
	Deployer     string
	SysCfgOwner  string

	// rollup actors
	Proposer     string
	Batcher      string
	SequencerP2P string

	// prefunded L1/L2 accounts for testing
	Alice   string
	Bob     string
	Mallory string

	Batcher1 string
	Batcher2 string
	Batcher3 string
	Batcher4 string
}

// Secrets computes the private keys for all mnemonic paths,
// which can then be kept around for fast precomputed private key access.
func (m *MnemonicConfig) Secrets() (*Secrets, error) {
	wallet, err := hdwallet.NewFromMnemonic(m.Mnemonic)
	if err != nil {
		return nil, fmt.Errorf("failed to create wallet: %w", err)
	}
	account := func(path string) accounts.Account {
		return accounts.Account{URL: accounts.URL{Path: path}}
	}

	deployer, err := wallet.PrivateKey(account(m.Deployer))
	if err != nil {
		return nil, err
	}
	cliqueSigner, err := wallet.PrivateKey(account(m.CliqueSigner))
	if err != nil {
		return nil, err
	}
	sysCfgOwner, err := wallet.PrivateKey(account(m.SysCfgOwner))
	if err != nil {
		return nil, err
	}
	proposer, err := wallet.PrivateKey(account(m.Proposer))
	if err != nil {
		return nil, err
	}
	batcher, err := wallet.PrivateKey(account(m.Batcher))
	if err != nil {
		return nil, err
	}
	sequencerP2P, err := wallet.PrivateKey(account(m.SequencerP2P))
	if err != nil {
		return nil, err
	}
	alice, err := wallet.PrivateKey(account(m.Alice))
	if err != nil {
		return nil, err
	}
	bob, err := wallet.PrivateKey(account(m.Bob))
	if err != nil {
		return nil, err
	}
	mallory, err := wallet.PrivateKey(account(m.Mallory))
	if err != nil {
		return nil, err
	}
	batcher1, err := wallet.PrivateKey(account(m.Batcher1))
	if err != nil {
		return nil, err
	}
	batcher2, err := wallet.PrivateKey(account(m.Batcher2))
	if err != nil {
		return nil, err
	}
	batcher3, err := wallet.PrivateKey(account(m.Batcher3))
	if err != nil {
		return nil, err
	}
	batcher4, err := wallet.PrivateKey(account(m.Batcher4))
	if err != nil {
		return nil, err
	}

	return &Secrets{
		Deployer:     deployer,
		SysCfgOwner:  sysCfgOwner,
		CliqueSigner: cliqueSigner,
		Proposer:     proposer,
		Batcher:      batcher,
		SequencerP2P: sequencerP2P,
		Alice:        alice,
		Bob:          bob,
		Mallory:      mallory,
		Wallet:       wallet,
		Batcher1:     batcher1,
		Batcher2:     batcher2,
		Batcher3:     batcher3,
		Batcher4:     batcher4,
	}, nil
}

// Secrets bundles secp256k1 private keys for all common rollup actors for testing purposes.
type Secrets struct {
	Deployer     *ecdsa.PrivateKey
	CliqueSigner *ecdsa.PrivateKey
	SysCfgOwner  *ecdsa.PrivateKey

	// rollup actors
	Proposer     *ecdsa.PrivateKey
	Batcher      *ecdsa.PrivateKey
	SequencerP2P *ecdsa.PrivateKey

	// prefunded L1/L2 accounts for testing
	Alice   *ecdsa.PrivateKey
	Bob     *ecdsa.PrivateKey
	Mallory *ecdsa.PrivateKey

	// Share the wallet to be able to generate more accounts
	Wallet *hdwallet.Wallet

	// Batchers for testing
	Batcher1 *ecdsa.PrivateKey
	Batcher2 *ecdsa.PrivateKey
	Batcher3 *ecdsa.PrivateKey
	Batcher4 *ecdsa.PrivateKey
}

// EncodePrivKey encodes the given private key in 32 bytes
func EncodePrivKey(priv *ecdsa.PrivateKey) hexutil.Bytes {
	privkey := make([]byte, 32)
	blob := priv.D.Bytes()
	copy(privkey[32-len(blob):], blob)
	return privkey
}

func EncodePrivKeyToString(priv *ecdsa.PrivateKey) string {
	return hexutil.Encode(EncodePrivKey(priv))
}

// Addresses computes the ethereum address of each account,
// which can then be kept around for fast precomputed address access.
func (s *Secrets) Addresses() *Addresses {
	return &Addresses{
		Deployer:     crypto.PubkeyToAddress(s.Deployer.PublicKey),
		CliqueSigner: crypto.PubkeyToAddress(s.CliqueSigner.PublicKey),
		SysCfgOwner:  crypto.PubkeyToAddress(s.SysCfgOwner.PublicKey),
		Proposer:     crypto.PubkeyToAddress(s.Proposer.PublicKey),
		Batcher:      crypto.PubkeyToAddress(s.Batcher.PublicKey),
		SequencerP2P: crypto.PubkeyToAddress(s.SequencerP2P.PublicKey),
		Alice:        crypto.PubkeyToAddress(s.Alice.PublicKey),
		Bob:          crypto.PubkeyToAddress(s.Bob.PublicKey),
		Mallory:      crypto.PubkeyToAddress(s.Mallory.PublicKey),
		Batcher1:     crypto.PubkeyToAddress(s.Batcher1.PublicKey),
		Batcher2:     crypto.PubkeyToAddress(s.Batcher2.PublicKey),
		Batcher3:     crypto.PubkeyToAddress(s.Batcher3.PublicKey),
		Batcher4:     crypto.PubkeyToAddress(s.Batcher3.PublicKey),
	}
}

// Addresses bundles the addresses for all common rollup addresses for testing purposes.
type Addresses struct {
	Deployer     common.Address
	CliqueSigner common.Address
	SysCfgOwner  common.Address

	// rollup actors
	Proposer     common.Address
	Batcher      common.Address
	SequencerP2P common.Address

	// prefunded L1/L2 accounts for testing
	Alice   common.Address
	Bob     common.Address
	Mallory common.Address

	// Batchers for testing
	Batcher1 common.Address
	Batcher2 common.Address
	Batcher3 common.Address
	Batcher4 common.Address
}

func (a *Addresses) All() []common.Address {
	return []common.Address{
		a.Deployer,
		a.CliqueSigner,
		a.SysCfgOwner,
		a.Proposer,
		a.Batcher,
		a.SequencerP2P,
		a.Alice,
		a.Bob,
		a.Mallory,
		a.Batcher1,
		a.Batcher2,
		a.Batcher3,
		a.Batcher4,
	}
}
