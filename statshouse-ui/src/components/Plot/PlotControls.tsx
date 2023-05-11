// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import React, { ChangeEvent, memo, useCallback, useEffect, useMemo } from 'react';
import * as utils from '../../view/utils';
import { getTimeShifts, promQLMetric, timeShiftAbbrevExpand } from '../../view/utils';
import { MetricItem } from '../../hooks';
import { PlotControlFrom, PlotControlTimeShifts, PlotControlTo, Select, SelectOptionProps } from '../index';
import { TagControl } from '../../view/TagControl';
import { ReactComponent as SVGFiles } from 'bootstrap-icons/icons/files.svg';
import { ReactComponent as SVGLightning } from 'bootstrap-icons/icons/lightning.svg';
import { ReactComponent as SVGPcDisplay } from 'bootstrap-icons/icons/pc-display.svg';
import { ReactComponent as SVGCode } from 'bootstrap-icons/icons/code.svg';
import {
  selectorLastError,
  selectorParams,
  selectorParamsTagSync,
  selectorParamsTimeShifts,
  selectorPlotsData,
  selectorPlotsDataByIndex,
  selectorSetLastError,
  selectorSetParams,
  selectorSetTimeRange,
  selectorTimeRange,
  useStore,
} from '../../store';
import { globalSettings } from '../../common/settings';
import { filterHasTagID, metricKindToWhat, metricMeta, queryWhat, whatToWhatDesc } from '../../view/api';
import produce from 'immer';
import { PLOT_TYPE, PlotParams } from '../../common/plotQueryParams';
import cn from 'classnames';

