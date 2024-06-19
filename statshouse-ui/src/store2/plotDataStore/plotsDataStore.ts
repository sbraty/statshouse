// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import uPlot from 'uplot';
import { type MetricType, type QueryWhat } from 'api/enum';
import { type SelectOptionProps, type UPlotWrapperPropsScales } from '../../components';
import { type PlotKey, type PlotParams, type TimeRange } from 'url2';
import { type QuerySeriesMeta } from 'api/query';
import { type StoreSlice } from '../createStore';
import { type StatsHouseStore } from '../statsHouseStore';
import { loadPlotData } from './loadPlotData';
import { getClearPlotsData } from './getClearPlotsData';
import { debug } from 'common/debug';

export type PlotValues = {
  rawValue: number | null;
  value: string;
  metricName: string;
  label: string;
  baseLabel: string;
  timeShift: number;
  max_host: string;
  total: number;
  percent: string;
  max_host_percent: string;
  top_max_host: string;
  top_max_host_percent: string;
};

export type TopInfo = {
  top: string;
  total: string;
  info: string;
};

export type PlotData = {
  data: uPlot.AlignedData;
  dataView: uPlot.AlignedData;
  bands?: uPlot.Band[];
  series: uPlot.Series[];
  seriesShow: boolean[];
  scales: UPlotWrapperPropsScales;
  promQL: string;
  metricName: string;
  metricWhat: string;
  whats: QueryWhat[];
  metricUnit: MetricType;
  error: string;
  error403?: string;
  errorSkipCount: number;
  seriesTimeShift: number[];
  lastPlotParams?: PlotParams;
  lastTimeRange?: TimeRange;
  lastTimeShifts?: number[];
  lastQuerySeriesMeta?: QuerySeriesMeta[];
  receiveErrors: number;
  receiveWarnings: number;
  samplingFactorSrc: number;
  samplingFactorAgg: number;
  mappingFloodEvents: number;
  legendValueWidth: number;
  legendMaxDotSpaceWidth: number;
  legendNameWidth: number;
  legendPercentWidth: number;
  legendMaxHostWidth: number;
  legendMaxHostPercentWidth: number;
  topInfo?: TopInfo;
  maxHostLists: SelectOptionProps[][];
  promqltestfailed?: boolean;
  promqlExpand: boolean;
};

export type PlotsDataStore = {
  plotsData: Partial<Record<PlotKey, PlotData>>;
  togglePromqlExpand(plotKey: PlotKey, status?: boolean): void;
  loadPlotData(plotKey: PlotKey): void;
};
export const plotsDataStore: StoreSlice<StatsHouseStore, PlotsDataStore> = (setState, getState, store) => {
  store.subscribe((state, prevState) => {
    if (state.params !== prevState.params) {
      setState((s) => {
        getClearPlotsData(state, prevState).forEach((plotKey) => {
          delete s.plotsData[plotKey];
        });
      });
      state.viewOrderPlot.forEach((plotKey) => {
        loadPlotData(plotKey, state.params).then((updatePlotData) => {
          if (updatePlotData) {
            setState(updatePlotData);
          }
        });
      });
    }
  });
  return {
    plotsData: {},
    togglePromqlExpand(plotKey, status) {
      setState((state) => {
        const pd = state.plotsData[plotKey];
        if (pd) {
          pd.promqlExpand = status ?? !pd.promqlExpand;
        }
      });
    },
    loadPlotData(plotKey) {
      debug.log('load', plotKey);
    },
  };
};
