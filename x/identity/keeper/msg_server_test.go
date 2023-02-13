package keeper_test

import (
	"github.com/HankBreck/archive/x/identity/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Register Issuer
func (suite *KeeperTestSuite) TestRegisterIssuer() {
	creator := suite.TestAccs[0]
	defaultMsg := types.MsgRegisterIssuer{
		Creator:     creator.String(),
		Name:        "Test name",
		MoreInfoUri: "google.com",
	}
	tests := map[string]struct {
		inputMsg *types.MsgRegisterIssuer
		expErr   bool
	}{
		"nil_message": {
			inputMsg: nil,
			expErr:   true,
		},
		"invalid_creator": {
			inputMsg: &types.MsgRegisterIssuer{
				Creator:     "creator",
				Name:        "Test name",
				MoreInfoUri: "google.com",
			},
			expErr: true,
		},
		"simple_register": {
			inputMsg: &defaultMsg,
			expErr:   false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx
			msgServer := suite.msgServer

			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			_, err := msgServer.RegisterIssuer(sdk.WrapSDKContext(suite.Ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.AssertEventEmitted(suite.Ctx, types.TypeMsgRegisterIssuer, 0)
			} else {
				suite.NoError(err)
				suite.AssertEventEmitted(suite.Ctx, types.TypeMsgRegisterIssuer, 1)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRegisterIssuer_DuplicateIssuer() {
	// Setup Test
	goCtx := sdk.WrapSDKContext(suite.Ctx)
	msgServer := suite.msgServer
	creator := suite.TestAccs[0]
	msg := &types.MsgRegisterIssuer{
		Creator:     creator.String(),
		Name:        "Test Issuer",
		MoreInfoUri: "google.com",
	}

	// Register creator as issuer for first time
	_, err := msgServer.RegisterIssuer(goCtx, msg)
	suite.NoError(err)

	// Register creator as issuer for second time
	_, err = msgServer.RegisterIssuer(goCtx, msg)
	suite.EqualError(err, types.ErrExistingIssuer.Error())
}

// Issuer Certificate
func (suite *KeeperTestSuite) TestIssueCertificate() {
	creator := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	defaultMsg := types.MsgIssueCertificate{
		Creator:           creator.String(),
		Recipient:         recipient.String(),
		Salt:              "salt",
		MetadataSchemaUri: "google.com",
		Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}},
	}
	tests := map[string]struct {
		inputIssuer *sdk.AccAddress
		inputMsg    *types.MsgIssueCertificate
		expErr      bool
	}{
		"nil_message": {
			inputIssuer: &creator,
			inputMsg:    nil,
			expErr:      true,
		},
		"invalid_creator": {
			inputIssuer: &creator,
			inputMsg: &types.MsgIssueCertificate{
				Creator:           "creator",
				Recipient:         recipient.String(),
				Salt:              "salt",
				MetadataSchemaUri: "google.com",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: "Bar"}},
			},
			expErr: true,
		},
		"invalid_issuer": {
			inputIssuer: nil,
			inputMsg:    &defaultMsg,
			expErr:      true,
		},
		"invalid_recipient": {
			inputIssuer: &creator,
			inputMsg: &types.MsgIssueCertificate{
				Creator:           creator.String(),
				Recipient:         "recipient",
				Salt:              "salt",
				MetadataSchemaUri: "google.com",
				Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}},
			},
			expErr: true,
		},
		"simple_register": {
			inputIssuer: &creator,
			inputMsg:    &defaultMsg,
			expErr:      false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			k := suite.App.IdentityKeeper
			ctx := suite.Ctx
			msgServer := suite.msgServer

			// Register inputIssuer as an issuer
			if test.inputIssuer != nil {
				k.SetIssuer(ctx, types.Issuer{
					Creator:     test.inputIssuer.String(),
					Name:        "Test Issuer",
					MoreInfoUri: "google.com",
				})
			}

			// Reset event manager
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			// Issue the Certificate
			res, err := msgServer.IssueCertificate(sdk.WrapSDKContext(ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.Nil(res)
				suite.AssertEventEmitted(ctx, types.TypeMsgIssueCertificate, 0)
			} else {
				// Assert correct output
				suite.NotNil(res)
				suite.NoError(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgIssueCertificate, 1)
				certificateId := res.Id
				recipientAddr, _ := sdk.AccAddressFromBech32(test.inputMsg.Recipient)
				suite.Equal(uint64(0), certificateId)

				// Assert correct state transition
				// Initial recipient is a pending member
				hasMember, err := k.HasPendingMember(ctx, certificateId, recipientAddr)
				suite.NoError(err)
				suite.True(hasMember)
				// Initial recipient is an operator
				hasOperator, err := k.HasOperator(ctx, certificateId, recipientAddr)
				suite.NoError(err)
				suite.True(hasOperator)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestIssueCertificate_IncrementalIds() {
	k := suite.App.IdentityKeeper
	msgServer := suite.msgServer
	creator := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	msg := &types.MsgIssueCertificate{
		Creator:           creator.String(),
		Recipient:         recipient.String(),
		Salt:              "salt",
		MetadataSchemaUri: "google.com",
		Hashes:            []*types.HashEntry{{Field: "foo", Hash: "bar"}},
	}
	k.SetIssuer(suite.Ctx, types.Issuer{
		Creator:     creator.String(),
		Name:        "test name",
		MoreInfoUri: "google.com",
	})

	// Capture first ID
	res, err := msgServer.IssueCertificate(sdk.WrapSDKContext(suite.Ctx), msg)
	suite.NoError(err)
	firstId := res.Id

	// Capture second ID
	res, err = msgServer.IssueCertificate(sdk.WrapSDKContext(suite.Ctx), msg)
	suite.NoError(err)
	secondId := res.Id

	// Ensure secondId is one higher than firstId
	suite.Equal(secondId, firstId+1)
}

// Accept Identity
func (suite *KeeperTestSuite) TestAcceptIdentity() {
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	defaultMsg := types.MsgAcceptIdentity{
		Creator: recipient.String(),
		Id:      uint64(0), // Overwritten inside test
	}
	tests := map[string]struct {
		inputIssuer    *sdk.AccAddress
		inputRecipient *sdk.AccAddress
		inputMsg       *types.MsgAcceptIdentity
		expErr         bool
	}{
		"nil_message": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg:       nil,
			expErr:         true,
		},
		"invalid_creator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg: &types.MsgAcceptIdentity{
				Creator: "creator",
				Id:      uint64(0),
			},
			expErr: true,
		},
		"nonpending_sender": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg: &types.MsgAcceptIdentity{
				Creator: suite.TestAccs[8].String(),
				Id:      uint64(0),
			},
			expErr: true,
		},
		"simple_register": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg:       &defaultMsg,
			expErr:         false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx
			k := suite.App.IdentityKeeper
			msgServer := suite.msgServer
			id, _ := suite.PrepareCertificate(*test.inputIssuer, test.inputRecipient)
			if test.inputMsg != nil {
				test.inputMsg.Id = id
			}

			// Reset event manager
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			// Test AcceptIdentity
			res, err := msgServer.AcceptIdentity(sdk.WrapSDKContext(ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgAcceptIdentity, 0)
				hasPending, _ := k.HasPendingMember(ctx, id, *test.inputRecipient)
				suite.True(hasPending)
				hasAccepted, _ := k.HasMember(ctx, id, *test.inputRecipient)
				suite.False(hasAccepted)
			} else {
				suite.NoError(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgAcceptIdentity, 1)
				suite.NotNil(res)
				hasPending, _ := k.HasPendingMember(ctx, id, *test.inputRecipient)
				hasAccepted, _ := k.HasMember(ctx, id, *test.inputRecipient)
				hasOper, _ := k.HasOperator(ctx, id, *test.inputRecipient)
				suite.False(hasPending)
				suite.True(hasAccepted)
				suite.True(hasOper)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestAcceptIdentity_InvalidCertificateId() {
	suite.SetupTest()
	ctx := suite.Ctx
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	msgServer := suite.msgServer
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	invalidId := id + 1

	res, err := msgServer.AcceptIdentity(sdk.WrapSDKContext(ctx), &types.MsgAcceptIdentity{
		Creator: recipient.String(),
		Id:      invalidId,
	})
	suite.Error(err)
	suite.Nil(res)
}

// Reject Identity
func (suite *KeeperTestSuite) TestRejectIdentity() {
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	defaultMsg := types.MsgRejectIdentity{
		Creator: recipient.String(),
		Id:      uint64(0), // Overwritten inside test
	}
	tests := map[string]struct {
		inputIssuer    *sdk.AccAddress
		inputRecipient *sdk.AccAddress
		inputMsg       *types.MsgRejectIdentity
		expErr         bool
	}{
		"nil_message": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg:       nil,
			expErr:         true,
		},
		"invalid_creator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg: &types.MsgRejectIdentity{
				Creator: "creator",
				Id:      uint64(0),
			},
			expErr: true,
		},
		"nonpending_sender": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg: &types.MsgRejectIdentity{
				Creator: suite.TestAccs[8].String(),
				Id:      uint64(0),
			},
			expErr: true,
		},
		"simple_register": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg:       &defaultMsg,
			expErr:         false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx
			k := suite.App.IdentityKeeper
			msgServer := suite.msgServer
			id, _ := suite.PrepareCertificate(*test.inputIssuer, test.inputRecipient)
			if test.inputMsg != nil {
				test.inputMsg.Id = id
			}

			// Reset event manager
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			// Test RejectIdentity
			res, err := msgServer.RejectIdentity(sdk.WrapSDKContext(ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgRejectIdentity, 0)
				hasPending, _ := k.HasPendingMember(ctx, id, *test.inputRecipient)
				suite.True(hasPending)
				hasAccepted, _ := k.HasMember(ctx, id, *test.inputRecipient)
				suite.False(hasAccepted)
			} else {
				suite.NoError(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgRejectIdentity, 1)
				suite.NotNil(res)
				hasPending, _ := k.HasPendingMember(ctx, id, *test.inputRecipient)
				hasAccepted, _ := k.HasMember(ctx, id, *test.inputRecipient)
				hasOper, _ := k.HasOperator(ctx, id, *test.inputRecipient)
				suite.False(hasPending)
				suite.False(hasAccepted)
				suite.False(hasOper)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRejectIdentity_InvalidCertificateId() {
	suite.SetupTest()
	ctx := suite.Ctx
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	msgServer := suite.msgServer
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	invalidId := id + 1

	res, err := msgServer.RejectIdentity(sdk.WrapSDKContext(ctx), &types.MsgRejectIdentity{
		Creator: recipient.String(),
		Id:      invalidId,
	})
	suite.Error(err)
	suite.Nil(res)
}

// Renounce Identity
func (suite *KeeperTestSuite) TestRenounceIdentity() {
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	defaultMsg := types.MsgRenounceIdentity{
		Creator: recipient.String(),
		Id:      uint64(0), // Overwritten inside test
	}
	tests := map[string]struct {
		inputIssuer    *sdk.AccAddress
		inputRecipient *sdk.AccAddress
		inputMsg       *types.MsgRenounceIdentity
		expErr         bool
	}{
		"nil_message": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg:       nil,
			expErr:         true,
		},
		"invalid_creator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg: &types.MsgRenounceIdentity{
				Creator: "creator",
				Id:      uint64(0),
			},
			expErr: true,
		},
		"nonmember_sender": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg: &types.MsgRenounceIdentity{
				Creator: suite.TestAccs[8].String(),
				Id:      uint64(0),
			},
			expErr: true,
		},
		"simple_register": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMsg:       &defaultMsg,
			expErr:         false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx
			k := suite.App.IdentityKeeper
			msgServer := suite.msgServer
			id, _ := suite.PrepareCertificate(*test.inputIssuer, test.inputRecipient)
			suite.SetMembers(id, []sdk.AccAddress{*test.inputRecipient})
			if test.inputMsg != nil {
				test.inputMsg.Id = id
			}
			k.RemoveOperators(ctx, id, []sdk.AccAddress{*test.inputRecipient})

			// Reset event manager
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			// Ensure the member is an accepted member
			hasMember, _ := k.HasMember(ctx, id, *test.inputRecipient)
			suite.True(hasMember)

			// Test RejectIdentity
			res, err := msgServer.RenounceIdentity(sdk.WrapSDKContext(ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgRenounceIdentity, 0)
				// Ensure membership is unaffected by the failed call
				hasAccepted, _ := k.HasMember(ctx, id, *test.inputRecipient)
				suite.True(hasAccepted)
			} else {
				suite.NoError(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgRenounceIdentity, 1)
				suite.NotNil(res)
				// Ensure recipient is no longer an accepted member
				hasAccepted, _ := k.HasMember(ctx, id, *test.inputRecipient)
				suite.False(hasAccepted)
			}
		})
	}
}

