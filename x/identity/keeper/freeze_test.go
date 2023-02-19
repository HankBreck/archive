package keeper_test

func (suite *KeeperTestSuite) TestFreeze() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)

	// Identities are not frozen by default
	suite.False(k.IsFrozen(suite.Ctx, id))

	// Freeze the identity
	err := k.Freeze(suite.Ctx, id)
	suite.NoError(err)

	// Ensure it is frozen in state
	suite.True(k.IsFrozen(suite.Ctx, id))
}
