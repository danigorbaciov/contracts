package kyber_test

import (
	"context"
	"os"
    "fmt"
	"testing"
    "math/big"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
    "github.com/tokencard/contracts/v2/pkg/bindings/mocks"
	"github.com/tokencard/contracts/v2/pkg/bindings/mocks/kyber"
    "github.com/tokencard/contracts/v2/pkg/bindings"
    . "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
    . "github.com/tokencard/contracts/v2/test/shared"
)


var Wallet *bindings.Wallet
var WalletAddress common.Address

var KyberNetworkProxy *kyber.KyberNetworkProxy
var KyberNetworkProxyAddress common.Address

var KyberNetwork *kyber.KyberNetwork
var KyberNetworkAddress common.Address

var FeeBurner *kyber.FeeBurner
var FeeBurnerAddress common.Address

var ExpectedRate *kyber.ExpectedRate
var ExpectedRateAddress common.Address

var KNC *mocks.BurnerToken
var KNCAddress common.Address

func TestTokenWhitelistSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "kyber Suite")
}

var _ = BeforeEach(func() {
	err := InitializeBackend()
	Expect(err).ToNot(HaveOccurred())

    var tx *types.Transaction

    KyberNetworkProxyAddress, tx, KyberNetworkProxy, err = kyber.DeployKyberNetworkProxy(BankAccount.TransactOpts(), Backend, Owner.Address())
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    KyberNetworkAddress, tx, KyberNetwork, err = kyber.DeployKyberNetwork(BankAccount.TransactOpts(), Backend, Owner.Address())
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())


    KNCAddress, tx, KNC, err = mocks.DeployBurnerToken(Owner.TransactOpts(), Backend)
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    kncEthPrec := new(big.Int)
    kncEthPrec.SetString("1273294871580578838478",10)
    FeeBurnerAddress, tx, FeeBurner, err = kyber.DeployFeeBurner(BankAccount.TransactOpts(), Backend, Owner.Address(), KNCAddress, KyberNetworkAddress, kncEthPrec)
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    ExpectedRateAddress, tx, ExpectedRate, err = kyber.DeployExpectedRate(BankAccount.TransactOpts(), Backend, KyberNetworkAddress, KNCAddress, Owner.Address())
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    WalletAddress, tx, Wallet, err = bindings.DeployWallet(BankAccount.TransactOpts(), Backend, Owner.Address(), true, ENSRegistryAddress, TokenWhitelistName, ControllerName, LicenceName, EthToWei(1))
	Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    //set network addresses and params
    tx, err = KyberNetworkProxy.SetKyberNetworkContract(Owner.TransactOpts(), KyberNetworkAddress)
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    tx, err = KyberNetwork.SetKyberProxy(Owner.TransactOpts(), KyberNetworkProxyAddress)
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    tx, err = KyberNetwork.SetFeeBurner(Owner.TransactOpts(), FeeBurnerAddress)
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    tx, err = KyberNetwork.SetExpectedRate(Owner.TransactOpts(), ExpectedRateAddress)
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    tx, err = KyberNetwork.SetParams(Owner.TransactOpts(), big.NewInt(100000000000), big.NewInt(20))
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())

    tx, err = KyberNetwork.SetEnable(Owner.TransactOpts(), true)
    Expect(err).ToNot(HaveOccurred())
	Backend.Commit()
	Expect(isSuccessful(tx)).To(BeTrue())
})

var _ = AfterEach(func() {
    td := CurrentGinkgoTestDescription()

	if td.Failed {
		fmt.Fprintf(GinkgoWriter, "\nLast Executed Smart Contract Line for %s:%d\n", td.FileName, td.LineNumber)
		fmt.Fprintln(GinkgoWriter, TestRig.LastExecuted())
	}
	err := Backend.Close()
	Expect(err).ToNot(HaveOccurred())

})

var _ = AfterSuite(func() {
	// TestRig.ExpectMinimumCoverage("mocks/kyber/KyberNetworkProxy.sol", 0.0)
	TestRig.PrintGasUsage(os.Stdout)
})

func isSuccessful(tx *types.Transaction) bool {
	r, err := Backend.TransactionReceipt(context.Background(), tx.Hash())
	Expect(err).ToNot(HaveOccurred())
	return r.Status == types.ReceiptStatusSuccessful
}