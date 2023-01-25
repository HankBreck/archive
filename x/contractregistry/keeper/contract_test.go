package keeper_test

import (
	"github.com/HankBreck/archive/x/contractregistry/types"
)

func (suite *KeeperTestSuite) TestAppendContract() {
	k := suite.App.ContractregistryKeeper
	defaultContract := &types.Contract{
		Description:       "dummy contract",
		Authors:           []string{"hank", "david"},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
		MoreInfoUri:       "google.com",
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}

	tests := map[string]struct {
		inputContract *types.Contract
		expPanic      bool
		expContract   *types.Contract
	}{
		"simple_append": {
			inputContract: defaultContract,
			expPanic:      false,
			expContract: &types.Contract{
				Id:                uint64(0),
				Description:       "dummy contract",
				Authors:           []string{"hank", "david"},
				ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
				MoreInfoUri:       "google.com",
				TemplateUri:       "google.com/template",
				TemplateSchemaUri: "google.com/template-schema",
			},
		},
		"overwrite_preset_id": {
			inputContract: &types.Contract{
				Id:                uint64(10), // this should be overwritten
				Description:       "dummy contract",
				Authors:           []string{"hank", "david"},
				ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
				MoreInfoUri:       "google.com",
				TemplateUri:       "google.com/template",
				TemplateSchemaUri: "google.com/template-schema",
			},
			expPanic: false,
			expContract: &types.Contract{
				Id:                uint64(1),
				Description:       "dummy contract",
				Authors:           []string{"hank", "david"},
				ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
				MoreInfoUri:       "google.com",
				TemplateUri:       "google.com/template",
				TemplateSchemaUri: "google.com/template-schema",
			},
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			if !test.expPanic {
				id := k.AppendContract(suite.Ctx, *test.inputContract)
				actual, _ := k.GetContract(suite.Ctx, id)
				suite.Equal(*test.expContract, *actual)
			} else {
				suite.Panics(func() {
					k.AppendContract(suite.Ctx, *test.inputContract)
				})
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAppendContract_SequentialIds() {
	k := suite.App.ContractregistryKeeper
	contract := types.Contract{
		Description:       "dummy contract",
		Authors:           []string{"hank", "david"},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
		MoreInfoUri:       "google.com",
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}

	// Append two contracts
	firstId := k.AppendContract(suite.Ctx, contract)
	secondId := k.AppendContract(suite.Ctx, contract)

	// Ensure secondId is one greater than firstId
	suite.Equal(secondId, firstId+1)
}

func (suite *KeeperTestSuite) TestGetContract() {
	k := suite.App.ContractregistryKeeper
	ids := suite.PrepareContracts(1)
	// Take from PrepareContracts
	expected := types.Contract{
		Description:       "dummy contract",
		Authors:           []string{"hank", "david"},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
		MoreInfoUri:       "google.com",
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}

	actual, err := k.GetContract(suite.Ctx, ids[0])
	suite.NoError(err)
	suite.Equal(expected, *actual)
}

func (suite *KeeperTestSuite) TestHasContract() {
	k := suite.App.ContractregistryKeeper
	ids := suite.PrepareContracts(1)

	hasData := k.HasContract(suite.Ctx, ids[0])
	suite.True(hasData)
}
