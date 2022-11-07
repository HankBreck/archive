package keeper_test

import (
	"archive/x/contractregistry/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO: Test msgServer like this https://github.com/osmosis-labs/osmosis/blob/main/x/tokenfactory/keeper/msg_server_test.go

// TestRegisterContractMsg tests TypeMsgRegisterContract message is emitted on a successful registration
func (suite *KeeperTestSuite) TestRegisterContractMsg() {
	defaultMsg := &types.MsgRegisterContract{
		Creator:     suite.TestAccs[0].String(),
		Description: "test description",
		Authors:     []string{"hank", "david"},
		ContactInfo: &types.ContactInfo{
			Method: types.ContactMethod_Email,
			Value:  "hank@archive.com",
		},
		MoreInfoUri:       "google.com/more-info",
		SigningDataSchema: []byte("{test: 1}"),
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}
	tests := map[string]struct {
		inputMsg      *types.MsgRegisterContract
		expEventCount int
	}{
		"simple_register": {
			inputMsg:      defaultMsg,
			expEventCount: 1,
		},
		"nil_msg": {
			inputMsg:      nil,
			expEventCount: 0,
		},
	}
	for name, test := range tests {
		suite.Run(name, func() {
			ctx := suite.Ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))
			// Test RegisterContract
			suite.msgServer.RegisterContract(sdk.WrapSDKContext(ctx), test.inputMsg)
			// Ensure events are emitted
			suite.AssertEventEmitted(ctx, types.TypeMsgRegisterContract, test.expEventCount)
		})
	}
}

func (suite *KeeperTestSuite) TestRegisterContract() {
	k := suite.App.ContractregistryKeeper

	// Register a contract via msgServer
	defaultMsg := &types.MsgRegisterContract{
		Creator:     suite.TestAccs[0].String(),
		Description: "test description",
		Authors:     []string{"hank", "david"},
		ContactInfo: &types.ContactInfo{
			Method: types.ContactMethod_Email,
			Value:  "hank@archive.com",
		},
		MoreInfoUri:       "google.com/more-info",
		SigningDataSchema: []byte("{test: 1}"),
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}
	res, err := suite.msgServer.RegisterContract(sdk.WrapSDKContext(suite.Ctx), defaultMsg)
	suite.NoError(err)
	suite.Equal(uint64(0), res.Id)

	// Ensure actual and expected contracts match
	actualContract, _ := k.GetContract(suite.Ctx, res.Id)
	expectedContract := types.Contract{
		Id:          uint64(0),
		Description: "test description",
		Authors:     []string{"hank", "david"},
		ContactInfo: &types.ContactInfo{
			Method: types.ContactMethod_Email,
			Value:  "hank@archive.com",
		},
		MoreInfoUri:       "google.com/more-info",
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}
	suite.Equal(expectedContract, *actualContract)

	// Ensure actual and expected signing data match
	actualSigningData, _ := k.GetSigningData(suite.Ctx, res.Id)
	expectedSigningData := []byte("{test: 1}")
	suite.Equal(expectedSigningData, actualSigningData.Bytes())
}
