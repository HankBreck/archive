package keeper_test

import (
	"archive/x/contractregistry/types"
)

// AppendContract test cases
//		Success on implementation case
//		Overwrites a pre-set ID
// 		Failure on id that is greater than or equal to k.getContractCount(ctx)
// 		Failure on id that is less than k.getContractCount(ctx)
// 		Failure on duplicate ID

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
				suite.Equal(test.expContract, *actual)
			} else {
				suite.Panics(func() {
					k.AppendContract(suite.Ctx, *test.inputContract)
				})
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDuplicateAppendSigningData() {
	k := suite.App.ContractregistryKeeper

	suite.PrepareContracts(1)

	var firstData types.RawSigningData
	firstData.UnmarshalJSON([]byte("test 1"))
	var secondData types.RawSigningData
	secondData.UnmarshalJSON([]byte("test 2"))

	// Set the first signing data for id 0
	err := k.SetSigningData(suite.Ctx, firstData, 0)
	suite.NoError(err)

	// Try to set the second signing data for the same id
	err = k.SetSigningData(suite.Ctx, secondData, 0)
	suite.Error(err)

	actualData, err := k.GetSigningData(suite.Ctx, 0)
	suite.NoError(err)
	suite.Equal(firstData, actualData)
}

func (suite *KeeperTestSuite) TestGetContract() {
	k := suite.App.ContractregistryKeeper
	var expected types.RawSigningData
	expected.UnmarshalJSON([]byte("test"))

	ids := suite.PrepareContracts(1)
	k.SetSigningData(suite.Ctx, expected, ids[0])

	actual, err := k.GetSigningData(suite.Ctx, ids[0])
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *KeeperTestSuite) TestHasContract() {
	k := suite.App.ContractregistryKeeper
	ids := suite.PrepareContracts(1)
	k.SetSigningData(suite.Ctx, []byte("test"), ids[0])

	hasData := k.HasSigningData(suite.Ctx, ids[0])
	suite.True(hasData)
}
