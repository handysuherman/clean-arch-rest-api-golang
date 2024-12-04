package config

import "time"

type RedisKey string

const (
	REDIS RedisKey = "redis"
)

type Redis struct {
	Servers   []string       `mapstructure:"servers"`
	DB        int            `mapstructure:"db"`
	Password  string         `mapstructure:"password"`
	AppID     string         `mapstructure:"app_id"`
	PoolSize  int            `mapstructure:"pool_size"`
	EnableTLS bool           `mapstructure:"enable_tls"`
	Prefixes  *RedisPrefixes `mapstructure:"prefixes"`
}

type RedisPrefixes struct {
	CreateConsumerIdempotency *Prefixes `mapstructure:"create_consumer_idempotency"`
	UpdateConsumerIdempotency *Prefixes `mapstructure:"update_consumer_idempotency"`
	Consumer                  *Prefixes `mapstructure:"consumer"`

	CreateConsumerLoanLimitIdempotency *Prefixes `mapstructure:"create_consumer_loan_limit_idempotency"`
	UpdateConsumerLoanLimitIdempotency *Prefixes `mapstructure:"update_consumer_loan_limit_idempotency"`
	ConsumerLoanLimit                  *Prefixes `mapstructure:"consumer_loan_limit"`

	CreateConsumerTransactionIdempotency *Prefixes `mapstructure:"create_consumer_transaction_idempotency"`
	UpdateConsumerTransactionIdempotency *Prefixes `mapstructure:"update_consumer_transaction_idempotency"`
	ConsumerTransaction                  *Prefixes `mapstructure:"consumer_transaction"`

	CreateAffiliatedDealerIdempotency *Prefixes `mapstructure:"create_affiliated_dealer_idempotency"`
	UpdateAffiliatedDealerIdempotency *Prefixes `mapstructure:"update_affiliated_dealer_idempotency"`
	AffiliatedDealer                  *Prefixes `mapstructure:"affiliated_dealer"`
}

type Prefixes struct {
	Prefix             string        `mapstructure:"prefix"`
	ExpirationDuration time.Duration `mapstructure:"expiration_duration"`
}
