import { RouterReducerState } from '@ngrx/router-store';
import { favoriteBookState, initialfavoriteBookState } from './favoriteBook.state';
import { initialConfigState, IConfigState } from './config.state';

export interface AppState {
  router?: RouterReducerState;
  favoriteBooks: favoriteBookState;
  config: IConfigState;
}

export const initialAppState: AppState = {
  favoriteBooks: initialfavoriteBookState,
  config: initialConfigState
};

export function getInitialState(): AppState {
  return initialAppState;
}
