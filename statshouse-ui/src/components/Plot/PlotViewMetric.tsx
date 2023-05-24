// Copyright 2022 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import uPlot from 'uplot';
import { calcYRange } from '../../common/calcYRange';
import { PlotSubMenu } from './PlotSubMenu';
import { PlotHeader } from './PlotHeader';
import { LegendItem, PlotLegend, UPlotPluginPortal, UPlotWrapper, UPlotWrapperPropsOpts } from '../index';
import { formatSI, now, timeRangeAbbrevExpand } from '../../view/utils';
import { queryURLCSV } from '../../view/api';
import { black, grey, greyDark } from '../../view/palette';
import produce from 'immer';
import {
  PlotValues,
  selectorBaseRange,
  selectorLiveMode,
  selectorLoadMetricsMeta,
  selectorMetricsMetaByName,
  selectorNumQueriesPlotByIndex,
  selectorParamsPlotsByIndex,
  selectorParamsTimeShifts,
  selectorPlotLastError,
  selectorPlotsDataByIndex,
  selectorSetLiveMode,
  selectorSetParamsPlots,
  selectorSetPlotShow,
  selectorSetPreviews,
  selectorSetTimeRange,
  selectorSetUPlotWidth,
  selectorSetYLockChange,
  selectorThemeDark,
  selectorTimeRange,
  selectorUPlotsWidthByIndex,
  useStore,
} from '../../store';
import { xAxisValues, xAxisValuesCompact } from '../../common/axisValues';
import cn from 'classnames';
import { PlotEventOverlay } from './PlotEventOverlay';
import { useUPlotPluginHooks } from '../../hooks';

const unFocusAlfa = 1;
const rightPad = 16;
const font =
  '12px system-ui, -apple-system, "Segoe UI", Roboto, "Helvetica Neue", Arial, "Noto Sans", "Liberation Sans", sans-serif'; // keep in sync with $font-family-sans-serif

function xRangeStatic(u: uPlot, dataMin: number | null, dataMax: number | null): [number, number] {
  if (dataMin === null || dataMax === null) {
    const t = now();
    return [t - 3600, t];
  }
  return [dataMin, dataMax];
}

