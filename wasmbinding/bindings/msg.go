package bindings

type ArchiveMsg struct {
	// Contracts can send CDA approvals as a witness
	WitnessApproveCda *WitnessApproveCda `json:"witness_approve_cda,omitempty"`
	// Contracts can void CDAs
	VoidCda *VoidCda `json:"void_cda,omitempty"`
}

type WitnessApproveCda struct {
	// the ID of the CDA to approve
	CdaId uint64 `json:"cda_id"`
	// the on-chain data specified in MsgCreateCda
	SigningData []byte `json:"signing_data"`
}

type VoidCda struct {
	// the ID of the CDA to void
	CdaId uint64 `json:"cda_id"`
}
