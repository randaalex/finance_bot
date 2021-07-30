package telegram

const (
	btnUpdateTransactionCategory = "btnUpdateTransactionCategory"
	btnSetTransactionCategory    = "btnSetTransactionCategory"
	btnDeleteTransaction         = "btnDeleteTransaction"
)

type Message struct {
	MessageID string
	ChatID    int64
}

func (m Message) MessageSig() (string, int64) {
	return m.MessageID, m.ChatID
}
