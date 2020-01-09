import { Action } from '@ngrx/store';
import { FavoriteBook } from '../../core/models/favoriteBook.model';

export enum EfavoriteBookActions {
  GetfavoriteBooks = '[favoriteBook] Get favoriteBooks',
  GetfavoriteBooksSuccess = '[favoriteBook] Get favoriteBooks Success',
}
export class GetfavoriteBooks implements Action {
    public readonly type = EfavoriteBookActions.GetfavoriteBooks
}

export class GetfavoriteBooksSuccess implements Action {
public readonly type = EfavoriteBookActions.GetfavoriteBooksSuccess;
  constructor(public payload: FavoriteBook[]) {}
}

export type favoriteBookActions = GetfavoriteBooks | GetfavoriteBooksSuccess;