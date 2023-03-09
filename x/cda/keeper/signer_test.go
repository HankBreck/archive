package keeper_test

func (suite *KeeperTestSuite) TestAppendSignerCDA() {
	k := suite.App.CdaKeeper
	inputSignerId := uint64(1)
	inputCdaId := uint64(2)

	suite.False(k.SignerHasCda(suite.Ctx, inputSignerId, inputCdaId))
	k.AppendSignerCDA(suite.Ctx, inputSignerId, inputCdaId)
	suite.True(k.SignerHasCda(suite.Ctx, inputSignerId, inputCdaId))
}
