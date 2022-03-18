package booter

import (
	"fmt"
	"os"

	eoscvault "github.com/eoscanada/eosc/vault"
	"github.com/spf13/viper"
	zsw "github.com/zhongshuwen/zswchain-go"
)

func (b *booter) newKeyBagFromVault(vaultFile string) (*zsw.KeyBag, error) {
	if _, err := os.Stat(vaultFile); err != nil {
		return nil, fmt.Errorf("vault file %q missing: %s", vaultFile, err)
	}

	vault, err := eoscvault.NewVaultFromWalletFile(vaultFile)
	if err != nil {
		return nil, fmt.Errorf("loading vault: %s", err)
	}

	boxer, err := eoscvault.SecretBoxerForType(vault.SecretBoxWrap, viper.GetString("global-kms-gcp-keypath"))
	if err != nil {
		return nil, fmt.Errorf("secret boxer: %s", err)
	}

	if err := vault.Open(boxer); err != nil {
		return nil, err
	}

	return vault.KeyBag, nil
}
