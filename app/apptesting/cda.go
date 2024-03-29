package apptesting

import (
	"time"

	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestHelper) ApproveCda(signer sdk.AccAddress, cdaId uint64, signerId uint64) error {
	k := s.App.CdaKeeper
	cda, err := k.GetCDA(s.Ctx, cdaId)
	if err != nil {
		return err
	}

	err = k.SetApproval(s.Ctx, cda.Id, signerId)
	if err != nil {
		return err
	}

	return nil
}

func (s *KeeperTestHelper) PrepareCdas(signers []*sdk.AccAddress, count int) ([]uint64, []uint64) {
	cdaIds := make([]uint64, count)
	k := s.App.CdaKeeper
	issuer := s.TestAccs[0]
	signerIds := make([]uint64, len(signers))
	for i, signer := range signers {
		signerIds[i], _ = s.PrepareCertificate(issuer, signer)
	}

	for i := 0; i < count; i++ {
		cda := s.GetTemplateCda(*signers[0], signerIds)

		// Store CDA & grab cda id
		id := k.AppendCDA(s.Ctx, cda)
		for _, signerId := range cda.SignerIdentities {
			k.AppendSignerCDA(s.Ctx, signerId, id)
		}

		err := k.SetSigningData(s.Ctx, id, s.GetSigningData())
		if err != nil {
			panic(err)
		}

		cdaIds[i] = id
	}
	return cdaIds, signerIds
}

func (s *KeeperTestHelper) PrepareVoidedCdaForSigners(signers []*sdk.AccAddress) uint64 {
	k := s.App.CdaKeeper
	issuer := s.TestAccs[0]
	signerIds := make([]uint64, len(signers))
	for i, signer := range signers {
		signerIds[i], _ = s.PrepareCertificate(issuer, signer)
	}

	cda := s.GetTemplateCda(*signers[0], signerIds)

	// Store CDA & grab cda id
	id := k.AppendCDA(s.Ctx, cda)
	for _, signerId := range cda.SignerIdentities {
		k.AppendSignerCDA(s.Ctx, signerId, id)
	}

	err := k.SetSigningData(s.Ctx, id, s.GetSigningData())
	if err != nil {
		panic(err)
	}

	return id
}

func (s *KeeperTestHelper) PrepareContract() types.Contract {
	k := s.App.CdaKeeper
	contract := types.Contract{
		Description:       "",
		Authors:           []string{},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Phone, Value: "(123) 456-7890"},
		MoreInfoUri:       "",
		TemplateUri:       "",
		TemplateSchemaUri: "",
		WitnessCodeId:     0,
	}
	id := k.AppendContract(s.Ctx, types.Contract{
		Description:       "",
		Authors:           []string{},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Phone, Value: "(123) 456-7890"},
		MoreInfoUri:       "",
		TemplateUri:       "",
		TemplateSchemaUri: "",
		WitnessCodeId:     0,
	})
	contract.Id = id

	return contract
}

func (s *KeeperTestHelper) PrepareContractWithSchema(codeId uint64, schema types.RawSigningData) (*types.Contract, error) {
	k := s.App.CdaKeeper
	contract := types.Contract{
		Description:       "",
		Authors:           []string{},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Phone, Value: "(123) 456-7890"},
		MoreInfoUri:       "",
		TemplateUri:       "",
		TemplateSchemaUri: "",
		WitnessCodeId:     codeId,
	}
	id := k.AppendContract(s.Ctx, contract)
	contract.Id = id
	err := k.SetSigningDataSchema(s.Ctx, id, schema)
	if err != nil {
		return nil, err
	}

	return &contract, nil
}

func (s *KeeperTestHelper) GetTemplateCda(creator sdk.AccAddress, signerIds []uint64) types.CDA {
	return types.CDA{
		Creator:          creator.String(),
		SignerIdentities: signerIds,
		ContractId:       1,
		LegalMetadataUri: "bafkreifbcafazw72o3hogmftvf2bfc7n7t67movnrarx26nyzdz6j6ohpe",
		UtcExpireTime:    time.Date(2100, time.September, 10, 9, 0, 0, 0, time.UTC), // Wednesday, September 1, 2100 9:00:00 AM UTC
		Status:           types.CDA_Pending,
	}
}

func (s *KeeperTestHelper) GetSigningData() types.RawSigningData {
	var data types.RawSigningData
	data.UnmarshalJSON([]byte(`
	{
		"ownerships": [
			{ "owner": "address", "ownership_proportion": 1 },
			{ "owner": "address2", "ownership_proportion": 99 }
		]
	}`))
	return data
}

func (suite *KeeperTestHelper) GetTestDoc() []byte {
	return []byte(`
	{
		"ownerships": [
			{ "owner": "address", "ownership_proportion": 1 },
			{ "owner": "address2", "ownership_proportion": 99 }
		]
	}`)
}

func (suite *KeeperTestHelper) GetTestSchema() []byte {
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
