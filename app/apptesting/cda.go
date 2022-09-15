package apptesting

import (
	"archive/x/cda/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (s *KeeperTestHelper) PrepareCdasForOwner(owner sdk.AccAddress, count int) []uint64 {
	ids := make([]uint64, count)
	k := s.App.CdaKeeper
	ownerships := make([]*types.Ownership, 1)
	ownerships[0] = &types.Ownership{
		Owner:     owner.String(),
		Ownership: uint64(100),
	}

	for i := 0; i < count; i++ {
		var cda = types.CDA{
			Creator:    owner.String(),
			Cid:        "QmSrnQXUtGqsVRcgY93CdWXf8GPE9Zjj7Tg3SZUgLKDN5W",
			Ownership:  ownerships,
			Expiration: 4123503529000, // Wednesday, September 1, 2100 5:38:49 PM
		}

		// Store CDA & grab cda id
		id := k.AppendCDA(s.Ctx, cda)
		for i := range cda.Ownership {
			owner := cda.Ownership[i]
			err := k.AppendOwnerCDA(s.Ctx, owner.Owner, id)
			// TODO: check if we need some sort of transaction/rollback option in case this fails
			if err != nil {
				panic(err)
			}
		}
		ids[i] = id
	}
	return ids
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