export const PlotControls = memo(function PlotControls_(props: {
  indexPlot: number;
  setBaseRange: (r: utils.timeRangeAbbrev) => void;
  sel: PlotParams;
  setSel: (state: React.SetStateAction<PlotParams>, replaceUrl?: boolean) => void;
  meta: metricMeta;
  numQueries: number;
  metricsOptions: MetricItem[];
  clonePlot?: () => void;
}) {
  const { indexPlot, setBaseRange, sel, setSel, meta, numQueries, clonePlot, metricsOptions } = props;

  const timeShifts = useStore(selectorParamsTimeShifts);
  const params = useStore(selectorParams);
  const setParams = useStore(selectorSetParams);

  const timeRange = useStore(selectorTimeRange);
  const setTimeRange = useStore(selectorSetTimeRange);

  const syncTags = useStore(selectorParamsTagSync);

  const lastError = useStore(selectorLastError);
  const setLastError = useStore(selectorSetLastError);

  const selectorPlotData = useMemo(() => selectorPlotsDataByIndex.bind(undefined, indexPlot), [indexPlot]);
  const plotData = useStore(selectorPlotData);
  const plotsData = useStore(selectorPlotsData);

  const clearLastError = useCallback(() => {
    setLastError('');
  }, [setLastError]);

  const eventPlotList = useMemo<SelectOptionProps[]>(() => {
    const eventPlots: SelectOptionProps[] = params.plots
      .map((p, indexP) => [p, indexP] as [PlotParams, number])
      .filter(([p]) => p.type === PLOT_TYPE.Event && p.metricName !== '')
      .map(([p, indexP]) => {
        const metricName =
          p.customName || (p.metricName !== promQLMetric ? p.metricName : plotsData[indexP].nameMetric);
        const what =
          p.metricName === promQLMetric
            ? plotsData[indexP].whats.map((qw) => whatToWhatDesc(qw)).join(', ')
            : p.what.map((qw) => whatToWhatDesc(qw)).join(', ');
        const name = metricName + (what ? ': ' + what : '');
        return {
          value: indexP.toString(),
          name,
        };
      });
    return eventPlots;
  }, [params.plots, plotsData]);

  // keep meta up-to-date when sel.metricName changes (e.g. because of navigation)
  useEffect(() => {
    const whats = metricKindToWhat(meta.kind);
    if (meta.name === sel.metricName && sel.what.some((qw) => whats.indexOf(qw) === -1)) {
      // console.log('reset what', meta, sel.metricName, sel.what, whats);
      setSel(
        (s) => ({
          ...s,
          what: [whats[0]],
        }),
        true
      );
    }
  }, [meta.kind, meta.name, sel.metricName, sel.what, setSel]);

  const onCustomAggChange = useCallback(
    (e: ChangeEvent<HTMLSelectElement>) => {
      const customAgg = parseInt(e.target.value);
      const timeShiftsSet = getTimeShifts(customAgg);
      const shifts = timeShifts.filter(
        (v) => timeShiftsSet.find((shift) => timeShiftAbbrevExpand(shift) === v) !== undefined
      );
      setParams((p) => ({ ...p, timeShifts: shifts }));
      setSel((s) => ({
        ...s,
        customAgg: customAgg,
      }));
    },
    [setParams, setSel, timeShifts]
  );

  const onNumSeriesChange = useCallback(
    (e: ChangeEvent<HTMLSelectElement>) => {
      setSel((s) => ({
        ...s,
        numSeries: parseInt(e.target.value),
      }));
    },
    [setSel]
  );

  const onV2Change = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      setSel((s) => ({
        ...s,
        useV2: e.target.checked,
      }));
    },
    [setSel]
  );
  const onHostChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      setSel((s) => ({
        ...s,
        maxHost: e.target.checked,
      }));
    },
    [setSel]
  );

  const onMetricChange = useCallback(
    (value?: string | string[]) => {
      if (typeof value !== 'string') {
        return;
      }
      setSel((s) => {
        const whats = metricKindToWhat(meta.kind);
        return {
          ...s,
          metricName: value,
          customName: '',
          what: s.what.some((qw) => whats.indexOf(qw) === -1) ? [whats[0]] : s.what,
          groupBy: [],
          filterIn: {},
          filterNotIn: {},
        };
      });
    },
    [meta.kind, setSel]
  );

  const onWhatChange = useCallback(
    (value?: string | string[]) => {
      const whatValue = Array.isArray(value) ? value : value ? [value] : [];
      if (!whatValue.length) {
        const whats = metricKindToWhat(meta.kind);
        whatValue.push(whats[0]);
      }
      setSel((s) => ({
        ...s,
        what: whatValue as queryWhat[],
      }));
    },
    [meta.kind, setSel]
  );

  const whatOption = useMemo(
    () =>
      metricKindToWhat(meta.kind).map((w) => ({
        value: w,
        name: whatToWhatDesc(w),
        disabled: w === '-',
        splitter: w === '-',
      })),
    [meta.kind]
  );

  const toPromql = useCallback(() => {
    setSel(
      produce((s) => {
        const whats = metricKindToWhat(meta.kind);

        s.metricName = promQLMetric;
        s.what = [whats[0]];
        s.groupBy = [];
        s.filterIn = {};
        s.filterNotIn = {};
        s.promQL = plotData.promQL;
      })
    );
  }, [meta.kind, plotData.promQL, setSel]);

  const eventsChange = useCallback(
    (value: string | string[] = []) => {
      setSel(
        produce((s) => {
          s.events = Array.isArray(value) ? value.map((v) => parseInt(v)) : [parseInt(value)];
        })
      );
    },
    [setSel]
  );

  return (
    <div>
      <div
        className="alert alert-danger d-flex align-items-center justify-content-between"
        hidden={lastError === ''}
        role="alert"
      >
        <small className="overflow-force-wrap font-monospace">{lastError}</small>
        <button type="button" className="btn-close" aria-label="Close" onClick={clearLastError} />
      </div>

      <form spellCheck="false">
        <div className="d-flex mb-2">
          <div className="col input-group">
            <Select
              value={sel.metricName}
              options={metricsOptions}
              onChange={onMetricChange}
              valueToInput={true}
              className="sh-select form-control"
              classNameList="dropdown-menu"
            />
            {!!clonePlot && (
              <button
                type="button"
                onClick={clonePlot}
                className="btn btn-outline-primary"
                title="Duplicate plot to new tab"
              >
                <SVGFiles />
              </button>
            )}
          </div>
          {sel.type === PLOT_TYPE.Metric && (
            <button type="button" className="btn btn-outline-primary ms-3" onClick={toPromql} title="PromQL">
              <SVGCode />
            </button>
          )}
        </div>
        <div className="row mb-3">
          <div className="d-flex align-items-baseline">
            <Select
              value={sel.what}
              onChange={onWhatChange}
              options={whatOption}
              multiple
              onceSelectByClick
              className={cn('sh-select form-control', sel.type === PLOT_TYPE.Metric && ' me-4')}
              classNameList="dropdown-menu"
            />
            {sel.type === PLOT_TYPE.Metric && (
              <div className="form-check form-switch">
                <input
                  className="form-check-input"
                  type="checkbox"
                  value=""
                  id="switchMaxHost"
                  checked={sel.maxHost}
                  onChange={onHostChange}
                />
                <label className="form-check-label" htmlFor="switchMaxHost" title="Host">
                  <SVGPcDisplay />
                </label>
              </div>
            )}
          </div>
        </div>

        <div className="row mb-3 align-items-baseline">
          <PlotControlFrom timeRange={timeRange} setTimeRange={setTimeRange} setBaseRange={setBaseRange} />
          <div className="align-items-baseline mt-2">
            <PlotControlTo timeRange={timeRange} setTimeRange={setTimeRange} />
          </div>
          <PlotControlTimeShifts className="w-100 mt-2" />
        </div>

        <div className="row mb-3 align-items-baseline">
          <div className={cn(sel.type === PLOT_TYPE.Metric ? 'col-4' : 'col-8')}>
            <select
              className={`form-select ${sel.customAgg > 0 ? 'border-warning' : ''}`}
              value={sel.customAgg}
              onChange={onCustomAggChange}
            >
              <option value={0}>Auto</option>
              <option value={-1}>Auto (low)</option>
              <option value={1}>1 second</option>
              <option value={5}>5 seconds</option>
              <option value={15}>15 seconds</option>
              <option value={60}>1 minute</option>
              <option value={5 * 60}>5 minutes</option>
              <option value={15 * 60}>15 minutes</option>
              <option value={60 * 60}>1 hour</option>
              <option value={4 * 60 * 60}>4 hours</option>
              <option value={24 * 60 * 60}>24 hours</option>
              <option value={7 * 24 * 60 * 60}>7 days</option>
              <option value={31 * 24 * 60 * 60}>1 month</option>
            </select>
          </div>
          {sel.type === PLOT_TYPE.Metric && (
            <div className="col-4">
              <select className="form-select" value={sel.numSeries} onChange={onNumSeriesChange}>
                <option value="1">Top 1</option>
                <option value="2">Top 2</option>
                <option value="3">Top 3</option>
                <option value="4">Top 4</option>
                <option value="5">Top 5</option>
                <option value="10">Top 10</option>
                <option value="20">Top 20</option>
                <option value="30">Top 30</option>
                <option value="40">Top 40</option>
                <option value="50">Top 50</option>
                <option value="100">Top 100</option>
                <option value="-1">Bottom 1</option>
                <option value="-2">Bottom 2</option>
                <option value="-3">Bottom 3</option>
                <option value="-4">Bottom 4</option>
                <option value="-5">Bottom 5</option>
                <option value="-10">Bottom 10</option>
                <option value="-20">Bottom 20</option>
                <option value="-30">Bottom 30</option>
                <option value="-40">Bottom 40</option>
                <option value="-50">Bottom 50</option>
                <option value="-100">Bottom 100</option>
              </select>
            </div>
          )}
          <div className="col-4 d-flex justify-content-end">
            <div className="form-check form-switch">
              <input
                className="form-check-input"
                type="checkbox"
                value=""
                id="useV2Input"
                checked={sel.useV2}
                disabled={globalSettings.disabled_v1}
                onChange={onV2Change}
              />
              <label className="form-check-label" htmlFor="useV2Input" title="Use StatsHouse v2">
                <SVGLightning />
              </label>
            </div>
          </div>
        </div>

        {numQueries !== 0 && (
          <div className="text-center">
            <div className="text-info spinner-border spinner-border-sm m-5" role="status" aria-hidden="true" />
          </div>
        )}

        {numQueries === 0 &&
          (meta.tags || []).map((t, index) =>
            t.description === '-' && !filterHasTagID(sel, index) ? null : (
              <TagControl
                key={`${meta.name} tag${index}`}
                tag={t}
                indexTag={index}
                indexPlot={indexPlot}
                tagID={`key${index}`}
                sync={syncTags.some((g) => g[indexPlot] === index)}
              />
            )
          )}

        {numQueries === 0 && (meta.string_top_name || meta.string_top_description || filterHasTagID(sel, -1)) && (
          <TagControl
            key={`${meta.name} tag_s`}
            tag={{
              name: meta.string_top_name ? meta.string_top_name : 'tag_s',
              description: meta.string_top_description,
            }}
            indexTag={-1}
            indexPlot={indexPlot}
            tagID={`skey`}
          />
        )}
        {sel.type === PLOT_TYPE.Metric && !!eventPlotList.length && (
          <div>
            <Select
              value={sel.events.map((e) => e.toString())}
              onChange={eventsChange}
              className="sh-select form-control"
              classNameList="dropdown-menu"
              showSelected={true}
              onceSelectByClick
              multiple
              options={eventPlotList}
            />
          </div>
        )}
      </form>
    </div>
  );
});
