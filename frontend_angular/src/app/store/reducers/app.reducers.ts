import { ActionReducerMap } from '@ngrx/store';
import { routerReducer } from '@ngrx/router-store';
import { AppState } from '../state/app.state';
import { configReducers } from './config.reducers';
import { favoriteBookReducers } from './favoriteBook.reducers';
import { FavoriteBook } from '../../core/models/favoriteBook.model';

export const appReducers: ActionReducerMap<AppState, any> = {
  router: routerReducer,
  favoriteBooks: favoriteBookReducers,
  config: configReducers
};

export interface State {
  favoriteBook: FavoriteBook[];
}

export const selectfavoriteBookList = (state: State) => state.favoriteBook;