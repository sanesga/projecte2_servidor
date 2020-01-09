import {
  ActionReducer,
  ActionReducerMap,
  createFeatureSelector,
  createSelector,
  MetaReducer
} from '@ngrx/store';
import { environment } from '../../environments/environment';

import { FavoriteBook } from '../core/models/favoriteBook.model';
import * as fromFavoriteBooks from './favoriteBooks';

export interface State {
  favoriteBooks: FavoriteBook[];
}

export const reducers: ActionReducerMap<State> = {
  favoriteBooks: fromFavoriteBooks.reducer
};


export const metaReducers: MetaReducer<State>[] = !environment.production ? [] : [];

// Selectors
export const getFavoriteBooksList = (state: State) => state.favoriteBooks;