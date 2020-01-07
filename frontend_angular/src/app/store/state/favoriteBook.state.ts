import { FavoriteBook } from '../../core/models/favoriteBook.model';

export interface favoriteBookState {
  favoriteBooks: FavoriteBook[];
}

export const initialfavoriteBookState: favoriteBookState = {
  favoriteBooks: null,
};