func (suite *KeeperTestSuite) TestRenounceIdentity_InvalidCertificateId() {
	suite.SetupTest()
	k := suite.App.IdentityKeeper
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, []sdk.AccAddress{recipient})
	k.RemoveOperators(suite.Ctx, id, []sdk.AccAddress{recipient})

	invalidId := id + 1
	msg := &types.MsgRenounceIdentity{
		Creator: recipient.String(),
		Id:      invalidId,
	}
	res, err := msgServer.RenounceIdentity(sdk.WrapSDKContext(suite.Ctx), msg)
	suite.EqualError(err, types.ErrNonexistentCertificate.Wrapf("No identity found for ID %d", msg.Id).Error())
	suite.Nil(res)
}

// Sender still an operator
func (suite *KeeperTestSuite) TestRenounceIdentity_StillOperator() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, []sdk.AccAddress{recipient})

	msg := &types.MsgRenounceIdentity{
		Creator: recipient.String(),
		Id:      id,
	}
	res, err := msgServer.RenounceIdentity(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := types.ErrExistingOperator.Wrapf("address (%s) must be demoted from operator before it can be removed as a member", recipient.String()).Error()
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

// Update Members
func (suite *KeeperTestSuite) TestUpdateMembers() {
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	defaultMsg := types.MsgUpdateMembers{
		Creator:  recipient.String(),
		Id:       uint64(0), // Overwritten inside test
		ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
		ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
	}
	defaultInputMembers := []sdk.AccAddress{suite.TestAccs[8], suite.TestAccs[9], suite.TestAccs[10]}
	tests := map[string]struct {
		inputIssuer    *sdk.AccAddress
		inputRecipient *sdk.AccAddress
		inputMembers   []sdk.AccAddress
		inputMsg       *types.MsgUpdateMembers
		expErr         bool
	}{
		"nil_message": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputMsg:       nil,
			expErr:         true,
		},
		"invalid_creator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputMsg: &types.MsgUpdateMembers{
				Creator:  "creator",
				Id:       uint64(0),
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
			},
			expErr: true,
		},
		"unauthorized_sender": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputMsg: &types.MsgUpdateMembers{
				Creator:  suite.TestAccs[8].String(), // member but not operator
				Id:       uint64(0),
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
			},
			expErr: true,
		},
		"remove_operator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputMsg: &types.MsgUpdateMembers{
				Creator:  issuer.String(),
				Id:       uint64(0),
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{recipient.String()}, // fail to remove operator
			},
			expErr: true,
		},
		"update_by_operator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputMsg:       &defaultMsg,
			expErr:         false,
		},
		"update_by_issuer": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputMsg: &types.MsgUpdateMembers{
				Creator:  issuer.String(),
				Id:       uint64(0), // Overwritten inside test
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
			},
			expErr: false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			k := suite.App.IdentityKeeper
			msgServer := suite.msgServer
			ctx := suite.Ctx

			// Mock existing identity members
			id, _ := suite.PrepareCertificate(*test.inputIssuer, test.inputRecipient)
			existingMembers := append(test.inputMembers, *test.inputRecipient)
			err := suite.SetMembers(id, existingMembers)
			suite.NoError(err)

			if test.inputMsg != nil {
				test.inputMsg.Id = id

				// Ensure each toAdd is not a member and each toRemove is a member
				for _, addr := range test.inputMsg.ToAdd {
					hasMember, _ := k.HasMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.False(hasMember)
				}
				for _, addr := range test.inputMsg.ToRemove {
					hasMember, _ := k.HasMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.True(hasMember)
				}
			}

			// Reset event manager
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			// Test UpdateMembers
			res, err := msgServer.UpdateMembers(sdk.WrapSDKContext(ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgUpdateMembers, 0)
				// Ensure membership is unaffected by the failed call
				hasMember, _ := k.HasMember(ctx, id, *test.inputRecipient)
				suite.True(hasMember)
			} else {
				suite.NoError(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgUpdateMembers, 1)
				suite.NotNil(res)
				// Ensure each toRemove address has no privilages
				for _, addr := range test.inputMsg.ToRemove {
					hasMember, _ := k.HasMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.False(hasMember)
				}
				// Ensure each toAdd address is only a pending member
				for _, addr := range test.inputMsg.ToAdd {
					hasPending, _ := k.HasPendingMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.True(hasPending)
					hasAccepted, _ := k.HasMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					hasOperator, _ := k.HasOperator(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.False(hasAccepted)
					suite.False(hasOperator)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateMembers_InvalidCertificateId() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingMembers := suite.TestAccs[2:5]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, existingMembers)

	toAdd := []string{suite.TestAccs[2].String(), suite.TestAccs[3].String(), suite.TestAccs[4].String()}
	toRemove := []string{suite.TestAccs[5].String(), suite.TestAccs[6].String(), suite.TestAccs[7].String()}
	invalidId := id + 1
	msg := &types.MsgUpdateMembers{
		Creator:  recipient.String(),
		Id:       invalidId,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateMembers(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", invalidId).Error()
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestUpdateMembers_InvalidToAdd() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingMembers := suite.TestAccs[2:5]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, existingMembers)

	toAdd := []string{suite.TestAccs[2].String(), "invalid address", suite.TestAccs[4].String()}
	toRemove := []string{suite.TestAccs[5].String(), suite.TestAccs[6].String(), suite.TestAccs[7].String()}
	msg := &types.MsgUpdateMembers{
		Creator:  recipient.String(),
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateMembers(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := "decoding bech32 failed: invalid character in string: ' '"
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestUpdateMembers_InvalidToRemove() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingMembers := suite.TestAccs[2:5]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, existingMembers)

	toAdd := []string{suite.TestAccs[2].String(), suite.TestAccs[3].String(), suite.TestAccs[4].String()}
	toRemove := []string{suite.TestAccs[5].String(), "invalid address", suite.TestAccs[7].String()}
	msg := &types.MsgUpdateMembers{
		Creator:  recipient.String(),
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateMembers(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := "decoding bech32 failed: invalid character in string: ' '"
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

// Update Operators
func (suite *KeeperTestSuite) TestUpdateOperators() {
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	defaultMsg := types.MsgUpdateOperators{
		Creator:  recipient.String(),
		Id:       uint64(0), // Overwritten inside test
		ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
		ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
	}

	defaultInputMembers := []sdk.AccAddress{suite.TestAccs[2], suite.TestAccs[3]}
	defaultInputOperators := []sdk.AccAddress{suite.TestAccs[8], suite.TestAccs[9], suite.TestAccs[10]}
	tests := map[string]struct {
		inputIssuer    *sdk.AccAddress
		inputRecipient *sdk.AccAddress
		inputMembers   []sdk.AccAddress
		inputOperators []sdk.AccAddress
		inputMsg       *types.MsgUpdateOperators
		expErr         bool
	}{
		"nil_message": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputOperators: defaultInputOperators,
			inputMsg:       nil,
			expErr:         true,
		},
		"invalid_creator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputOperators: defaultInputOperators,
			inputMsg: &types.MsgUpdateOperators{
				Creator:  "creator",
				Id:       uint64(0),
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
			},
			expErr: true,
		},
		"unauthorized_sender": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputOperators: defaultInputOperators,
			inputMsg: &types.MsgUpdateOperators{
				Creator:  suite.TestAccs[12].String(), // non-operator account
				Id:       uint64(0),
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
			},
			expErr: true,
		},
		"duplicate_toAdd": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputOperators: defaultInputOperators,
			inputMsg: &types.MsgUpdateOperators{
				Creator:  issuer.String(),
				Id:       uint64(0),
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[2].String()},
				ToRemove: []string{},
			},
			expErr: true,
		},
		"update_by_operator": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputOperators: defaultInputOperators,
			inputMsg:       &defaultMsg,
			expErr:         false,
		},
		"update_by_issuer": {
			inputIssuer:    &issuer,
			inputRecipient: &recipient,
			inputMembers:   defaultInputMembers,
			inputOperators: defaultInputOperators,
			inputMsg: &types.MsgUpdateOperators{
				Creator:  issuer.String(),
				Id:       uint64(0), // Overwritten inside test
				ToAdd:    []string{suite.TestAccs[2].String(), suite.TestAccs[3].String()},
				ToRemove: []string{suite.TestAccs[8].String(), suite.TestAccs[9].String()},
			},
			expErr: false,
		},
	}

	for name, test := range tests {
		suite.Run(name, func() {
			suite.SetupTest()
			k := suite.App.IdentityKeeper
			msgServer := suite.msgServer
			ctx := suite.Ctx
			id, _ := suite.PrepareCertificate(*test.inputIssuer, test.inputRecipient)

			// Mock existing identity operators (to be demoted)
			// existingOperators := append(test.inputOperators)
			err := suite.AddOperators(id, test.inputOperators)
			suite.NoError(err)

			// Mock existing identity members (to be added)
			err = suite.SetMembers(id, test.inputMembers)
			suite.NoError(err)

			if test.inputMsg != nil {
				test.inputMsg.Id = id

				// Ensure each toAdd is not an operator and each toRemove is an operator
				for _, addr := range test.inputMsg.ToAdd {
					hasOper, _ := k.HasOperator(ctx, id, sdk.MustAccAddressFromBech32(addr))
					hasMember, _ := k.HasMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.False(hasOper)
					suite.True(hasMember)
				}
				for _, addr := range test.inputMsg.ToRemove {
					hasOper, _ := k.HasOperator(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.True(hasOper)
				}
			}

			// Reset event manager
			ctx = ctx.WithEventManager(sdk.NewEventManager())
			suite.Equal(0, len(ctx.EventManager().Events()))

			// Test UpdateOperators
			res, err := msgServer.UpdateOperators(sdk.WrapSDKContext(ctx), test.inputMsg)
			if test.expErr {
				suite.Error(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgUpdateOperators, 0)
				// Ensure operator status is unaffected by the failed call
				hasOperator, _ := k.HasOperator(ctx, id, *test.inputRecipient)
				suite.True(hasOperator)
			} else {
				suite.NoError(err)
				suite.AssertEventEmitted(ctx, types.TypeMsgUpdateOperators, 1)
				suite.NotNil(res)
				// Ensure each toRemove address is no longer an operator
				for _, addr := range test.inputMsg.ToRemove {
					hasOperator, _ := k.HasOperator(ctx, id, sdk.MustAccAddressFromBech32(addr))
					hasMember, _ := k.HasMember(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.False(hasOperator)
					suite.True(hasMember)
				}
				// Ensure each toAdd address is an operator
				for _, addr := range test.inputMsg.ToAdd {
					hasOper, _ := k.HasOperator(ctx, id, sdk.MustAccAddressFromBech32(addr))
					suite.True(hasOper)
				}
			}
		})
	}
}

func (suite *KeeperTestSuite) TestUpdateOperators_InvalidCertificateId() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingMembers := suite.TestAccs[2:5]
	existingOperators := suite.TestAccs[5:8]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, existingMembers)
	suite.AddOperators(id, existingOperators)

	toAdd := []string{suite.TestAccs[2].String(), suite.TestAccs[3].String(), suite.TestAccs[4].String()}
	toRemove := []string{suite.TestAccs[5].String(), suite.TestAccs[6].String(), suite.TestAccs[7].String()}
	invalidId := id + 1
	msg := &types.MsgUpdateOperators{
		Creator:  recipient.String(),
		Id:       invalidId,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateOperators(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := types.ErrNonexistentCertificate.Wrapf("no certificate found for ID: %d", invalidId).Error()
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestUpdateOperators_InvalidToAdd() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingMembers := suite.TestAccs[2:5]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, existingMembers) // needed to promote to operator

	toAdd := []string{suite.TestAccs[2].String(), "invalid address", suite.TestAccs[4].String()}
	toRemove := []string{}
	msg := &types.MsgUpdateOperators{
		Creator:  recipient.String(),
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateOperators(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := "decoding bech32 failed: invalid character in string: ' '"
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestUpdateOperators_InvalidToRemove() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingOperators := suite.TestAccs[5:8]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.AddOperators(id, existingOperators)

	toAdd := []string{}
	toRemove := []string{suite.TestAccs[5].String(), "invalid address", suite.TestAccs[7].String()}
	msg := &types.MsgUpdateOperators{
		Creator:  recipient.String(),
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateOperators(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := "decoding bech32 failed: invalid character in string: ' '"
	suite.EqualError(err, errorString)
	suite.Nil(res)
}

func (suite *KeeperTestSuite) TestUpdateOperators_NotAMember() {
	suite.SetupTest()
	msgServer := suite.msgServer
	issuer := suite.TestAccs[0]
	recipient := suite.TestAccs[1]
	existingMembers := suite.TestAccs[5:8]
	id, _ := suite.PrepareCertificate(issuer, &recipient)
	suite.SetMembers(id, existingMembers)

	toAdd := []string{string(suite.TestAccs[3].String())} // nonexisting member
	toRemove := []string{}
	msg := &types.MsgUpdateOperators{
		Creator:  recipient.String(),
		Id:       id,
		ToAdd:    toAdd,
		ToRemove: toRemove,
	}
	res, err := msgServer.UpdateOperators(sdk.WrapSDKContext(suite.Ctx), msg)
	errorString := sdkerrors.ErrNotFound.Wrapf("new operator is not a member of identity %d", id).Error()
	suite.EqualError(err, errorString)
	suite.Nil(res)
}
