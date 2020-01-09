import { Injectable } from '@angular/core';
import { Effect, ofType, Actions } from '@ngrx/effects';
import { Store } from '@ngrx/store';
import { of } from 'rxjs';
import { switchMap} from 'rxjs/operators';
import { AppState } from '../state/app.state';
import {
  GetfavoriteBooksSuccess,
  EfavoriteBookActions,
  GetfavoriteBooks,
} from '../actions/favoriteBook.actions';
import { FavoriteBookService } from '../../core/services/favoriteBook.service';
import { FavoriteBookHttp } from '../../core/models/favoriteBook-http.model';

@Injectable()
export class FavoriteBookEffects {

  constructor(
    private _favoriteBookService: FavoriteBookService,
    private _actions$: Actions,
    private _store: Store<AppState>
  ) {}

  @Effect()
  getfavoriteBooks$ = this._actions$.pipe(
    ofType<GetfavoriteBooks>(EfavoriteBookActions.GetfavoriteBooks),
    switchMap(() => this._favoriteBookService.getFavoriteBooks()),
    switchMap((FavoriteBookHttp: FavoriteBookHttp) =>{
      console.log("estamos en favoriteBook effects");
      console.log(FavoriteBookHttp);
      return of(new GetfavoriteBooksSuccess(FavoriteBookHttp.favoriteBooks))
    }));
}

