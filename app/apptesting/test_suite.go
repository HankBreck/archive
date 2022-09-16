package apptesting

import (
	"time"

	"archive/app"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
	tmtypes "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/stretchr/testify/suite"
)

type KeeperTestHelper struct {
	suite.Suite

	App         *app.App
	Ctx         sdk.Context
	QueryHelper *baseapp.QueryServiceTestHelper
	TestAccs    []sdk.AccAddress
}

func (s *KeeperTestHelper) Setup() {
	s.App = app.Setup(false)
	s.Ctx = s.App.NewContext(false, tmtypes.Header{Height: 1, ChainID: "casper-1", Time: time.Now().UTC()})
	s.QueryHelper = &baseapp.QueryServiceTestHelper{
		GRPCQueryRouter: s.App.GRPCQueryRouter(),
		Ctx:             s.Ctx,
	}

	s.TestAccs = CreateRandomAccounts(10)
}

func CreateRandomAccounts(num int) []sdk.AccAddress {
	testAddrs := make([]sdk.AccAddress, num)
	for i := 0; i < num; i++ {
		pubKey := ed25519.GenPrivKey().PubKey()
		pubKey.Address()
		testAddrs[i] = sdk.AccAddress(pubKey.Address())
	}
	return testAddrs
}
