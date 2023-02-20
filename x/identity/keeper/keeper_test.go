package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/HankBreck/archive/app/apptesting"
	"github.com/HankBreck/archive/x/identity/keeper"
	"github.com/HankBreck/archive/x/identity/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
	msgServer   types.MsgServer
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.IdentityKeeper)
}
