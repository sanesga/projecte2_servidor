import { RouterReducerState } from '@ngrx/router-store';
import { favoriteBookState, initialfavoriteBookState } from './favoriteBook.state';

export interface AppState {
  router?: RouterReducerState;
  favoriteBooks: favoriteBookState;
}

export const initialAppState: AppState = {
  favoriteBooks: initialfavoriteBookState,
};

export function getInitialState(): AppState {
  return initialAppState;
}
