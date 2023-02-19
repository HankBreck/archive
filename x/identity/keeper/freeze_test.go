package keeper_test

import "github.com/HankBreck/archive/x/identity/types"

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

func (suite *KeeperTestSuite) TestFreeze_InvalidCertificate() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	invalidId := id + 1

	// Try to freeze a nonexistent identity
	err := k.Freeze(suite.Ctx, invalidId)
	suite.EqualError(err, types.ErrNonexistentCertificate.Wrapf("certificate not found for id %d", invalidId).Error())

	// Ensure it did not update state
	suite.False(k.IsFrozen(suite.Ctx, invalidId))
}
