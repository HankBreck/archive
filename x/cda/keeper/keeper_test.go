package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"archive/app/apptesting"
	"archive/x/cda/types"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper
	queryClient types.QueryClient
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)

	// Osmosis configures module params here
}
