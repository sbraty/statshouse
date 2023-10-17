// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package metajournal

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/vkcom/statshouse/internal/data_model/gen2/tlmetadata"
	"github.com/vkcom/statshouse/internal/format"
)

func newMetricStorage(loader MetricsStorageLoader) *MetricsStorage {
	result := MakeMetricsStorage("", nil, nil)
	result.journal.metaLoader = loader
	return result
}

const metricID = 123456
const namespace1ID = 332
const group1Id = 233
const id6 = 62488
const id7 = 62418

var group1 = format.MetricsGroup{
	ID:     group1Id,
	Name:   "group1",
	Weight: 1,
}

var group2Metric6 = format.MetricMetaValue{MetricID: id6, Name: group1.Name + "_metric6"}
var testMetric7 = format.MetricMetaValue{MetricID: id7, Name: "test_metric7"}

func TestMetricsStorage(t *testing.T) {
	events := []tlmetadata.Event{}
	var version int64 = 1
	incVersion := func() int64 {
		defer func() {
			version++
		}()
		return version
	}
	m := newMetricStorage(func(ctx context.Context, lastVersion int64, returnIfEmpty bool) ([]tlmetadata.Event, int64, error) {
		var result []tlmetadata.Event
		for _, e := range events {
			if e.Version > lastVersion {
				result = append(result, e)
			}
		}
		var v int64
		if len(events) > 0 {
			v = events[len(events)-1].Version
		}
		return result, v, nil
	})
	// actually 1 test, but grouped by small test case (need to run together)
	t.Run("updateJournal test(each other depends on previous)", func(t *testing.T) {
		descrField := "__description"
		metric := format.BuiltinMetrics[format.BuiltinMetricIDAPIBRS]
		metric.MetricID = metricID
		metric.Name = "test_metric"
		metric.NamespacedName = metric.Name
		metric.GroupID = 0
		_ = metric.RestoreCachedInfo()
		metricBytes, err := metric.MarshalBinary()
		require.NoError(t, err)
		dashboard := format.DashboardMeta{
			DashboardID: 94,
			Name:        "abc",
			JSONData: map[string]interface{}{
				descrField: "abc",
			},
		}
		dashboardBytes, err := json.Marshal(dashboard.JSONData)
		require.NoError(t, err)

		namespace := format.NamespaceMeta{
			ID:     332,
			Name:   "namespace",
			Weight: 1,
		}

		t.Run("create metric", func(t *testing.T) {
			metric.Version = incVersion()
			events = []tlmetadata.Event{
				{
					Id:        int64(metric.MetricID),
					Name:      metric.Name,
					EventType: format.MetricEvent,
					Version:   metric.Version,
					Data:      string(metricBytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 1)
			require.Len(t, m.journal.journalOld, 1)

			require.Equal(t, metric.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 1)
			require.Len(t, m.metricsByName, 1)
			require.Equal(t, *metric, *m.metricsByID[metric.MetricID])
			require.Equal(t, *metric, *m.metricsByName[namespacedName{
				namespace: 0,
				name:      metric.Name,
			}])
		})

		t.Run("edit metric", func(t *testing.T) {
			descr := "new descr"
			metric.Version = incVersion()
			metric.Description = descr
			bytes, err := metric.MarshalBinary()
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(metric.MetricID),
					Name:      metric.Name,
					EventType: format.MetricEvent,
					Version:   metric.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 1)
			require.Len(t, m.journal.journalOld, 1)

			require.Equal(t, metric.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 1)
			require.Len(t, m.metricsByName, 1)
			require.Equal(t, *metric, *m.metricsByID[metric.MetricID])
			require.Equal(t, *metric, *m.metricsByName[namespacedName{
				namespace: 0,
				name:      metric.Name,
			}])
		})

		t.Run("rename metric", func(t *testing.T) {
			descr := "new descr"
			metric.Name = "test_abc"
			metric.NamespacedName = metric.Name
			metric.Version = incVersion()
			metric.Description = descr
			bytes, err := metric.MarshalBinary()
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(metric.MetricID),
					Name:      metric.Name,
					EventType: format.MetricEvent,
					Version:   metric.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 1)
			require.Len(t, m.journal.journalOld, 1)

			require.Equal(t, metric.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 1)
			require.Len(t, m.metricsByName, 1)
			require.Equal(t, *metric, *m.metricsByID[metric.MetricID])
			require.Equal(t, *metric, *m.metricsByName[namespacedName{
				namespace: 0,
				name:      metric.Name,
			}])
		})

		t.Run("create dashboard", func(t *testing.T) {
			dashboard.Version = incVersion()
			events = []tlmetadata.Event{
				{
					Id:        int64(dashboard.DashboardID),
					Name:      dashboard.Name,
					EventType: format.DashboardEvent,
					Version:   dashboard.Version,
					Data:      string(dashboardBytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 2)
			require.Len(t, m.journal.journalOld, 1)

			require.Equal(t, dashboard.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 1)
			require.Len(t, m.metricsByName, 1)
			require.Equal(t, dashboard, *m.dashboardByID[dashboard.DashboardID])
		})

		t.Run("edit dashboard", func(t *testing.T) {
			descr := "new descr"
			dashboard.Version = incVersion()
			dashboard.JSONData = map[string]interface{}{
				descrField: descr,
			}
			bytes, err := json.Marshal(dashboard.JSONData)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(dashboard.DashboardID),
					Name:      dashboard.Name,
					EventType: format.DashboardEvent,
					Version:   dashboard.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 2)
			require.Len(t, m.journal.journalOld, 1)

			require.Equal(t, dashboard.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 1)
			require.Len(t, m.metricsByName, 1)
			require.Equal(t, dashboard, *m.dashboardByID[dashboard.DashboardID])
		})
		t.Run("broadcast version clients", func(t *testing.T) {
			metric.Version = incVersion()
			metric.MetricID = 1235
			metric.Name = "test1_metric3"
			events = []tlmetadata.Event{
				{
					Id:        int64(metric.MetricID),
					Name:      metric.Name,
					EventType: format.MetricEvent,
					Version:   metric.Version,
					Data:      string(metricBytes),
				},
			}
			waitCh := m.journal.waitVersion(metric.Version)
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)
			t.Run("got version update by broadcast", func(t *testing.T) {
				select {
				case <-waitCh:
				default:
					t.Fatal("Expected to get update of journal")
				}
			})
			t.Run("got version update after journal update", func(t *testing.T) {
				select {
				case <-m.journal.waitVersion(metric.Version):
				default:
					t.Fatal("Expected to get update of journal")
				}
			})
		})
		t.Run("same version and different types", func(t *testing.T) {
			dashboard.Version = incVersion()
			dashboard.DashboardID = metric.MetricID
			dashboard.JSONData = map[string]interface{}{}
			bytes, err := json.Marshal(dashboard.JSONData)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(dashboard.DashboardID),
					Name:      dashboard.Name,
					EventType: format.DashboardEvent,
					Version:   dashboard.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 4)
			require.Len(t, m.journal.journalOld, 2)

			require.Equal(t, dashboard.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 2)
			require.Len(t, m.metricsByName, 2)
			require.Len(t, m.dashboardByID, 2)
			require.Equal(t, dashboard, *m.dashboardByID[dashboard.DashboardID])
		})

		t.Run("empty journal", func(t *testing.T) {
			events = []tlmetadata.Event{}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 4)
			require.Len(t, m.journal.journalOld, 2)

			require.Equal(t, dashboard.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 2)
			require.Len(t, m.metricsByName, 2)
			require.Len(t, m.dashboardByID, 2)
			require.Equal(t, dashboard, *m.dashboardByID[dashboard.DashboardID])
		})

		t.Run("group created (check metric added to group)", func(t *testing.T) {
			group1.Version = incVersion()
			bytesGroup2, err := json.Marshal(group1)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(group1.ID),
					Name:      group1.Name,
					EventType: format.MetricsGroupEvent,
					Version:   group1.Version,
					Data:      string(bytesGroup2),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 5)

			require.Equal(t, group1.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 2)
			require.Len(t, m.metricsByName, 2)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
		})

		t.Run("metric renamed (check old metric removed from group and new metric added)", func(t *testing.T) {
			metric.Version = incVersion()
			metric.Name = group1.Name + "_metric"
			bytes, err := json.Marshal(metric)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(metric.MetricID),
					Name:      metric.Name,
					EventType: format.MetricEvent,
					Version:   metric.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 5)

			require.Equal(t, metric.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 2)
			require.Len(t, m.metricsByName, 2)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 1)
			require.Contains(t, m.metricsNamesByGroup, int32(group1Id))
			require.Contains(t, m.metricsNamesByGroup[group1Id], metric.MetricID)
		})
		t.Run("metric created (check new metric in group)", func(t *testing.T) {
			var id int32 = 65463
			metricCopy := *metric
			metricCopy.Version = incVersion()
			metricCopy.MetricID = id
			metricCopy.Name = group1.Name + "_metric5"
			bytes, err := json.Marshal(metricCopy)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(metricCopy.MetricID),
					Name:      metricCopy.Name,
					EventType: format.MetricEvent,
					Version:   metricCopy.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 6)
			require.Len(t, m.journal.journalOld, 3)

			require.Equal(t, metricCopy.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 3)
			require.Len(t, m.metricsByName, 3)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 1)
			require.Contains(t, m.metricsNamesByGroup, int32(group1Id))
			require.Len(t, m.metricsNamesByGroup[group1.ID], 2)
			require.Contains(t, m.metricsNamesByGroup[group1.ID], metricCopy.MetricID)
			require.Contains(t, m.metricsNamesByGroup[group1Id], metric.MetricID)
		})
		t.Run("group renamed (check old metric removed from group and metric added)", func(t *testing.T) {
			group1.Version = incVersion()
			group1.Name = "group3"
			bytes, err := json.Marshal(group1)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(group1.ID),
					Name:      group1.Name,
					EventType: format.MetricsGroupEvent,
					Version:   group1.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 6)
			require.Len(t, m.journal.journalOld, 3)

			require.Equal(t, group1.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 3)
			require.Len(t, m.metricsByName, 3)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
		})
		t.Run("broadcast prom clients", func(t *testing.T) {
			v1 := incVersion()
			v2 := incVersion()
			events = []tlmetadata.Event{
				{
					Id:        351525,
					Name:      "-",
					EventType: format.PromConfigEvent,
					Version:   v1,
					Data:      "abc",
				},
				{
					Id:        351525,
					Name:      "-",
					EventType: format.PromConfigEvent,
					Version:   v2,
					Data:      "def",
				},
			}
			var promConfgString string
			var promConfgVersion int64
			m.applyPromConfig = func(configString string, version int64) {
				promConfgString = configString
				promConfgVersion = version
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)
			require.Equal(t, "def", promConfgString)
			require.Equal(t, v2, promConfgVersion)
		})

		t.Run("namespace created", func(t *testing.T) {
			namespace.Version = incVersion()
			events = []tlmetadata.Event{
				{
					Id:        int64(namespace.ID),
					Name:      namespace.Name,
					EventType: format.NamespaceEvent,
					Version:   namespace.Version,
					Data:      `{"weight": 1}`,
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 8)

			require.Equal(t, namespace.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 3)
			require.Len(t, m.metricsByName, 3)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
			require.Len(t, m.namespaceByID, 1)
		})

		t.Run("metric added to namespace", func(t *testing.T) {
			metric.Version = incVersion()
			metric.NamespaceID = namespace.ID
			metricBytes, err := metric.MarshalBinary()
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(metric.MetricID),
					Name:      metric.Name,
					EventType: format.MetricEvent,
					Version:   metric.Version,
					Data:      string(metricBytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 8)

			require.Equal(t, metric.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 3)
			require.Len(t, m.metricsByName, 3)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
			require.Len(t, m.namespaceByID, 1)

			require.Contains(t, m.metricsByID, metric.MetricID)
			require.Equal(t, *m.metricsByID[metric.MetricID].Namespace, namespace)
		})

		t.Run("group added to namespace", func(t *testing.T) {
			group1.Version = incVersion()
			group1.NamespaceID = namespace.ID
			groupBytes, err := metric.MarshalBinary()
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(group1.ID),
					Name:      group1.Name,
					EventType: format.MetricsGroupEvent,
					Version:   group1.Version,
					Data:      string(groupBytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 8)

			require.Equal(t, group1.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 3)
			require.Len(t, m.metricsByName, 3)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
			require.Len(t, m.namespaceByID, 1)

			require.Contains(t, m.groupsByID, group1.ID)
			require.Equal(t, *m.groupsByID[group1.ID].Namespace, namespace)
		})
		t.Run("metric created (check new metric not in group)", func(t *testing.T) {
			group2Metric6.Version = incVersion()
			group2Metric6.NamespaceID = 0
			metric6Bytes, err := json.Marshal(group2Metric6)
			require.NoError(t, err)
			testMetric7.Version = incVersion()
			testMetric7.NamespaceID = namespace1ID
			metric7Bytes, err := json.Marshal(testMetric7)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(group2Metric6.MetricID),
					Name:      group2Metric6.Name,
					EventType: format.MetricEvent,
					Version:   group2Metric6.Version,
					Data:      string(metric6Bytes),
				},
				{
					Id:        int64(testMetric7.MetricID),
					Name:      testMetric7.Name,
					EventType: format.MetricEvent,
					Version:   testMetric7.Version,
					Data:      string(metric7Bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 10)

			require.Equal(t, testMetric7.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 5)
			require.Len(t, m.metricsByName, 5)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
		})
		t.Run("metric change namespace (check metric in group)", func(t *testing.T) {
			group2Metric6.Version = incVersion()
			group2Metric6.NamespaceID = namespace.ID
			bytes, err := json.Marshal(group2Metric6)
			require.NoError(t, err)
			events = []tlmetadata.Event{
				{
					Id:        int64(group2Metric6.MetricID),
					Name:      group2Metric6.Name,
					EventType: format.MetricEvent,
					Version:   group2Metric6.Version,
					Data:      string(bytes),
				},
			}
			err = m.journal.updateJournal(nil)
			require.NoError(t, err)

			require.Len(t, m.journal.journal, 10)

			require.Equal(t, group2Metric6.Version, m.journal.versionLocked())

			require.Len(t, m.metricsByID, 5)
			require.Len(t, m.metricsByName, 5)
			require.Len(t, m.dashboardByID, 2)
			require.Len(t, m.groupsByID, 1)
			require.Len(t, m.metricsNamesByGroup, 0)
		})
		t.Run("get metric 6", func(t *testing.T) {
			metric := m.GetMetaMetricByName(namespace.Name + namespaceSeparator + group2Metric6.Name)
			require.Equal(t, group2Metric6.MetricID, metric.MetricID)
		})
	})
	t.Run("test getJournalDiffLocked3", func(t *testing.T) {
		a := tlmetadata.Event{
			Id:      1,
			Version: 1,
		}
		b := tlmetadata.Event{
			Id:      2,
			Version: 2,
		}
		c := tlmetadata.Event{
			Id:      3,
			Version: 3,
		}
		test := func(clientVersion int64, journal, expected []tlmetadata.Event) func(t *testing.T) {
			return func(t *testing.T) {
				m.journal.journal = journal
				res := m.journal.getJournalDiffLocked3(clientVersion)
				if len(expected) == 0 {
					require.Len(t, res.Events, 0)
				} else {
					require.Equal(t, expected, res.Events)
				}
				require.Equal(t, m.journal.versionLocked(), res.CurrentVersion)
			}
		}
		t.Run("empty journal", test(0, nil, nil))
		t.Run("full journal", test(0, []tlmetadata.Event{a, b, c}, []tlmetadata.Event{a, b, c}))
		t.Run("part of journal1", test(1, []tlmetadata.Event{a, b, c}, []tlmetadata.Event{b, c}))
		t.Run("part of journal2", test(2, []tlmetadata.Event{a, b, c}, []tlmetadata.Event{c}))
		t.Run("part of journal3", test(3, []tlmetadata.Event{a, b, c}, nil))
		t.Run("part of journal4", test(999, []tlmetadata.Event{a, b, c}, nil))
	})
}
