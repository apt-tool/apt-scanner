package scanner

import (
	"time"

	"github.com/sirupsen/logrus"

	"github.com/apt-tool/apt-scanner/internal/dnscache"
)

const (
	dnsRefreshTime   = 30 * time.Minute
	dnsLookupTimeout = 10 * time.Second
)

func NewDNSCache(logger *logrus.Logger) (*dnscache.Resolver, error) {
	dnsResolver, err := dnscache.New(dnsRefreshTime, dnsLookupTimeout, logger)
	if err != nil {
		return nil, err
	}

	return dnsResolver, nil
}
