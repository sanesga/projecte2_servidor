import { createSelector } from '@ngrx/store';

import { AppState } from '../state/app.state';
import { IConfigState } from '../state/config.state';

const configState = (state: AppState) => state.config;

export const selectConfig = createSelector(
  configState,
  (state: IConfigState) => state.config
);
