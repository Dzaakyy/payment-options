package repository

import (
	"payment-options/internal/models"
	"time"
)

type paymentRepo struct{}

func NewPaymentRepo() PaymentRepository {
	return &paymentRepo{}
}

func (r *paymentRepo) CallMobileBanking() models.PaymentMethod {
	time.Sleep(2 * time.Second) // Simulate network delay
	return models.PaymentMethod{
		Account: "MBANK123xxx",
		Status:  "Active",
		Amount: "1500000",
		Icon:    "https://sampleurl.com/mobilebanking.png",
	}
}

func (r *paymentRepo) CallVirtualAccount() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "VA987654xxx",
		Status:  "Active",
		Amount: "110000",
		Icon:    "https://sampleurl.com/virtualaccount.png",
	}
}

func (r *paymentRepo) CallEWallet() models.PaymentMethod {
	// time.Sleep(2 * time.Second)
	return models.PaymentMethod{
		Account: "EWALLETxxx",
		Status:  "Active",
		Amount: "300000",
		Icon:    "https://sampleurl.com/ewallet.png",
	}
}

func (r *paymentRepo) CallInstallmentPlan() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "INST123xxx",
		Status:  "Active",
		Amount: "210000",
		Icon:    "https://sampleurl.com/installment.png",
	}
}

func (r *paymentRepo) CallGiftCard() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "GIFT789xxx",
		Status:  "Active",
		Amount: "220000",
		Icon:    "https://sampleurl.com/giftcard.png",
	}
}

func (r *paymentRepo) CallDirectDebit() models.PaymentMethod {
	return models.PaymentMethod{
		Account: "DD45678xxx",
		Status:  "Active",
		Amount: "110000",
		Icon:    "https://sampleurl.com/directdebit.png",
	}
}
