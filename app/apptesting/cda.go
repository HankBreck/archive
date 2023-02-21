package apptesting

import (
	"time"

	"github.com/HankBreck/archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestHelper) ApproveCda(cdaId uint64, owner *sdk.AccAddress) error {
	k := s.App.CdaKeeper

	goCtx := sdk.WrapSDKContext(s.Ctx)
	res, err := k.Cda(goCtx, &types.QueryCdaRequest{
		Id: cdaId,
	})
	if err != nil {
		return err
	}

	signingData, err := k.GetSigningData(s.Ctx, cdaId)
	if err != nil {
		return err
	}

	msg := types.MsgApproveCda{
		Creator:     owner.String(),
		CdaId:       res.Cda.Id,
		SigningData: signingData,
	}
	err = k.SetApproval(s.Ctx, &msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *KeeperTestHelper) PrepareCdasForOwner(signers []*sdk.AccAddress, count int) []uint64 {
	ids := make([]uint64, count)
	k := s.App.CdaKeeper
	signingParties := make([]string, len(signers))
	for i, signer := range signers {
		signingParties[i] = signer.String()
	}

	for i := 0; i < count; i++ {
		var cda = types.CDA{
			Creator:          signers[0].String(),
			SigningParties:   signingParties,
			ContractId:       0,
			LegalMetadataUri: "bafkreifbcafazw72o3hogmftvf2bfc7n7t67movnrarx26nyzdz6j6ohpe",
			UtcExpireTime:    time.Date(2100, time.September, 10, 9, 0, 0, 0, time.UTC), // Wednesday, September 1, 2100 9:00:00 AM UTC
			Status:           types.CDA_Pending,
		}

		// Store CDA & grab cda id
		id := k.AppendCDA(s.Ctx, cda)
		for _, signer := range cda.SigningParties {
			err := k.AppendOwnerCDA(s.Ctx, signer, id)
			if err != nil {
				panic(err)
			}
		}

		err := k.SetSigningData(s.Ctx, id, s.GetSigningData())
		if err != nil {
			panic(err)
		}

		ids[i] = id
	}
	return ids
}

func (s *KeeperTestHelper) PrepareVoidedCdaForSigners(signers []*sdk.AccAddress) uint64 {
	k := s.App.CdaKeeper
	signingParties := make([]string, len(signers))
	for i, signer := range signers {
		signingParties[i] = signer.String()
	}
	cda := types.CDA{
		Creator:          signers[0].String(),
		SigningParties:   signingParties,
		ContractId:       0,
		LegalMetadataUri: "bafkreifbcafazw72o3hogmftvf2bfc7n7t67movnrarx26nyzdz6j6ohpe",
		UtcExpireTime:    time.Date(2100, time.September, 10, 9, 0, 0, 0, time.UTC), // Wednesday, September 1, 2100 9:00:00 AM UTC
		Status:           types.CDA_Voided,
	}

	// CreateCda() logic
	id := k.AppendCDA(s.Ctx, cda)
	for _, signer := range cda.SigningParties {
		err := k.AppendOwnerCDA(s.Ctx, signer, id)
		if err != nil {
			panic(err)
		}
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

func (s *KeeperTestHelper) GetCdas(ids []uint64) []*types.CDA {
	k := s.App.CdaKeeper
	result := make([]*types.CDA, len(ids))
	goCtx := sdk.WrapSDKContext(s.Ctx)

	for i, id := range ids {
		req := types.QueryCdaRequest{Id: id}
		res, err := k.Cda(goCtx, &req)
		if err != nil {
			panic(err)
		}
		if res == nil {
			panic("Could not fetch CDA!")
		}
		result[i] = res.Cda
	}
	return result
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
