import React, { memo, useCallback, useMemo } from 'react';
import { type PlotKey } from 'url2';
import { Button, Tooltip } from 'components';
import { ReactComponent as SVGExclamationTriangleFill } from 'bootstrap-icons/icons/exclamation-triangle-fill.svg';
import { ReactComponent as SVGArrowCounterclockwise } from 'bootstrap-icons/icons/arrow-counterclockwise.svg';
import { useStatsHouseShallow } from '../../../store2';

export type PlotHealsStatusProps = {
  className?: string;
  plotKey: PlotKey;
};
export function _PlotHealsStatus({ className, plotKey }: PlotHealsStatusProps) {
  const { lastError, plotHealsTimeout, interval, numQueries, clearPlotError, loadPlotData } = useStatsHouseShallow(
    ({ plotsData, plotHeals, liveMode, clearPlotError, loadPlotData }) => ({
      lastError: plotsData[plotKey]?.error,
      numQueries: plotsData[plotKey]?.numQueries ?? 0,
      plotHealsTimeout: plotHeals[plotKey]?.timeout,
      interval: liveMode.interval,
      clearPlotError,
      loadPlotData,
    })
  );
  const healsInfo = useMemo(() => {
    if (plotHealsTimeout && interval < plotHealsTimeout) {
      return `plot update timeout ${plotHealsTimeout} sec`;
    }
    return undefined;
  }, [interval, plotHealsTimeout]);
  const clearLastError = useCallback(() => {
    clearPlotError(plotKey);
  }, [clearPlotError, plotKey]);
  const reload = useCallback(() => {
    clearPlotError(plotKey);
    loadPlotData(plotKey);
  }, [clearPlotError, loadPlotData, plotKey]);
  return (
    <Tooltip
      titleClassName="alert alert-danger p-0"
      horizontal="left"
      vertical="out-bottom"
      hover
      style={{ width: 24, height: 24 }}
      open={lastError ? undefined : false}
      title={
        !!lastError && (
          <div className="d-flex flex-nowrap align-items-center justify-content-between" role="alert">
            <Button type="button" className="btn" aria-label="Reload" onClick={reload}>
              <SVGArrowCounterclockwise />
            </Button>
            <div>
              <pre className="my-0 mx-1 overflow-force-wrap font-monospace">{lastError}</pre>
              {!!healsInfo && (
                <pre className="my-0 mx-1 overflow-force-wrap font-monospace text-secondary">{healsInfo}</pre>
              )}
            </div>
            <Button type="button" className="btn btn-close" aria-label="Close" onClick={clearLastError} />
          </div>
        )
      }
    >
      {numQueries > 0 ? (
        <div className="text-info spinner-border spinner-border-sm" role="status" aria-hidden="true" />
      ) : !!lastError ? (
        <div>
          <SVGExclamationTriangleFill className="text-danger" />
        </div>
      ) : null}
    </Tooltip>
  );
}
export const PlotHealsStatus = memo(_PlotHealsStatus);
