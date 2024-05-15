import './../../testMock/matchMedia.mock';
import {
  urlEncode,
  urlEncodeGlobalParam,
  urlEncodeGroup,
  urlEncodeGroups,
  urlEncodePlot,
  urlEncodePlotFilters,
  urlEncodePlots,
  urlEncodeTimeRange,
  urlEncodeVariable,
  urlEncodeVariables,
  urlEncodeVariableSource,
} from './urlEncode';
import { GroupInfo, PlotParams, QueryParams, VariableParams } from './queryParams';
import {
  getDefaultParams,
  getNewGroup,
  getNewPlot,
  getNewVariable,
  getNewVariableSource,
  promQLMetric,
  removeValueChar,
} from './lib';
import { GET_PARAMS, METRIC_TYPE, PLOT_TYPE, QUERY_WHAT, TIME_RANGE_KEYS_TO } from '../../api/enum';
import { THEMES } from '../../store';

jest.useFakeTimers().setSystemTime(new Date('2020-01-01 00:00:00'));

const params: QueryParams = getDefaultParams();

describe('urlStore urlEncode', () => {
  test('urlEncodeVariableSource', () => {
    const dParam = {
      ...getNewVariableSource(),
      id: '0',
    };
    const prefix = GET_PARAMS.variablePrefix + '0.' + GET_PARAMS.variableSourcePrefix + dParam.id + '.';
    expect(urlEncodeVariableSource(prefix, dParam)).toEqual([[prefix + GET_PARAMS.variableSourceMetricName, '']]);
    expect(urlEncodeVariableSource(prefix, dParam, dParam)).toEqual([]);
    expect(
      urlEncodeVariableSource(
        prefix,
        { ...dParam, metric: 'm1', tag: '1', filterIn: { '0': ['v1', 'v2'] }, filterNotIn: { '1': ['n1', 'ns2'] } },
        dParam
      )
    ).toEqual([
      [prefix + GET_PARAMS.variableSourceMetricName, 'm1'],
      [prefix + GET_PARAMS.variableSourceTag, '1'],
      [prefix + GET_PARAMS.variableSourceFilter, '0-v1'],
      [prefix + GET_PARAMS.variableSourceFilter, '0-v2'],
      [prefix + GET_PARAMS.variableSourceFilter, '1~n1'],
      [prefix + GET_PARAMS.variableSourceFilter, '1~ns2'],
    ]);
  });
  test('urlEncodeVariable', () => {
    const dParam: VariableParams = {
      ...getNewVariable(),
      id: '0',
    };
    const prefix = GET_PARAMS.variablePrefix + dParam.id + '.';

    expect(urlEncodeVariable(dParam)).toEqual([[prefix + GET_PARAMS.variableName, '']]);
    expect(urlEncodeVariable(dParam, dParam)).toEqual([]);

    const dParam2: VariableParams = {
      ...dParam,
      name: 'var1',
      description: 'desc1',
      groupBy: true,
      negative: true,
      values: ['v1', 'v2'],
      source: {
        '0': {
          id: '0',
          metric: 'm1',
          tag: '1',
          filterIn: { '0': ['v1', 'v2'] },
          filterNotIn: { '1': ['n1', 'ns2'] },
        },
        '1': {
          id: '1',
          metric: 'm2',
          tag: '2',
          filterIn: {},
          filterNotIn: {},
        },
      },
      sourceOrder: ['0', '1'],
      link: [
        ['0', '0'],
        ['1', '0'],
      ],
    };
    const prefix2 = GET_PARAMS.variablePrefix + dParam2.id + '.';
    expect(urlEncodeVariable(dParam2, dParam)).toEqual([
      [prefix2 + GET_PARAMS.variableName, dParam2.name],
      [prefix2 + GET_PARAMS.variableDescription, dParam2.description],
      [prefix2 + GET_PARAMS.variableLinkPlot, '0.0-1.0'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0.' + GET_PARAMS.variableSourceMetricName, 'm1'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0.' + GET_PARAMS.variableSourceTag, '1'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0.' + GET_PARAMS.variableSourceFilter, '0-v1'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0.' + GET_PARAMS.variableSourceFilter, '0-v2'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0.' + GET_PARAMS.variableSourceFilter, '1~n1'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0.' + GET_PARAMS.variableSourceFilter, '1~ns2'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '1.' + GET_PARAMS.variableSourceMetricName, 'm2'],
      [prefix2 + GET_PARAMS.variableSourcePrefix + '1.' + GET_PARAMS.variableSourceTag, '2'],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name, 'v1'],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name, 'v2'],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name + '.' + GET_PARAMS.variableGroupBy, '1'],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name + '.' + GET_PARAMS.variableNegative, '1'],
    ]);
    expect(
      urlEncodeVariable(
        {
          ...dParam2,
          values: [],
          groupBy: false,
          negative: false,
          source: { '1': dParam2.source['1'] },
          sourceOrder: ['1'],
        },
        dParam2
      )
    ).toEqual([
      [prefix2 + GET_PARAMS.variableSourcePrefix + '0', removeValueChar],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name, removeValueChar],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name + '.' + GET_PARAMS.variableGroupBy, '0'],
      [GET_PARAMS.variableValuePrefix + '.' + dParam2.name + '.' + GET_PARAMS.variableNegative, '0'],
    ]);
  });
  test('urlEncodeVariables', () => {
    const params2: QueryParams = {
      ...params,
      variables: {
        '0': {
          id: '0',
          name: 'var1',
          description: 'desc1',
          groupBy: true,
          negative: true,
          values: ['v1', 'v2'],
          source: {
            '0': {
              id: '0',
              metric: 'm1',
              tag: '1',
              filterIn: { '0': ['v1', 'v2'] },
              filterNotIn: { '1': ['n1', 'ns2'] },
            },
            '1': {
              id: '1',
              metric: 'm2',
              tag: '2',
              filterIn: {},
              filterNotIn: {},
            },
          },
          sourceOrder: ['0', '1'],
          link: [
            ['0', '0'],
            ['1', '0'],
          ],
        },
        '1': {
          id: '1',
          name: 'var2',
          description: 'desc2',
          groupBy: true,
          negative: true,
          values: ['v21', 'v22'],
          source: {
            '0': {
              id: '0',
              metric: 'm1',
              tag: '1',
              filterIn: { '0': ['v1', 'v2'] },
              filterNotIn: { '1': ['n1', 'ns2'] },
            },
            '1': {
              id: '1',
              metric: 'm2',
              tag: '2',
              filterIn: {},
              filterNotIn: {},
            },
          },
          sourceOrder: ['0', '1'],
          link: [
            ['0', '0'],
            ['1', '0'],
          ],
        },
      },
      orderVariables: ['1', '0'],
    };
    expect(urlEncodeVariables(params)).toEqual([]);
    expect(
      urlEncodeVariables(
        {
          ...params2,
          variables: {
            '1': params2.variables['1'],
          },
          orderVariables: ['1'],
        },
        params2
      )
    ).toEqual([
      ['v0', removeValueChar],
      ['ov', '1'],
    ]);
  });
  test('urlEncodeGroup', () => {
    const dParam: GroupInfo = {
      ...getNewGroup(),
      id: '0',
    };
    expect(urlEncodeGroup(dParam)).toEqual([['g0.t', '']]);
    expect(urlEncodeGroup(dParam, dParam)).toEqual([]);
    expect(
      urlEncodeGroup({ ...dParam, show: false, size: '4', count: 4, name: 'n', description: 'd' }, dParam)
    ).toEqual([
      ['g0.t', 'n'],
      ['g0.d', 'd'],
      ['g0.n', '4'],
      ['g0.s', '4'],
      ['g0.v', '0'],
    ]);
    expect(urlEncodeGroup({ ...dParam, show: true }, { ...dParam, show: false })).toEqual([['g0.v', '1']]);
  });
  test('urlEncodeGroups', () => {
    const params2: QueryParams = {
      ...params,
      groups: {
        '0': {
          ...getNewGroup(),
          id: '0',
        },
        '1': {
          ...getNewGroup(),
          id: '1',
        },
      },
      orderGroup: ['1', '0'],
    };
    expect(urlEncodeGroups(params)).toEqual([]);
    expect(urlEncodeGroups(params2, params2)).toEqual([]);
    expect(
      urlEncodeGroups(
        {
          ...params2,
          groups: {
            '1': params2.groups[1],
          },
          orderGroup: ['1'],
        },
        params2
      )
    ).toEqual([
      ['g0', removeValueChar],
      ['og', '1'],
    ]);
  });
  test('urlEncodePlotFilters', () => {
    expect(urlEncodePlotFilters(GET_PARAMS.plotPrefix + '1.', {}, {})).toEqual([]);
    expect(urlEncodePlotFilters(GET_PARAMS.plotPrefix + '1.', { '0': ['val1'] }, { '1': ['val1'] })).toEqual([
      ['t1.qf', '0-val1'],
      ['t1.qf', '1~val1'],
    ]);
  });
  test('urlEncodePlot', () => {
    const dParam: PlotParams = {
      ...getNewPlot(),
      id: '0',
    };
    const dParam2: PlotParams = {
      ...dParam,
      metricName: 'm1',
      customName: 'cn',
      customDescription: 'cd',
      promQL: 'promql',
      metricUnit: METRIC_TYPE.microsecond,
      what: [QUERY_WHAT.count, QUERY_WHAT.maxCountHost],
      customAgg: 5,
      groupBy: ['2', '3'],
      filterIn: { '0': ['val'] },
      filterNotIn: { '1': ['noval'] },
      numSeries: 10,
      useV2: false,
      yLock: {
        min: -100,
        max: 100,
      },
      maxHost: true,
      type: PLOT_TYPE.Event,
      events: ['0', '1'],
      eventsBy: ['0', '2'],
      eventsHide: ['0', '2'],
      totalLine: true,
      filledGraph: false,
      timeShifts: [200, 300],
    };
    expect(urlEncodePlot(dParam)).toEqual([['s', '']]);
    expect(urlEncodePlot(dParam, dParam)).toEqual([]);
    expect(urlEncodePlot(dParam2, dParam)).toEqual([
      ['s', 'm1'],
      ['q', 'promql'],
      ['cn', 'cn'],
      ['cd', 'cd'],
      ['mt', 'microsecond'],
      ['qw', 'count'],
      ['qw', 'max_count_host'],
      ['g', '5'],
      ['qb', '2'],
      ['qb', '3'],
      ['qf', '0-val'],
      ['qf', '1~noval'],
      ['n', '10'],
      ['v', '1'],
      ['yl', '-100'],
      ['yh', '100'],
      ['mh', '1'],
      ['qt', '1'],
      ['qe', '0'],
      ['qe', '1'],
      ['eb', '0'],
      ['eb', '2'],
      ['eh', '0'],
      ['eh', '2'],
      ['vtl', '1'],
      ['vfg', '0'],
      ['lts', '200'],
      ['lts', '300'],
    ]);
    expect(
      urlEncodePlot(
        {
          ...dParam2,
          useV2: true,
          maxHost: false,
          totalLine: false,
          filledGraph: true,
        },
        dParam2
      )
    ).toEqual([
      ['v', '2'],
      ['mh', '0'],
      ['vtl', '0'],
      ['vfg', '1'],
    ]);
    expect(
      urlEncodePlot(
        {
          ...dParam,
          promQL: 'promql',
        },
        dParam
      )
    ).toEqual([['q', 'promql']]);
    expect(
      urlEncodePlot(
        {
          ...dParam,
          metricName: promQLMetric,
          promQL: '',
        },
        { ...dParam, promQL: 'promql' }
      )
    ).toEqual([['q', '']]);
    expect(
      urlEncodePlot(
        {
          ...dParam,
          promQL: '',
        },
        { ...dParam, promQL: 'promql' }
      )
    ).toEqual([['q', '']]);
    expect(
      urlEncodePlot({
        ...dParam,
        metricName: promQLMetric,
      })
    ).toEqual([['q', '']]);
  });
  test('urlEncodePlots', () => {
    const params2: QueryParams = {
      ...params,
      plots: {
        '0': {
          ...getNewPlot(),
          id: '0',
        },
        '1': {
          ...getNewPlot(),
          id: '1',
        },
      },
      orderPlot: ['0', '1'],
    };
    expect(urlEncodePlots(params)).toEqual([]);
    expect(urlEncodePlots(params, params)).toEqual([]);
    expect(
      urlEncodePlots(
        {
          ...params2,
          plots: {
            '1': {
              ...getNewPlot(),
              id: '1',
            },
          },
          orderPlot: ['1'],
        },
        params2
      )
    ).toEqual([
      ['t0', removeValueChar],
      ['op', '1'],
    ]);
  });
  test('urlEncodeTimeRange', () => {
    expect(urlEncodeTimeRange(params)).toEqual([]);
    expect(urlEncodeTimeRange(params, params)).toEqual([]);
    expect(
      urlEncodeTimeRange(
        {
          ...params,
          timeRange: {
            ...params.timeRange,
            from: -1000,
            urlTo: TIME_RANGE_KEYS_TO.Now,
          },
        },
        params
      )
    ).toEqual([
      ['t', '0'],
      ['f', '-1000'],
    ]);
  });
  test('urlEncodeGlobalParam', () => {
    expect(urlEncodeGlobalParam(params)).toEqual([]);
    expect(urlEncodeGlobalParam(params, params)).toEqual([]);
    expect(
      urlEncodeGlobalParam(
        {
          ...params,
          dashboardId: '1',
          dashboardName: 'dn',
          dashboardDescription: 'dd',
          dashboardVersion: 123,
          eventFrom: 2000,
          live: true,
          tabNum: '5',
          theme: THEMES.Light,
          timeShifts: [2000, 4000],
        },
        { ...params }
      )
    ).toEqual([
      ['live', '1'],
      ['theme', 'light'],
      ['tn', '5'],
      ['ts', '2000'],
      ['ts', '4000'],
      ['ef', '2000'],
      ['id', '1'],
      ['dn', 'dn'],
      ['dd', 'dd'],
      ['dv', '123'],
    ]);
    expect(
      urlEncodeGlobalParam(
        {
          ...params,
          live: false,
        },
        { ...params, live: true }
      )
    ).toEqual([['live', '0']]);
  });

  test('urlEncode', () => {
    expect(urlEncode(params)).toEqual([]);
  });
});