import React, { useCallback, useEffect, useMemo, useRef, useState } from 'react';
import { type PlotViewProps } from './PlotView';
import { black, grey, greyDark } from 'view/palette';
import { buildThresholdList, useIntersectionObserver, useStateToRef, useUPlotPluginHooks } from 'hooks';
import cn from 'classnames';
import { useStatsHouseShallow } from 'store2';
import { yAxisSize } from 'common/settings';
import { PlotHealsStatus } from './PlotHealsStatus';
import { PlotHeader } from './PlotHeader';
import { PlotSubMenu } from './PlotSubMenu';
import { LegendItem, UPlotPluginPortal, UPlotWrapper, UPlotWrapperPropsOpts } from 'components';
import { PlotLegend } from '../PlotLegend';
import { dateRangeFormat } from './dateRangeFormat';
import { dataIdxNearest } from 'common/dataIdxNearest';
import { font, getYAxisSize, xAxisValues, xAxisValuesCompact } from 'common/axisValues';
import { formatByMetricType, getMetricType, splitByMetricType } from 'common/formatByMetricType';
import { METRIC_TYPE } from 'api/enum';
import { xRangeStatic } from './xRangeStatic';
import { calcYRange } from 'common/calcYRange';
import css from './style.module.css';
import { PlotEventOverlay } from './PlotEventOverlay';
import { type PlotValues } from 'store2/plotDataStore';
import { useThemeStore } from 'store';
import { incrs } from './constants';

const rightPad = 16;
const threshold = buildThresholdList(1);

// const themeDark = false;
// const xAxisSize = 16;
const unFocusAlfa = 1;
const yLockDefault = { min: 0, max: 0 };
const syncGroup = '1';

