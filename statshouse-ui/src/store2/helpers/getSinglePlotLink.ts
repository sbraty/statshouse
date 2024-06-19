// Copyright 2024 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import { getDefaultParams, getNewPlot, type PlotKey, type QueryParams, urlEncode } from 'url2';
import { addPlot } from './addPlot';

export function getSinglePlotLink(plotKey: PlotKey, params: QueryParams, saveParams?: QueryParams): string {
  const nextParams = addPlot(params.plots[plotKey] ?? getNewPlot(), getDefaultParams());
  return '?' + new URLSearchParams(urlEncode(nextParams, saveParams)).toString();
}
