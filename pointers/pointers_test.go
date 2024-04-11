package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("depositing BTC", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawing BTC", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, ErrInsufficentFunds)
		assertBalance(t, wallet, startingBalance)

	})

}

// Test helpers
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Want %s but got %s", want, got)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Wanted an error but didn't get one.")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got an error but didn't want one")
	}
}