export function PlotViewMetric({ className, plotKey }: PlotViewProps) {
  const {
    yLock,
    numSeries,
    error403,
    isEmbed,
    isDashboard,
    metricUnit,
    metricUnitData,
    topInfo,
    data,
    series,
    scales,
    seriesShow,
    plotWhat,
    plotDataWhat,
    legendNameWidth,
    legendValueWidth,
    legendMaxHostWidth,
    legendMaxDotSpaceWidth,
    setPlotVisibility,
    setPlotYLock,
    setTimeRange,
    setLiveMode,
    createPlotPreview,
    setPlotShow,
    resetZoom,
  } = useStatsHouseShallow(
    ({
      plotsData,
      params: { plots, tabNum },
      metricMeta,
      isEmbed,
      baseRange,
      setPlotVisibility,
      setPlotYLock,
      setTimeRange,
      setLiveMode,
      createPlotPreview,
      setPlotShow,
      resetZoom,
    }) => {
      const plot = plots[plotKey];
      const plotData = plotsData[plotKey];
      return {
        plotWhat: plot?.what,
        plotDataWhat: plotData?.whats,
        topInfo: plotData?.topInfo,
        yLock: plot?.yLock,
        numSeries: plot?.numSeries ?? 0,
        error403: plotData?.error403 ?? '',
        metricUnit: plot?.metricUnit,
        metricUnitData: plotData?.metricUnit ?? metricMeta[plot?.metricName ?? '']?.metric_type,
        data: plotData?.data,
        series: plotData?.series,
        scales: plotData?.scales,
        seriesShow: plotData?.seriesShow,
        legendNameWidth: plotData?.legendNameWidth,
        legendValueWidth: plotData?.legendValueWidth,
        legendMaxHostWidth: plotData?.legendMaxHostWidth,
        legendMaxDotSpaceWidth: plotData?.legendMaxDotSpaceWidth,
        isEmbed,
        isDashboard: +tabNum < 0,
        baseRange,
        setPlotVisibility,
        setPlotYLock,
        setTimeRange,
        setLiveMode,
        createPlotPreview,
        setPlotShow,
        resetZoom,
      };
    }
  );
  const themeDark = useThemeStore((s) => s.dark);
  const compact = isDashboard || isEmbed;
  const yLockRef = useStateToRef(yLock ?? yLockDefault);
  const getAxisStroke = useCallback(() => (themeDark ? grey : black), [themeDark]);

  const [cursorLock, setCursorLock] = useState(false);

  const uPlotRef = useRef<uPlot>();
  const [legend, setLegend] = useState<LegendItem[]>([]);

  const [pluginEventOverlay, pluginEventOverlayHooks] = useUPlotPluginHooks();

  const topPad = compact ? 8 : 16;
  const xAxisSize = compact ? 32 : 48;

  const resetZoomRef = useStateToRef(resetZoom);

  const onSetSelect = useCallback(
    (u: uPlot) => {
      if (u.status === 1) {
        const xMin = u.posToVal(u.select.left, 'x');
        const xMax = u.posToVal(u.select.left + u.select.width, 'x');
        const yMin = u.posToVal(u.select.top + u.select.height, 'y');
        const yMax = u.posToVal(u.select.top, 'y');
        const xOnly = u.select.top === 0;
        if (!xOnly) {
          setPlotYLock(plotKey, true, { min: yMin, max: yMax });
        } else {
          setLiveMode(false);
          setTimeRange({ from: Math.floor(xMin), to: Math.ceil(xMax) });
        }
      }
    },
    [setPlotYLock, plotKey, setLiveMode, setTimeRange]
  );

  const metricType = useMemo(() => {
    if (metricUnit != null) {
      return metricUnit;
    }
    return getMetricType(plotDataWhat?.length ? plotDataWhat : plotWhat, metricUnitData);
  }, [metricUnit, metricUnitData, plotDataWhat, plotWhat]);

  const opts = useMemo<UPlotWrapperPropsOpts>(() => {
    const grid: uPlot.Axis.Grid = {
      stroke: themeDark ? greyDark : grey,
      width: 1 / devicePixelRatio,
    };
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
        sync: {
          key: syncGroup,
          filters: {
            sub(event) {
              return event !== 'mouseup' && event !== 'mousedown';
            },
            pub(event) {
              return event !== 'mouseup' && event !== 'mousedown';
            },
          },
        },
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
          values: (_, splits) => splits.map(formatByMetricType(metricType)),
          size: getYAxisSize(yAxisSize),
          font: font,
          stroke: getAxisStroke,
          splits: metricType === METRIC_TYPE.none ? undefined : splitByMetricType(metricType),
          incrs,
        },
      ],
      scales: {
        x: { auto: false, range: xRangeStatic },
        y: {
          auto: (u) => !yLockRef.current || (yLockRef.current.min === 0 && yLockRef.current.max === 0),
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
          value: dateRangeFormat, //'{DD}/{MM}/{YY} {H}:{mm}:{ss}',
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
  }, [compact, getAxisStroke, metricType, pluginEventOverlay, themeDark, topPad, xAxisSize, yLockRef]);

  const onReady = useCallback(
    (u: uPlot) => {
      if (uPlotRef.current !== u) {
        uPlotRef.current = u;
      }
      // setUPlotWidth(indexPlot, u.bbox.width);
      u.over.onclick = () => {
        // @ts-ignore
        setCursorLock(u.cursor._lock);
      };
      u.over.ondblclick = () => {
        resetZoomRef.current(plotKey);
      };
      u.setCursor({ top: -10, left: -10 }, false);
    },
    [plotKey, resetZoomRef]
  );

  const onUpdatePreview = useCallback(
    (u: uPlot) => {
      createPlotPreview(plotKey, u);
    },
    [createPlotPreview, plotKey]
  );

  const [fixHeight, setFixHeight] = useState<number>(0);
  const divOut = useRef<HTMLDivElement>(null);
  const visible = useIntersectionObserver(divOut?.current, threshold, undefined, 0);
  const onMouseOver = useCallback(() => {
    if (divOut.current && !isEmbed) {
      setFixHeight(divOut.current.getBoundingClientRect().height);
    }
  }, [isEmbed]);
  const onMouseOut = useCallback(() => {
    setFixHeight(0);
  }, []);

  const onLegendFocus = useCallback((index: number, focus: boolean) => {
    if ((uPlotRef.current?.cursor as { _lock: boolean })._lock) {
      return;
    }
    uPlotRef.current?.setSeries(index, { focus }, true);
  }, []);
  //
  const onLegendShow = useCallback(
    (index: number, show: boolean, single: boolean) => {
      setPlotShow(plotKey, index - 1, show, single);
    },
    [plotKey, setPlotShow]
  );
  useEffect(() => {
    seriesShow?.forEach((show, idx) => {
      if (uPlotRef.current?.series[idx + 1] && uPlotRef.current?.series[idx + 1].show !== show) {
        uPlotRef.current?.setSeries(idx + 1, { show }, true);
      }
    });
  }, [seriesShow]);

  useEffect(() => {
    setPlotVisibility(plotKey, visible > 0);
  }, [plotKey, setPlotVisibility, visible]);

  return (
    <div
      className={cn(
        'plot-view',
        compact ? 'plot-compact' : 'plot-full',
        isDashboard && 'plot-dash',
        fixHeight > 0 && isDashboard && 'plot-hover',
        className
      )}
      ref={divOut}
      style={
        {
          '--legend-name-width': `${legendNameWidth}px`,
          '--legend-value-width': `${legendValueWidth}px`,
          '--legend-max-host-width': `${legendMaxHostWidth}px`,
          '--legend-dot-space-width': `${legendMaxDotSpaceWidth}px`,
          height: fixHeight > 0 && isDashboard ? `${fixHeight}px` : undefined,
        } as React.CSSProperties
      }
      onMouseOver={onMouseOver}
      onMouseOut={onMouseOut}
    >
      <div className="plot-view-inner">
        <div
          className="d-flex align-items-center position-relative"
          style={{
            marginRight: `${rightPad}px`,
          }}
        >
          {/*loader*/}
          <div
            style={{ width: `${yAxisSize}px` }}
            className="flex-shrink-0 d-flex justify-content-end align-items-center pe-3"
          >
            <PlotHealsStatus plotKey={plotKey} />
          </div>
          {/*header*/}
          <div className="d-flex flex-column flex-grow-1 overflow-force-wrap">
            <PlotHeader plotKey={plotKey} />
            {!compact && <PlotSubMenu plotKey={plotKey} />}
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
              className={cn('w-100 h-100 position-absolute top-0 start-0', cursorLock && css.cursorLock)}
              onUpdateLegend={setLegend}
            >
              <UPlotPluginPortal hooks={pluginEventOverlayHooks} zone="over">
                <PlotEventOverlay
                  plotKey={plotKey}
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
              plotKey={plotKey}
              legend={legend as LegendItem<PlotValues>[]}
              onLegendShow={onLegendShow}
              onLegendFocus={onLegendFocus}
              compact={compact && !(fixHeight > 0 && isDashboard)}
              unit={metricType}
            />
            {topInfo && (!compact || (fixHeight > 0 && isDashboard)) && (
              <div className="pb-3">
                Showing {numSeries > 0 ? 'top' : 'bottom'} {topInfo.top} out of {topInfo.total} total series
                {topInfo.info}
              </div>
            )}
          </div>
        )}
      </div>
    </div>
  );
}
