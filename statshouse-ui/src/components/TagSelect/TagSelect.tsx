// Copyright 2023 V Kontakte LLC
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

import React, { memo } from 'react';
import { ReactComponent as SVGSortAlphaDown } from 'bootstrap-icons/icons/sort-alpha-down.svg';
import { ReactComponent as SVGLayers } from 'bootstrap-icons/icons/layers.svg';
import { Select, SelectOptionProps } from '../Select';
import { ToggleButton } from '../UI';

export type TagSelectProps = {
  values?: string[];
  onChange?: (value?: string | string[] | undefined, name?: string | undefined) => void;
  onFocus?: () => void;
  onBlur?: () => void;
  options?: SelectOptionProps[];
  moreOption?: boolean;
  groupBy?: boolean;
  setGroupBy?: (value: boolean) => void;
  sort?: boolean;
  setSort?: (value: boolean) => void;
  negative?: boolean;
  setNegative?: (value: boolean) => void;
  loading?: boolean;
  placeholder?: string;
};
export function _TagSelect({
  values,
  onChange,
  onBlur,
  onFocus,
  options,
  moreOption,
  sort,
  negative,
  groupBy,
  setGroupBy,
  setSort,
  setNegative,
  loading,
  placeholder,
}: TagSelectProps) {
  return (
    <>
      <Select
        options={options}
        value={values}
        multiple
        className="sh-select form-control"
        classNameList="dropdown-menu"
        loading={loading}
        placeholder={placeholder}
        showSelected={false}
        onceSelectByClick
        moreItems={moreOption}
        showCountItems
        onChange={onChange}
        onFocus={onFocus}
        onBlur={onBlur}
      />
      <ToggleButton
        title="Negate next selection"
        className="btn btn-outline-primary"
        checked={negative}
        onChange={setNegative}
      >
        −
      </ToggleButton>
      <ToggleButton title="Sort alphabetically" className="btn btn-outline-primary" checked={sort} onChange={setSort}>
        <SVGSortAlphaDown />
      </ToggleButton>
      {!!setGroupBy && (
        <ToggleButton title="Group by" className="btn btn-outline-primary" checked={groupBy} onChange={setGroupBy}>
          <SVGLayers />
        </ToggleButton>
      )}
      {/*<div className="input-group-text bg-body">
        <SwitchBox title="Group by" checked={groupBy} onChange={setGroupBy}>
          <SVGLayers />
        </SwitchBox>
      </div>*/}
    </>
  );
}

export const TagSelect = memo(_TagSelect);
