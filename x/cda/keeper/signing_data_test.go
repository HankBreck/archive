package keeper_test

import (
	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetSigningData test cases
//		Success on implementation case
// 		Failure on id that is greater than or equal to k.getContractCount(ctx)
// 		Failure on id that is less than k.getContractCount(ctx)
// 		Failure on duplicate ID

// GetSigningData test cases
//		What happens on an overflow?

func (suite *KeeperTestSuite) TestSetTemplateSigningData() {
	k := suite.App.CdaKeeper
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
			err := k.SetTemplateSigningData(suite.Ctx, test.inputData, test.inputId)
			if !test.expErr {
				suite.NoError(err)
				actualData, _ := k.GetTemplateSigningData(suite.Ctx, test.inputId)
				suite.Equal(test.expData, actualData)
			} else {
				suite.Error(err)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestDuplicateSetTemplateSigningData() {
	k := suite.App.CdaKeeper

	suite.PrepareContracts(1)

	var firstData types.RawSigningData
	firstData.UnmarshalJSON([]byte("test 1"))
	var secondData types.RawSigningData
	secondData.UnmarshalJSON([]byte("test 2"))

	// Set the first signing data for id 0
	err := k.SetTemplateSigningData(suite.Ctx, firstData, 0)
	suite.NoError(err)

	// Try to set the second signing data for the same id
	err = k.SetTemplateSigningData(suite.Ctx, secondData, 0)
	suite.Error(err)

	actualData, err := k.GetTemplateSigningData(suite.Ctx, 0)
	suite.NoError(err)
	suite.Equal(firstData, actualData)
}

func (suite *KeeperTestSuite) TestGetTemplateSigningData() {
	k := suite.App.CdaKeeper
	var expected types.RawSigningData
	expected.UnmarshalJSON([]byte("test"))

	ids := suite.PrepareContracts(1)
	k.SetTemplateSigningData(suite.Ctx, expected, ids[0])

	actual, err := k.GetTemplateSigningData(suite.Ctx, ids[0])
	suite.NoError(err)
	suite.Equal(expected, actual)
}

func (suite *KeeperTestSuite) TestHasTemplateSigningData() {
	k := suite.App.CdaKeeper
	ids := suite.PrepareContracts(1)
	k.SetTemplateSigningData(suite.Ctx, []byte("test"), ids[0])

	hasData := k.HasTemplateSigningData(suite.Ctx, ids[0])
	suite.True(hasData)
}

func (suite *KeeperTestSuite) TestMatchesTemplateSigningDataSchema() {
	k := suite.App.CdaKeeper
	var signingDataSchema types.RawSigningData
	signingDataSchema.UnmarshalJSON([]byte(getTestSchema()))
	res, _ := suite.msgServer.RegisterContract(sdk.WrapSDKContext(suite.Ctx), &types.MsgRegisterContract{
		Creator:     string(suite.TestAccs[0]),
		Description: "",
		Authors:     []string{},
		ContactInfo: &types.ContactInfo{
			Method: types.ContactMethod_Email,
			Value:  "breckenridgeh2@gmail.com",
		},
		MoreInfoUri:       "",
		SigningDataSchema: signingDataSchema,
		TemplateUri:       "",
		TemplateSchemaUri: "",
	})

	var inputData types.RawSigningData
	inputData.UnmarshalJSON([]byte(getTestDoc()))
	matches, err := k.MatchesTemplateSigningDataSchema(suite.Ctx, res.Id, inputData)
	suite.NoError(err)
	suite.True(matches)
}

func (suite *KeeperTestSuite) TestMatchesTemplateSigningDataSchema_NoMatch() {
	k := suite.App.CdaKeeper
	var signingDataSchema types.RawSigningData
	signingDataSchema.UnmarshalJSON(getTestSchema())
	res, _ := suite.msgServer.RegisterContract(sdk.WrapSDKContext(suite.Ctx), &types.MsgRegisterContract{
		Creator:     string(suite.TestAccs[0]),
		Description: "",
		Authors:     []string{},
		ContactInfo: &types.ContactInfo{
			Method: types.ContactMethod_Email,
			Value:  "breckenridgeh2@gmail.com",
		},
		MoreInfoUri:       "",
		SigningDataSchema: signingDataSchema,
		TemplateUri:       "",
		TemplateSchemaUri: "",
	})

	var inputData types.RawSigningData
	inputData.UnmarshalJSON([]byte(`
	{
		"notOwnerships": [
			{ "owner": "address", "ownership_proportion": 1 },
			{ "owner": "address2", "ownership_proportion": 99 }
		]
	}`))
	matches, err := k.MatchesTemplateSigningDataSchema(suite.Ctx, res.Id, inputData)
	suite.Error(err)
	suite.False(matches)
}

func (suite *KeeperTestSuite) TestMatchesTemplateSigningDataSchema_InvalidJSONSchema() {
	k := suite.App.CdaKeeper
	var signingDataSchema types.RawSigningData
	signingDataSchema.UnmarshalJSON([]byte(`"hello": "world"`)) // missing braces around JSON
	res, _ := suite.msgServer.RegisterContract(sdk.WrapSDKContext(suite.Ctx), &types.MsgRegisterContract{
		Creator:     string(suite.TestAccs[0]),
		Description: "",
		Authors:     []string{},
		ContactInfo: &types.ContactInfo{
			Method: types.ContactMethod_Email,
			Value:  "breckenridgeh2@gmail.com",
		},
		MoreInfoUri:       "",
		SigningDataSchema: signingDataSchema,
		TemplateUri:       "",
		TemplateSchemaUri: "",
	})

	var inputData types.RawSigningData
	inputData.UnmarshalJSON(getTestDoc())
	matches, err := k.MatchesTemplateSigningDataSchema(suite.Ctx, res.Id, inputData)
	suite.Error(err)
	suite.False(matches)
}

func getTestDoc() []byte {
	return []byte(`
	{
		"ownerships": [
			{ "owner": "address", "ownership_proportion": 1 },
			{ "owner": "address2", "ownership_proportion": 99 }
		]
	}`)
}

func getTestSchema() []byte {
	return []byte(`
	{
		"$schema": "https://json-schema.org/draft/2019-09/schema",
		"$id": "http://example.com/example.json",
		"type": "object",
		"default": {},
		"title": "Root Schema",
		"required": [
			"ownerships"
		],
		"properties": {
			"ownerships": {
				"type": "array",
				"default": [],
				"title": "The ownerships Schema",
				"items": {
					"type": "object",
					"default": {},
					"title": "A Schema",
					"required": [
						"owner",
						"ownership_proportion"
					],
					"properties": {
						"owner": {
							"type": "string",
							"default": "",
							"title": "The owner Schema",
							"examples": [
								"address"
							]
						},
						"ownership_proportion": {
							"type": "integer",
							"default": 0,
							"title": "The ownership_proportion Schema",
							"examples": [
								1
							]
						}
					},
					"examples": [{
						"owner": "address",
						"ownership_proportion": 1
					}]
				},
				"examples": [
					[{
						"owner": "address",
						"ownership_proportion": 1
					}]
				]
			}
		},
		"examples": [{
			"ownerships": [{
				"owner": "address",
				"ownership_proportion": 1
			}]
		}]
	}`)
}
