package usecase

import (
	"payment-options/internal/models"
	"payment-options/internal/repository"
	// "sync"
)

type paymentUsecase struct {
	repo repository.PaymentRepository
}

func NewPaymentUsecase(r repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: r}
}

func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
		result := make(map[string]models.PaymentMethod)
		result["mobilebanking"] = u.repo.CallMobileBanking()
		result["virtualaccount"] = u.repo.CallVirtualAccount()
		result["ewallet"] = u.repo.CallEWallet()
		result["installmentplan"] = u.repo.CallInstallmentPlan()
		result["giftcard"] = u.repo.CallGiftCard()
		result["directdebit"] = u.repo.CallDirectDebit()
		return result, nil
}

// func (u *paymentUsecase) GetPaymentOptions() (map[string]models.PaymentMethod, error) {
// 	var wg sync.WaitGroup
// 	result := make(map[string]models.PaymentMethod)
// 	mu := sync.Mutex{}

// 	wg.Add(6)

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		result["mobilebanking"] = u.repo.CallMobileBanking()
// 		mu.Unlock()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		result["virtualaccount"] = u.repo.CallVirtualAccount()
// 		mu.Unlock()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		result["ewallet"] = u.repo.CallEWallet()
// 		mu.Unlock()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		result["installmentplan"] = u.repo.CallInstallmentPlan()
// 		mu.Unlock()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		result["giftcard"] = u.repo.CallGiftCard()
// 		mu.Unlock()
// 	}()

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		result["directdebit"] = u.repo.CallDirectDebit()
// 		mu.Unlock()
// 	}()

// 	wg.Wait()
// 	return result, nil
// }
