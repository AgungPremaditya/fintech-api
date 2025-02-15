package ledgerentry_dtos

import "ledger-system/models"

type Mapper struct{}

func NewTransactionMapper() *Mapper {
	return &Mapper{}
}

func (m *Mapper) ToLedgerEntryResponse(ledgerEntry *models.LedgerEntry) *LedgerEntryTransferDTO {
	return &LedgerEntryTransferDTO{
		ID:       ledgerEntry.ID.String(),
		Type:     ledgerEntry.Type,
		Amount:   ledgerEntry.Amount,
		Balance:  ledgerEntry.Balance,
		WalletID: ledgerEntry.WalletID.String(),
	}
}

func (m *Mapper) ToTransactionTransferDetailResponse(debitEntry *models.LedgerEntry, creditEntry *models.LedgerEntry) *DetailTransferTransactionDTO {
	return &DetailTransferTransactionDTO{
		DebitEntry:  *m.ToLedgerEntryResponse(debitEntry),
		CreditEntry: *m.ToLedgerEntryResponse(creditEntry),
	}
}
