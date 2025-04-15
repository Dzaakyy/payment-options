package repository

import "payment-options/internal/models"

type PaymentRepository interface {
	CallMobileBanking() models.PaymentMethod
	CallVirtualAccount() models.PaymentMethod
	CallEWallet() models.PaymentMethod
	CallInstallmentPlan() models.PaymentMethod
	CallGiftCard() models.PaymentMethod
	CallDirectDebit() models.PaymentMethod
}
