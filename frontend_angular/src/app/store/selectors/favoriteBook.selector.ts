import { createSelector } from '@ngrx/store';
import { AppState } from '../state/app.state';
import { favoriteBookState } from '../state/favoriteBook.state';

const selectfavoriteBooks = (state: AppState) => state.favoriteBooks;

console.log("entra a favoriteBook.selector");

export const selectfavoriteBookList = createSelector(
  selectfavoriteBooks,
  (state: favoriteBookState) => {
    console.log("favoriteBooks: "+state.favoriteBooks);
    return state.favoriteBooks

  }
);