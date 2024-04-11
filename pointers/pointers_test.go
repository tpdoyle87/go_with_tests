package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("Want %s but got %s", want, got)
		}
	}

	assertError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("Wanted an error but didn't get one.")
		}
	}
	t.Run("depositing BTC", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdrawing BTC", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(15)}
		wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Insufficent funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err)
		assertBalance(t, wallet, startingBalance)

	})

}
