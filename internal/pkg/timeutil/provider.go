package timeutil

import "time"

type Provider interface {
	Now() time.Time
}

var defaultProvider Provider = NewActualProvider()

func Now() time.Time {
	return defaultProvider.Now()
}

func SetDefaultProvider(provider Provider) {
	defaultProvider = provider
}

type ActualProvider struct {
}

func (p *ActualProvider) Now() time.Time {
	return time.Now()
}

func NewActualProvider() *ActualProvider {
	return &ActualProvider{}
}

type DeterministicProvider struct {
	now time.Time
}

func (p *DeterministicProvider) Now() time.Time {
	return p.now
}

func NewDeterministicProvider(now time.Time) *DeterministicProvider {
	return &DeterministicProvider{now: now}
}