export function PlotViewMetric(props: {
  indexPlot: number;
  className?: string;
  dashboard?: boolean;
  compact: boolean;
  yAxisSize: number;
  group?: string;
}) {
  const { indexPlot, compact, yAxisSize, dashboard, className, group } = props;

  const selectorParamsPlot = useMemo(() => selectorParamsPlotsByIndex.bind(undefined, indexPlot), [indexPlot]);
  const sel = useStore(selectorParamsPlot);
  const setParamsPlots = useStore(selectorSetParamsPlots);
  const setSel = useMemo(() => setParamsPlots.bind(undefined, indexPlot), [indexPlot, setParamsPlots]);

  const timeShifts = useStore(selectorParamsTimeShifts);

  const timeRange = useStore(selectorTimeRange);
  const setTimeRange = useStore(selectorSetTimeRange);

  const baseRange = useStore(selectorBaseRange);

  const setPreviewImage = useStore(selectorSetPreviews);

  const live = useStore(selectorLiveMode);
  const setLive = useStore(selectorSetLiveMode);

  const selectorPlotsData = useMemo(() => selectorPlotsDataByIndex.bind(undefined, indexPlot), [indexPlot]);
  const {
    scales,
    series,
    seriesShow,
    data,
    legendNameWidth,
    legendValueWidth,
    legendMaxDotSpaceWidth,
    legendMaxHostWidth,
    mappingFloodEvents,
    samplingFactorSrc,
    samplingFactorAgg,
    receiveErrors,
    receiveWarnings,
    error: lastError,
    error403,
    topInfo,
  } = useStore(selectorPlotsData);

  const setPlotShow = useStore(selectorSetPlotShow);

  const setYLockChange = useStore(selectorSetYLockChange);
  const onYLockChange = useMemo(() => setYLockChange?.bind(undefined, indexPlot), [indexPlot, setYLockChange]);

  const setLastError = useStore(selectorPlotLastError);
  const selectorNumQueries = useMemo(() => selectorNumQueriesPlotByIndex.bind(undefined, indexPlot), [indexPlot]);
  const numQueries = useStore(selectorNumQueries);

  const selectorUPlotWidth = useMemo(() => selectorUPlotsWidthByIndex.bind(undefined, indexPlot), [indexPlot]);
  const width = useStore(selectorUPlotWidth);
  const setUPlotWeight = useStore(selectorSetUPlotWidth);

  const selectorPlotMetricsMeta = useMemo(
    () => selectorMetricsMetaByName.bind(undefined, sel.metricName ?? ''),
    [sel.metricName]
  );
  const meta = useStore(selectorPlotMetricsMeta);
  const loadMetricsMeta = useStore(selectorLoadMetricsMeta);

  const themeDark = useStore(selectorThemeDark);

  const uPlotRef = useRef<uPlot>();
  const [legend, setLegend] = useState<LegendItem[]>([]);

  const [pluginEventOverlay, pluginEventOverlayHooks] = useUPlotPluginHooks();

  useEffect(() => {
    if (sel.metricName) {
      loadMetricsMeta(sel.metricName);
    }
  }, [sel.metricName, loadMetricsMeta]);

  const clearLastError = useCallback(() => {
    setLastError(indexPlot, '');
  }, [indexPlot, setLastError]);

  const resetZoom = useCallback(() => {
    setSel(
      produce((s) => {
        s.yLock = { min: 0, max: 0 };
      })
    );
    setTimeRange(timeRangeAbbrevExpand(baseRange, now()));
  }, [setSel, setTimeRange, baseRange]);

  const topPad = compact ? 8 : 16;
  const xAxisSize = compact ? 32 : 48;

  // make sure "create plot" does not depend on lock range
  const yLockRef = useRef(sel.yLock);

  useEffect(() => {
    yLockRef.current = sel.yLock;
  }, [sel.yLock]);

  // all of this so that our "create plot" layout effect does not depend on the time range
  const resetZoomRef = useRef(resetZoom);
  useEffect(() => {
    resetZoomRef.current = resetZoom;
  }, [resetZoom]);

  const onSetSelect = useCallback(
    (u: uPlot) => {
      if (u.status === 1) {
        const xMin = u.posToVal(u.select.left, 'x');
        const xMax = u.posToVal(u.select.left + u.select.width, 'x');
        const yMin = u.posToVal(u.select.top + u.select.height, 'y');
        const yMax = u.posToVal(u.select.top, 'y');
        const xOnly = u.select.top === 0;
        if (!xOnly) {
          setSel(
            produce((s) => {
              s.yLock = { min: yMin, max: yMax }; // unfortunately will cause a re-scale from the effect
            })
          );
        } else {
          setLive(false);
          setTimeRange({ from: Math.floor(xMin), to: Math.ceil(xMax) });
        }
      }
    },
    [setLive, setSel, setTimeRange]
  );

  const getAxisStroke = useCallback(() => (themeDark ? grey : black), [themeDark]);

  const opts = useMemo<UPlotWrapperPropsOpts>(() => {
    const grid: uPlot.Axis.Grid = {
      stroke: themeDark ? greyDark : grey,
      width: 1 / devicePixelRatio,
    };
    const sync: uPlot.Cursor.Sync | undefined = group
      ? {
          key: group,
        }
      : undefined;
    return {
      pxAlign: false, // avoid shimmer in live mode
      padding: [topPad, rightPad, 0, 0],
      cursor: {
        lock: true,
        drag: {
          dist: 5, // try to prevent double-click-selections a bit
          x: true,
          y: true,
          uni: Infinity,
        },
        focus: {
          prox: Infinity, // always have one series focused
        },
        sync,
        dataIdx: dataIdxNearest,
      },
      focus: {
        alpha: unFocusAlfa, // avoid redrawing unfocused series
      },
      axes: [
        {
          grid: grid,
          ticks: grid,
          values: compact ? xAxisValuesCompact : xAxisValues,
          font: font,
          size: xAxisSize,
          stroke: getAxisStroke,
        },
        {
          grid: grid,
          ticks: grid,
          values: (_, splits) => splits.map(formatSI),
          size: yAxisSize,
          font: font,
          stroke: getAxisStroke,
        },
      ],
      scales: {
        x: { auto: false, range: xRangeStatic },
        y: {
          auto: false,
          range: (u: uPlot): uPlot.Range.MinMax => {
            const min = yLockRef.current.min;
            const max = yLockRef.current.max;
            if (min !== 0 || max !== 0) {
              return [min, max];
            }
            return calcYRange(u, true);
          },
        },
      },
      series: [
        {
          value: '{DD}/{MM}/{YY} {H}:{mm}:{ss}',
        },
      ],
      legend: {
        show: false,
        live: true, //!compact,
        markers: {
          width: devicePixelRatio > 1 ? 1.5 : 1,
        },
      },
      plugins: [pluginEventOverlay],
    };
  }, [compact, getAxisStroke, group, pluginEventOverlay, themeDark, topPad, xAxisSize, yAxisSize]);

  const linkCSV = useMemo(() => {
    const agg =
      sel.customAgg === -1
        ? `${Math.floor(width / 2)}`
        : sel.customAgg === 0
        ? `${Math.floor(width * devicePixelRatio)}`
        : `${sel.customAgg}s`;
    return queryURLCSV(sel, timeRange, timeShifts, agg);
  }, [sel, timeRange, timeShifts, width]);

  const onReady = useCallback(
    (u: uPlot) => {
      if (uPlotRef.current !== u) {
        uPlotRef.current = u;
      }
      setUPlotWeight(indexPlot, u.bbox.width);
      u.over.ondblclick = () => {
        resetZoomRef.current();
      };
      u.setCursor({ top: -10, left: -10 }, false);
    },
    [indexPlot, setUPlotWeight]
  );

  const onUpdatePreview = useMemo(() => setPreviewImage?.bind(undefined, indexPlot), [indexPlot, setPreviewImage]);

  const [fixHeight, setFixHeight] = useState<number>(0);
  const divOut = useRef<HTMLDivElement>(null);
  const onMouseOver = useCallback(() => {
    if (divOut.current) {
      setFixHeight(divOut.current.getBoundingClientRect().height);
    }
  }, []);
  const onMouseOut = useCallback(() => {
    setFixHeight(0);
  }, []);

  const onLegendFocus = useCallback((index: number, focus: boolean) => {
    if ((uPlotRef.current?.cursor as { _lock: boolean })._lock) {
      return;
    }
    uPlotRef.current?.setSeries(index, { focus }, true);
  }, []);

  const onLegendShow = useCallback(
    (index: number, show: boolean, single: boolean) => {
      setPlotShow(indexPlot, index - 1, show, single);
    },
    [indexPlot, setPlotShow]
  );
  useEffect(() => {
    seriesShow.forEach((show, idx) => {
      if (uPlotRef.current?.series[idx + 1].show !== show) {
        uPlotRef.current?.setSeries(idx + 1, { show }, true);
      }
    });
  }, [seriesShow]);

  return (
    <div
      className={cn(
        'plot-view',
        compact ? 'plot-compact' : 'plot-full',
        dashboard && 'plot-dash',
        fixHeight > 0 && dashboard && 'plot-hover',
        className
      )}
      ref={divOut}
      style={
        {
          '--legend-name-width': `${legendNameWidth}px`,
          '--legend-value-width': `${legendValueWidth}px`,
          '--legend-max-host-width': `${legendMaxHostWidth}px`,
          '--legend-dot-space-width': `${legendMaxDotSpaceWidth}px`,
          height: fixHeight > 0 && dashboard ? `${fixHeight}px` : undefined,
        } as React.CSSProperties
      }
      onMouseOver={onMouseOver}
      onMouseOut={onMouseOut}
    >
      <div className="plot-view-inner">
        <div className="d-flex align-items-center position-relative" style={{ marginRight: `${rightPad}px` }}>
          {/*loader*/}
          <div
            style={{ width: `${yAxisSize}px` }}
            className="flex-shrink-0 d-flex justify-content-end align-items-center pe-3"
          >
            {numQueries > 0 && (
              <div className="text-info spinner-border spinner-border-sm" role="status" aria-hidden="true" />
            )}
          </div>
          {/*header*/}
          <div className="d-flex flex-column flex-grow-1 overflow-force-wrap">
            {/*last error*/}
            {!!lastError && (
              <div className="alert alert-danger d-flex align-items-center justify-content-between" role="alert">
                <small className="overflow-force-wrap font-monospace">{lastError}</small>
                <button type="button" className="btn-close" aria-label="Close" onClick={clearLastError} />
              </div>
            )}
            <PlotHeader
              indexPlot={indexPlot}
              sel={sel}
              setParams={setSel}
              setLive={setLive}
              meta={meta}
              live={live}
              yLock={sel.yLock}
              setTimeRange={setTimeRange}
              onYLockChange={onYLockChange}
              onResetZoom={resetZoom}
              dashboard={dashboard}
              compact={compact}
            />
            {!compact && (
              /*meta*/
              <PlotSubMenu
                linkCSV={linkCSV}
                mappingFloodEvents={mappingFloodEvents}
                timeRange={timeRange}
                sel={sel}
                receiveErrors={receiveErrors}
                receiveWarnings={receiveWarnings}
                samplingFactorAgg={samplingFactorAgg}
                samplingFactorSrc={samplingFactorSrc}
              />
            )}
          </div>
        </div>
        <div
          className="position-relative w-100 z-1"
          style={
            {
              paddingTop: '61.8034%',
              '--plot-padding-top': `${topPad}px`,
            } as React.CSSProperties
          }
        >
          {error403 ? (
            <div className="text-bg-light w-100 h-100 position-absolute top-0 start-0 d-flex align-items-center justify-content-center">
              Access denied
            </div>
          ) : (
            <UPlotWrapper
              opts={opts}
              data={data}
              series={series}
              scales={scales}
              onReady={onReady}
              onSetSelect={onSetSelect}
              onUpdatePreview={onUpdatePreview}
              className="w-100 h-100 position-absolute top-0 start-0"
              onUpdateLegend={setLegend}
            >
              <UPlotPluginPortal hooks={pluginEventOverlayHooks} zone="over">
                <PlotEventOverlay
                  indexPlot={indexPlot}
                  hooks={pluginEventOverlayHooks}
                  flagHeight={Math.min(topPad, 10)}
                  compact={compact}
                />
              </UPlotPluginPortal>
            </UPlotWrapper>
          )}
        </div>
        {!error403 && (
          <div className="plot-legend">
            <PlotLegend
              indexPlot={indexPlot}
              legend={legend as LegendItem<PlotValues>[]}
              onLegendShow={onLegendShow}
              onLegendFocus={onLegendFocus}
              compact={compact && !(fixHeight > 0 && dashboard)}
            />
            {topInfo && (!compact || (fixHeight > 0 && dashboard)) && (
              <div className="pb-3">
                Showing top {topInfo.top} out of {topInfo.total} total series{topInfo.info}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}

// https://leeoniya.github.io/uPlot/demos/nearest-non-null.html
function dataIdxNearest(self: uPlot, seriesIdx: number, hoveredIdx: number, cursorXVal: number): number {
  const xValues = self.data[0];
  const yValues = self.data[seriesIdx];

  // todo: only scan in-view indices
  const mi = self.scales['x'].min;
  const ma = self.scales['x'].max;

  if (mi === undefined || ma === undefined || ma <= mi) return hoveredIdx; // not initialized, etc.
  const fromX = cursorXVal - (ma - mi) / 50; // TODO +- 2% of total area for now
  const toX = cursorXVal + (ma - mi) / 50; // TODO +- 2% of total area for now

  let nonNullLft = null;
  let nonNullRgt = null;

  for (let i = hoveredIdx; i-- > 0; ) {
    // do not include hoveredIdx
    if (xValues[i] < fromX) {
      break;
    }
    if (yValues[i] != null) {
      nonNullLft = i;
      break;
    }
    // Optimization, uncomment if 'if' above does not work well
    // for (let j = 1; j < self.data.length; ++j)
    //   if (self.data[j][i] != null) {
    //     nonNullLft = i;
    //     break;
    //   }
    // if (nonNullLft != null)
    //   break;
  }

  for (let i = hoveredIdx; i < yValues.length; i++) {
    // include hoveredIdx
    if (xValues[i] > toX) {
      break;
    }
    if (yValues[i] != null) {
      nonNullRgt = i;
      break;
    }
    // Optimization, uncomment if 'if' above does not work well
    // for (let j = 1; j < self.data.length; ++j)
    //   if (self.data[j][i] != null) {
    //     nonNullRgt = i;
    //     break;
    //   }
    // if (nonNullRgt != null)
    //   break;
  }

  const rgtVal = nonNullRgt == null ? Infinity : xValues[nonNullRgt];
  const lftVal = nonNullLft == null ? -Infinity : xValues[nonNullLft];

  const lftDelta = cursorXVal - lftVal;
  const rgtDelta = rgtVal - cursorXVal;

  const idx = lftDelta <= rgtDelta ? nonNullLft : nonNullRgt;
  if (idx !== null) return idx;
  // this code path includes returning index where only timestamps are set, but no related data points
  return hoveredIdx;
}
