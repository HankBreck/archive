package keeper_test

import (
	"time"

	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestGetCDA() {
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids, _ := suite.PrepareCdas(signers, 1)
	expected := suite.GetCdas(ids)[0]

	cda, err := suite.App.CdaKeeper.GetCDA(suite.Ctx, ids[0])
	suite.NoError(err)
	suite.Equal(*expected, *cda)
}

func (suite *KeeperTestSuite) TestGetCDA_InvalidCdaId() {
	signers := []*sdk.AccAddress{&suite.TestAccs[0]}
	ids, _ := suite.PrepareCdas(signers, 1)

	cda, err := suite.App.CdaKeeper.GetCDA(suite.Ctx, ids[0]+uint64(1))
	suite.EqualError(err, "Invalid CdaId. Please ensure the CDA exists for the given ID.")
	suite.Nil(cda)
}

func (suite *KeeperTestSuite) TestAppendCda() {
	k := suite.App.CdaKeeper
	issuer := suite.TestAccs[0]
	signers := []*sdk.AccAddress{&suite.TestAccs[1]}
	signerId, _ := suite.PrepareCertificate(issuer, signers[0])

	cda := types.CDA{
		Creator:          signers[0].String(),
		SignerIdentities: []uint64{signerId},
		ContractId:       0,
		LegalMetadataUri: "bafkreifbcafazw72o3hogmftvf2bfc7n7t67movnrarx26nyzdz6j6ohpe",
		UtcExpireTime:    time.Date(2100, time.September, 10, 9, 0, 0, 0, time.UTC), // Wednesday, September 1, 2100 9:00:00 AM UTC
		Status:           types.CDA_Voided,
	}
	firstId := k.AppendCDA(suite.Ctx, cda)
	secondId := k.AppendCDA(suite.Ctx, cda)

	cdas := suite.GetCdas([]uint64{firstId, secondId})
	cda.Id = 0
	suite.Equal(cda, *cdas[0])
	cda.Id = 1
	suite.Equal(cda, *cdas[1])
}

func (suite *KeeperTestSuite) TestAppendCda_OverwritesId() {
	k := suite.App.CdaKeeper
	issuer := suite.TestAccs[0]
	signers := []*sdk.AccAddress{&suite.TestAccs[1]}
	signerId, _ := suite.PrepareCertificate(issuer, signers[0])

	cda := types.CDA{
		Creator:          signers[0].String(),
		Id:               uint64(10), // Id should be 0
		SignerIdentities: []uint64{signerId},
		ContractId:       0,
		LegalMetadataUri: "bafkreifbcafazw72o3hogmftvf2bfc7n7t67movnrarx26nyzdz6j6ohpe",
		UtcExpireTime:    time.Date(2100, time.September, 10, 9, 0, 0, 0, time.UTC), // Wednesday, September 1, 2100 9:00:00 AM UTC
		Status:           types.CDA_Voided,
	}
	id := k.AppendCDA(suite.Ctx, cda)
	suite.Equal(uint64(0), id)
}
