package config

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/Shopify/voucher"
	"github.com/Shopify/voucher/grafeas"
)

// NewMetadataClient creates a new MetadataClient.
func NewMetadataClient(ctx context.Context) (voucher.MetadataClient, error) {
	keyring, err := getKeyRing()
	if nil != err {
		log.Println("could not load keyring from ejson, continuing without attestation support: ", err)
		keyring = nil
	}

	return grafeas.NewClient(
		ctx,
		viper.GetString("image_project"),
		viper.GetString("binauth_project"),
		keyring,
	)
}
