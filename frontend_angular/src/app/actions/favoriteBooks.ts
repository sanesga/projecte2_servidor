import { Action } from '@ngrx/store';
import { FavoriteBook } from '../core/models/favoriteBook.model';

export const FAVORITEBOOKUPDATE = '[FavoriteBook] UpdateAll';
export const FAVORITEBOOKUPDATED = '[FavoriteBook] UpdatedAll';

export class FavoriteBooksUpdateAction implements Action {
    type = FAVORITEBOOKUPDATE;
}

export class FavoriteBooksUpdatedAction implements Action {
    type = FAVORITEBOOKUPDATED;

    constructor(public payload: FavoriteBook[]) {}
}