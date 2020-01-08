import { ActionReducerMap } from '@ngrx/store';
import { routerReducer } from '@ngrx/router-store';
import { AppState } from '../state/app.state';
import { configReducers } from './config.reducers';
import { favoriteBookReducers } from './favoriteBook.reducers';

export const appReducers: ActionReducerMap<AppState, any> = {
  router: routerReducer,
  favoriteBooks: favoriteBookReducers,
  config: configReducers
};
