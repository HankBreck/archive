package keeper_test

import (
	"github.com/HankBreck/archive/x/identity/types"
)

// TODO:
//		Add test for GetIssuers (paged)

func (suite *KeeperTestSuite) TestSetIssuer() {
	k := suite.App.IdentityKeeper
	issuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
		Cost:        0,
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
		Cost:        0,
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
	creator := suite.TestAccs[0].String()
	expectedIssuer := types.Issuer{
		Creator:     creator,
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
		Cost:        0,
	}
	k.SetIssuer(suite.Ctx, expectedIssuer)

	// Retrieve expected issuer object
	actualIssuer, err := k.GetIssuer(suite.Ctx, creator)
	suite.NoError(err)
	suite.Equal(expectedIssuer, *actualIssuer)
}

func (suite *KeeperTestSuite) TestHasIssuer() {
	// Prepare state
	k := suite.App.IdentityKeeper
	creator := suite.TestAccs[0].String()
	expectedIssuer := types.Issuer{
		Creator:     creator,
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
		Cost:        0,
	}
	k.SetIssuer(suite.Ctx, expectedIssuer)

	suite.True(k.HasIssuer(suite.Ctx, creator))
	suite.False(k.HasIssuer(suite.Ctx, suite.TestAccs[1].String()))
}
