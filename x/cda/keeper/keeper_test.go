package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/HankBreck/archive/x/cda/keeper"
	"github.com/HankBreck/archive/x/cda/types"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	"github.com/HankBreck/archive/app/apptesting"
)

type KeeperTestSuite struct {
	apptesting.KeeperTestHelper

	queryClient types.QueryClient
	msgServer   types.MsgServer

	wasmKeeper *wasmkeeper.PermissionedKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.Setup()
	suite.queryClient = types.NewQueryClient(suite.QueryHelper)
	suite.msgServer = keeper.NewMsgServerImpl(suite.App.CdaKeeper)
	suite.wasmKeeper = wasmkeeper.NewDefaultPermissionKeeper(suite.App.WasmKeeper)
}

func (suite *KeeperTestSuite) PrepareContracts(count int) []uint64 {
	result := []uint64{}
	for i := 0; i < count; i++ {
		defaultContract := types.Contract{
			Description:       "dummy contract",
			Authors:           []string{"hank", "david"},
			ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
			MoreInfoUri:       "google.com",
			TemplateUri:       "google.com/template",
			TemplateSchemaUri: "google.com/template-schema",
			WitnessCodeId:     1,
		}
		id := suite.App.CdaKeeper.AppendContract(suite.Ctx, defaultContract)
		result = append(result, id)
	}
	return result
}
