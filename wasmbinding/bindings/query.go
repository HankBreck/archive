package bindings

type ArchiveQuery struct {
	// Contracts can fetch the signing data for a CDA
	SigningData *SigningData `json:"signing_data,omitempty"`
}

type SigningData struct {
	CdaId uint64 `json:"cda_id"`
}

type SigningDataResponse struct {
	SigningData []byte `json:"signing_data"`
}
