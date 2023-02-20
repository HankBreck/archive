package keeper_test

import (
	"github.com/HankBreck/archive/x/identity/types"
)

func (suite *KeeperTestSuite) TestAppendCertificate() {
	k := suite.App.IdentityKeeper

	// Register the issuer
	issuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
	}
	k.SetIssuer(suite.Ctx, issuer)

	// Add new certificate
	certificate := types.Certificate{
		Id:                0,
		IssuerAddress:     issuer.Creator,
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}
	id := k.AppendCertificate(suite.Ctx, certificate)

	// Assert the expected certificate is written to state
	actual, _ := k.GetCertificate(suite.Ctx, id)
	suite.Equal(certificate, *actual)
}

func (suite *KeeperTestSuite) TestAppendCertificate_OverwritePresetId() {
	k := suite.App.IdentityKeeper

	// Register the issuer
	issuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
	}
	k.SetIssuer(suite.Ctx, issuer)

	// Add new certificate
	certificate := types.Certificate{
		Id:                10, // should be overwritten
		IssuerAddress:     issuer.Creator,
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}
	id := k.AppendCertificate(suite.Ctx, certificate)
	suite.NotEqual(certificate.Id, id)
	suite.Equal(uint64(0), id)
}

func (suite *KeeperTestSuite) TestGetCertificate() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	id, _ := suite.PrepareCertificate(issuer, nil)
	expected := types.Certificate{
		Id:                uint64(0),
		IssuerAddress:     issuer.String(),
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}

	actual, err := k.GetCertificate(suite.Ctx, id)
	suite.NoError(err)
	suite.Equal(expected, *actual)
}

func (suite *KeeperTestSuite) TestGetCertificate_InvalidId() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	id, _ := suite.PrepareCertificate(issuer, nil)

	cert, err := k.GetCertificate(suite.Ctx, id+1)
	suite.Error(err)
	suite.Nil(cert)
}

func (suite *KeeperTestSuite) TestHasCertificate() {
	k := suite.App.IdentityKeeper
	issuer := suite.TestAccs[0]
	id, _ := suite.PrepareCertificate(issuer, nil)

	suite.True(k.HasCertificate(suite.Ctx, id))
	suite.False(k.HasCertificate(suite.Ctx, id+5))
}
