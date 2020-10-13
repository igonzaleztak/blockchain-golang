package clientap

// SecretStruct Struct that holds the encrypted keys that
// will be send in a transaction
type SecretStruct struct {
	AdminSecret  []byte
	ClientSecret []byte
}

// ResponseClient struct that contains the response to the client
type ResponseClient struct {
	TxSecretHash []byte
	TxDataHash   []byte
}

// ResponseSubroutine handles the responses of the goroutines
type ResponseSubroutine struct {
	txHash []byte
	Error  error
}
