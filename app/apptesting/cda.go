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
	}
	id := k.AppendContract(s.Ctx, types.Contract{
		Description:       "",
		Authors:           []string{},
		ContactInfo:       &types.ContactInfo{Method: types.ContactMethod_Phone, Value: "(123) 456-7890"},
		MoreInfoUri:       "",
		TemplateUri:       "",
		TemplateSchemaUri: "",
	})
	contract.Id = id

	return contract
}

func (s *KeeperTestHelper) GetTemplateCda(creator sdk.AccAddress, signerIds []uint64) types.CDA {
	return types.CDA{
		Creator:          creator.String(),
		SignerIdentities: signerIds,
		ContractId:       0,
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
