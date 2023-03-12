package keeper_test

import (
	"context"
	_ "embed"
	"testing"

	keepertest "github.com/HankBreck/archive/testutil/keeper"
	"github.com/HankBreck/archive/x/cda/keeper"
	"github.com/HankBreck/archive/x/cda/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CdaKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

// TODO:
// Write tests for msg server functions
// Create CDA
//		Test signing metadata valid json but not schema match
// Approve CDA
//
// 		Assert fails with error on invalid signing data
// 		Assert fails with error on non-owner Creator
// 		Assert fails with error on invalid cda.Status
//		Assert fails with error on a CdaId that does not exist

// go:embed x/cda/testdata/escrow_0.7.wasm
var hackatomWasm []byte

func (suite *KeeperTestSuite) TestCreateCda_Trash() {
	msgServer := suite.msgServer
	goCtx := sdk.WrapSDKContext(suite.Ctx)

	// Prepare contract for upload
	codeId, _ := suite.prepareWasm(suite.Ctx, suite.TestAccs[0])

	// Register contract with codeID
	contract, _ := suite.PrepareContractWithSchema(codeId, suite.getTestSchema())

	// Prepare signer id
	signerId, _ := suite.PrepareCertificate(suite.TestAccs[0], &suite.TestAccs[1])

	// Test CDA creation
	createRes, err := msgServer.CreateCda(goCtx, &types.MsgCreateCda{
		Creator:        suite.TestAccs[0].String(),
		SignerIds:      []uint64{signerId},
		ContractId:     contract.Id,
		SigningData:    suite.getTestDoc(),
		WitnessInitMsg: []byte{0x0}, // TODO: Add wasm init msg
	})
	suite.NoError(err)
	suite.NotNil(createRes)
}

func (suite *KeeperTestSuite) prepareWasm(ctx sdk.Context, creator sdk.AccAddress) (uint64, error) {
	codeId, _, err := suite.wasmKeeper.Create(ctx, creator, hackatomWasm, &wasmtypes.DefaultUploadAccess)
	return codeId, err
}
