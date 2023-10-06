// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import { apiMetricListFetch, MetricShortInfo } from '../../api/metricList';
import { useErrorStore } from '../errors';
import { createStore } from '../createStore';

export type MetricsListStore = {
  list: MetricShortInfo[];
  loading: boolean;
};

export const useMetricsListStore = createStore<MetricsListStore>(
  () => ({
    list: [],
    loading: false,
  }),
  'MetricsListStore'
);

let errorRemove: (() => void) | undefined;
export async function updateMetricsList() {
  if (errorRemove) {
    errorRemove();
    errorRemove = undefined;
  }
  useMetricsListStore.setState((state) => {
    state.loading = true;
  });
  const { response, error } = await apiMetricListFetch('metricsListState');
  if (response) {
    useMetricsListStore.setState((state) => {
      state.list = response.data.metrics.map((m) => ({ name: m.name, value: m.name }));
    });
  }
  if (error) {
    errorRemove = useErrorStore.getState().addError(error);
  }
  useMetricsListStore.setState((state) => {
    state.loading = false;
  });
}

updateMetricsList();
