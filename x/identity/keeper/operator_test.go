package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

func (suite *KeeperTestSuite) TestSetOperator() {
	k := suite.App.IdentityKeeper

	// Setup default values
	defaultOpers := []sdk.AccAddress{suite.TestAccs[2], suite.TestAccs[3]}
	defaultMembers := []sdk.AccAddress{suite.TestAccs[2], suite.TestAccs[3]}

	tests := map[string]struct {
		inputOpers   []sdk.AccAddress
		inputMembers []sdk.AccAddress
		expErr       bool
	}{
		"not_a_member": {
			inputOpers:   defaultOpers,
			inputMembers: []sdk.AccAddress{},
			expErr:       true,
		},
		"set_duplicate": {
			inputOpers:   []sdk.AccAddress{suite.TestAccs[4], suite.TestAccs[4]},
			inputMembers: []sdk.AccAddress{suite.TestAccs[4]},
			expErr:       true,
		},
		"simple_set": {
			inputOpers:   defaultOpers,
			inputMembers: defaultMembers,
			expErr:       false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			// Set up certificate
			issuer := suite.TestAccs[0]
			recipient := suite.TestAccs[1]
			id, _ := suite.PrepareCertificate(issuer, &recipient)
			suite.SetMembers(id, test.inputMembers)

			// Test setting operators
			err := k.SetOperators(suite.Ctx, id, test.inputOpers)
			if test.expErr {
				suite.Error(err)
			} else {
				suite.NoError(err)

				// Verify correct update
				for _, oper := range test.inputOpers {
					hasOp, _ := k.HasOperator(suite.Ctx, id, oper)
					suite.True(hasOp)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestSetOperator_NoCertificate() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	operators := []sdk.AccAddress{suite.TestAccs[0], suite.TestAccs[1]}
	invalidId := uint64(100)
	err := k.SetOperators(suite.Ctx, invalidId, operators)
	suite.Error(err)
}

func setupRemoveOpers(suite *KeeperTestSuite, opers []sdk.AccAddress) uint64 {
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.AddOperators(id, opers)
	return id
}

func (suite *KeeperTestSuite) TestRemoveOperators() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	opers := suite.TestAccs[2:4]
	id := setupRemoveOpers(suite, opers)

	err := k.RemoveOperators(suite.Ctx, id, opers)
	suite.NoError(err)

	// Ensure operators were correctly updated
	operAddrs, _, err := k.GetOperators(suite.Ctx, id, &query.PageRequest{})
	suite.NoError(err)
	suite.Len(operAddrs, 0)
}

func (suite *KeeperTestSuite) TestGetOperators() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	opers := suite.TestAccs[2:4]
	suite.AddOperators(id, opers)

	operAddrs, _, err := k.GetOperators(suite.Ctx, id, &query.PageRequest{})
	suite.NoError(err)
	for _, addr := range opers {
		suite.Contains(operAddrs, addr.String())
	}
}

func (suite *KeeperTestSuite) TestGetOperators_WithLimit() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	opers := suite.TestAccs[2:15]
	suite.AddOperators(id, opers)

	operAddrs, _, err := k.GetOperators(suite.Ctx, id, &query.PageRequest{Limit: 10})
	suite.NoError(err)
	suite.Len(operAddrs, 10)
}

func (suite *KeeperTestSuite) TestHasOperator() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	operator := suite.TestAccs[2]
	id, _ := suite.PrepareCertificate(issuer, &recipient)

	// Test unset operator
	hasOp, err := k.HasOperator(suite.Ctx, id, operator)
	suite.NoError(err)
	suite.False(hasOp)

	// Test set operator
	suite.AddOperators(id, []sdk.AccAddress{operator})
	hasOp, err = k.HasOperator(suite.Ctx, id, operator)
	suite.NoError(err)
	suite.True(hasOp)
}

func (suite *KeeperTestSuite) TestHasOperator_NoCertificate() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	operator := suite.TestAccs[2]
	invalidId := uint64(100)
	hasOp, err := k.HasOperator(suite.Ctx, invalidId, operator)
	suite.Error(err)
	suite.False(hasOp)
}
