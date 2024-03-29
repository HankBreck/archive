package keeper_test

import (
	"github.com/HankBreck/archive/x/identity/types"
	"github.com/cosmos/cosmos-sdk/types/query"
)

// TODO:
//		Add test for GetIssuers (paged)

func (suite *KeeperTestSuite) TestSetIssuer() {
	k := suite.App.IdentityKeeper
	issuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
	}
	err := k.SetIssuer(suite.Ctx, issuer)
	suite.NoError(err)
}

func (suite *KeeperTestSuite) TestSetIssuer_DuplicateIssuer() {
	k := suite.App.IdentityKeeper
	issuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
	}
	// First time setting works
	err := k.SetIssuer(suite.Ctx, issuer)
	suite.NoError(err)

	// Second time setting throws an error
	err = k.SetIssuer(suite.Ctx, issuer)
	suite.ErrorIs(err, types.ErrExistingIssuer)
}

func (suite *KeeperTestSuite) TestGetIssuer() {
	// Prepare state
	k := suite.App.IdentityKeeper
	creator := suite.TestAccs[0]
	expectedIssuer := types.Issuer{
		Creator:     creator.String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
	}
	k.SetIssuer(suite.Ctx, expectedIssuer)

	// Retrieve expected issuer object
	actualIssuer, err := k.GetIssuer(suite.Ctx, creator)
	suite.NoError(err)
	suite.Equal(expectedIssuer, *actualIssuer)
}

func (suite *KeeperTestSuite) TestGetIssuers() {
	k := suite.App.IdentityKeeper

	recoveredIssuers := []string{}
	expectedIssuers, _ := suite.MockIssuers(10)

	// Fetch first set of issuers
	issuers, pageRes, err := k.GetIssuers(suite.Ctx, &query.PageRequest{Limit: 5})
	suite.NoError(err)
	suite.NotNil(pageRes.NextKey)
	suite.Len(issuers, 5)
	for _, addr := range issuers {
		recoveredIssuers = append(recoveredIssuers, addr.String())
	}

	// Fetch second set of issuers
	issuers, pageRes, err = k.GetIssuers(suite.Ctx, &query.PageRequest{Limit: 5, Key: pageRes.NextKey})
	suite.NoError(err)
	suite.Nil(pageRes.NextKey)
	suite.Len(issuers, 5)
	for _, addr := range issuers {
		recoveredIssuers = append(recoveredIssuers, addr.String())
	}

	// Ensure expected and recovered issuers match (order unpredictable)
	suite.Len(recoveredIssuers, len(expectedIssuers))
	for _, addr := range expectedIssuers {
		suite.Contains(recoveredIssuers, addr.String())
	}
}

func (suite *KeeperTestSuite) TestHasIssuer() {
	// Prepare state
	k := suite.App.IdentityKeeper
	creator := suite.TestAccs[0]
	expectedIssuer := types.Issuer{
		Creator:     creator.String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
	}
	k.SetIssuer(suite.Ctx, expectedIssuer)

	suite.True(k.HasIssuer(suite.Ctx, creator))
	suite.False(k.HasIssuer(suite.Ctx, suite.TestAccs[1]))
}
