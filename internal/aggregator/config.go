// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package aggregator

import (
	"flag"
	"fmt"
	"strings"

	"github.com/vkcom/statshouse/internal/data_model"
)

type ConfigAggregatorRemote struct {
	InsertBudget         int // for single replica, in bytes per contributor, when many contributors
	InsertBudget100      int // for single replica, in bytes per contributor, when 100 contributors
	StringTopCountInsert int
}

type ConfigAggregator struct {
	ShortWindow        int
	RecentInserters    int
	HistoricInserters  int
	InsertHistoricWhen int

	KHAddr string

	CardinalityWindow int
	MaxCardinality    int

	ConfigAggregatorRemote

	SimulateRandomErrors float64

	MetadataNet     string
	MetadataAddr    string
	MetadataActorID uint64

	Cluster           string
	PreviousNumShards int
	ExternalPort      string

	AutoCreate bool
}

func DefaultConfigAggregator() ConfigAggregator {
	return ConfigAggregator{
		ShortWindow:          data_model.MaxShortWindow,
		RecentInserters:      4,
		HistoricInserters:    1,
		InsertHistoricWhen:   2,
		CardinalityWindow:    3600,
		MaxCardinality:       50000, // will be divided by NumShardReplicas on each aggregator
		SimulateRandomErrors: 0,
		Cluster:              "statlogs2",
		MetadataNet:          "tcp4",
		MetadataAddr:         "127.0.0.1:2442",

		ConfigAggregatorRemote: ConfigAggregatorRemote{
			InsertBudget:         400,
			InsertBudget100:      2500,
			StringTopCountInsert: 20,
		},
	}
}

func (c *ConfigAggregatorRemote) Bind(f *flag.FlagSet, d ConfigAggregatorRemote) {
	f.IntVar(&c.InsertBudget, "insert-budget", d.InsertBudget, "Aggregator will sample data before inserting into clickhouse. Bytes per contributor when # >> 100.")
	f.IntVar(&c.InsertBudget100, "insert-budget-100", d.InsertBudget100, "Aggregator will sample data before inserting into clickhouse. Bytes per contributor when # ~ 100.")
	f.IntVar(&c.StringTopCountInsert, "string-top-insert", d.StringTopCountInsert, "How many different strings per key is inserted by aggregator in string tops.")
}

func ValidateConfigAggregator(c ConfigAggregator) error {
	if c.ShortWindow > data_model.MaxShortWindow {
		return fmt.Errorf("short-window (%d) cannot be > %d", c.ShortWindow, data_model.MaxShortWindow)
	}
	if c.ShortWindow < 2 {
		return fmt.Errorf("short-window (%d) cannot be < 2", c.ShortWindow)
	}

	if c.CardinalityWindow < data_model.MinCardinalityWindow {
		return fmt.Errorf("--cardinality-window (%d) must be >= %d", c.CardinalityWindow, data_model.MinCardinalityWindow)
	}
	if c.MaxCardinality < data_model.MinMaxCardinality {
		return fmt.Errorf("--max-cardinality (%d) must be >= %d", c.MaxCardinality, data_model.MinMaxCardinality)
	}

	if c.InsertHistoricWhen < 1 {
		return fmt.Errorf("--insert-historic-when (%d) must be >= 1", c.InsertHistoricWhen)
	}
	if c.RecentInserters < 1 {
		return fmt.Errorf("--recent-inserters (%d) must be >= 1", c.RecentInserters)
	}
	if c.HistoricInserters < 1 {
		return fmt.Errorf("--historic-inserters (%d) must be >= 1", c.HistoricInserters)
	}
	if c.HistoricInserters > 4 { // Otherwise batching during historic inserts will become too small
		return fmt.Errorf("--historic-inserters (%d) must be <= 4", c.HistoricInserters)
	}

	return c.ConfigAggregatorRemote.Validate()
}

func (c *ConfigAggregatorRemote) Validate() error {
	if c.InsertBudget < 1 {
		return fmt.Errorf("insert-budget (%d) must be >= 1", c.InsertBudget)
	}
	if c.InsertBudget100 < 1 {
		return fmt.Errorf("insert-budget-100 (%d) must be >= 1", c.InsertBudget100)
	}
	if c.StringTopCountInsert < data_model.MinStringTopInsert {
		return fmt.Errorf("--string-top-insert (%d) must be >= %d", c.StringTopCountInsert, data_model.MinStringTopInsert)
	}

	return nil
}

func (c *ConfigAggregatorRemote) updateFromRemoteDescription(description string) error {
	var f flag.FlagSet
	f.Init("", flag.ContinueOnError)
	c.Bind(&f, *c)
	var s []string
	for _, v := range strings.Split(description, "\n") {
		v = strings.TrimSpace(v)
		if len(v) != 0 && !strings.HasPrefix(v, "#") {
			s = append(s, v)
		}
	}
	err := f.Parse(s)
	if err != nil {
		return err
	}
	return c.Validate()
}
