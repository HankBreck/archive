package keeper_test

import "archive/x/contractregistry/types"

// SetSigningData test cases
//		Success on implementation case
// 		Failure on id that is greater than or equal to k.getContractCount(ctx)
// 		Failure on id that is less than k.getContractCount(ctx)
// 		Failure on duplicate ID

// GetSigningData test cases
//		What happens on an overflow?

func (suite *KeeperTestSuite) TestSetSigningData() {
	k := suite.App.ContractregistryKeeper
	var defaultData types.RawSigningData = []byte(`{"test":1}`)
	defaultId := uint64(0)
	tests := map[string]struct {
		inputData types.RawSigningData
		inputId   uint64
		expErr    bool
		expData   types.RawSigningData
	}{
		"simple_set": {
			inputData: defaultData,
			inputId:   defaultId,
			expErr:    false,
			expData:   defaultData,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
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
