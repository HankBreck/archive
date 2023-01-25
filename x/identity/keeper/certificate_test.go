package keeper_test

import (
	"github.com/HankBreck/archive/x/identity/types"
)

func (suite *KeeperTestSuite) TestAppendCertificate() {
	k := suite.App.IdentityKeeper

	// Setup default values
	defaultIssuer := types.Issuer{
		Creator:     suite.TestAccs[0].String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com/more-info",
		Cost:        0,
	}
	defaultCert := types.Certificate{
		IssuerAddress:     defaultIssuer.Creator,
		Salt:              "salt",
		MetadataSchemaUri: "google.com/metadata-schema",
		Hashes: []*types.HashEntry{
			{Field: "field1", Hash: "hash1"},
			{Field: "field2", Hash: "hash2"},
		},
	}

	tests := map[string]struct {
		inputIssuer      *types.Issuer
		inputCertificate *types.Certificate
		expPanic         bool
		expCertificate   types.Certificate
	}{
		"simple_set": {
			inputIssuer:      &defaultIssuer,
			inputCertificate: &defaultCert,
			expPanic:         false,
			expCertificate: types.Certificate{
				Id:                uint64(0),
				IssuerAddress:     defaultIssuer.Creator,
				Salt:              "salt",
				MetadataSchemaUri: "google.com/metadata-schema",
				Hashes: []*types.HashEntry{
					{Field: "field1", Hash: "hash1"},
					{Field: "field2", Hash: "hash2"},
				},
			},
		},
		"overwrite_preset_id": {
			inputIssuer: &defaultIssuer,
			inputCertificate: &types.Certificate{
				Id:                uint64(10),
				IssuerAddress:     defaultIssuer.Creator,
				Salt:              "salt",
				MetadataSchemaUri: "google.com/metadata-schema",
				Hashes: []*types.HashEntry{
					{Field: "field1", Hash: "hash1"},
					{Field: "field2", Hash: "hash2"},
				},
			},
			expPanic: false,
			expCertificate: types.Certificate{
				Id:                uint64(1), // sequential IDs
				IssuerAddress:     defaultIssuer.Creator,
				Salt:              "salt",
				MetadataSchemaUri: "google.com/metadata-schema",
				Hashes: []*types.HashEntry{
					{Field: "field1", Hash: "hash1"},
					{Field: "field2", Hash: "hash2"},
				},
			},
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			// Register the issuer
			k.SetIssuer(suite.Ctx, *test.inputIssuer)

			if test.expPanic {
				suite.Panics(func() { k.AppendCertificate(suite.Ctx, *test.inputCertificate) })
			} else {
				id := k.AppendCertificate(suite.Ctx, *test.inputCertificate)

				// Assert the expected certificate is written to state
				actual, _ := k.GetCertificate(suite.Ctx, id)
				suite.Equal(test.expCertificate, *actual)
			}
		})
	}
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
