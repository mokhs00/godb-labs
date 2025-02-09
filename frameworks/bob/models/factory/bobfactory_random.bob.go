// Code generated by BobGen mysql v0.30.0. DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package factory

import (
	"strings"
	"time"

	"github.com/jaswdr/faker/v2"
)

var defaultFaker = faker.New()

func random_int64(f *faker.Faker) int64 {
	if f == nil {
		f = &defaultFaker
	}

	return f.Int64()
}

func random_string(f *faker.Faker) string {
	if f == nil {
		f = &defaultFaker
	}

	return strings.Join(f.Lorem().Words(f.IntBetween(1, 5)), " ")
}

func random_time_Time(f *faker.Faker) time.Time {
	if f == nil {
		f = &defaultFaker
	}

	year := time.Hour * 24 * 365
	min := time.Now().Add(-year)
	max := time.Now().Add(year)
	return f.Time().TimeBetween(min, max)
}
