package keeper_test

import (
	"archive/x/contractregistry/types"
)

// SetSigningData test cases
//		Success on implementation case
// 		Failure on id that is greater than or equal to k.getContractCount(ctx)
// 		Failure on id that is less than k.getContractCount(ctx)
// 		Failure on duplicate ID

// GetSigningData test cases
//		What happens on an overflow?

func (suite *KeeperTestSuite) TestSetSigningData() {
	k := suite.App.ContractregistryKeeper
	var defaultData types.RawSigningData
	defaultData.UnmarshalJSON([]byte(`{"test":1}`))
	defaultId := uint64(0)
	defaultContracts := []*types.Contract{{
		Description:       "dummy contract",
		Authors:           []string{"hank", "david"},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Email, Value: "hank@archive.com"},
		MoreInfoUri:       "google.com",
		TemplateUri:       "google.com/template",
		TemplateSchemaUri: "google.com/template-schema",
	}}

	tests := map[string]struct {
		inputContracts []*types.Contract
		inputData      types.RawSigningData
		inputId        uint64
		expErr         bool
		expData        types.RawSigningData
	}{
		"simple_set": {
			inputContracts: defaultContracts,
			inputData:      defaultData,
			inputId:        defaultId,
			expErr:         false,
			expData:        defaultData,
		},
		"no_contract_set": {
			inputContracts: []*types.Contract{},
			inputData:      defaultData,
			inputId:        defaultId,
			expErr:         true,
			expData:        nil,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			for _, contract := range test.inputContracts {
				k.AppendContract(suite.Ctx, *contract)
			}
			err := k.SetSigningData(suite.Ctx, test.inputData, test.inputId)
			if !test.expErr {
				suite.NoError(err)
				actualData, _ := k.GetSigningData(suite.Ctx, test.inputId)
				suite.Equal(test.expData, actualData)
			} else {
				suite.Error(err)
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

func (suite *KeeperTestSuite) TestGetSigningData() {
	k := suite.App.ContractregistryKeeper
	var expected types.RawSigningData
	expected.UnmarshalJSON([]byte("test"))

	ids := suite.PrepareContracts(1)
	k.SetSigningData(suite.Ctx, expected, ids[0])

	actual, err := k.GetSigningData(suite.Ctx, ids[0])
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *KeeperTestSuite) TestHasSigningData() {
	k := suite.App.ContractregistryKeeper
	ids := suite.PrepareContracts(1)
	k.SetSigningData(suite.Ctx, []byte("test"), ids[0])

	hasData := k.HasSigningData(suite.Ctx, ids[0])
	suite.True(hasData)
}
