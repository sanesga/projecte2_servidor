import { Injectable } from "@angular/core";
import { Actions, Effect, ofType } from "@ngrx/effects";
import { Observable } from 'rxjs';
import { Action } from '@ngrx/store';
import * as favoriteBooks from "../actions/favoriteBooks";
import { switchMap, map } from 'rxjs/operators';
import { FavoriteBooksService } from '../core';

@Injectable()
export class FavoriteBooksEffects {
    
    @Effect()
    update: Observable<Action> = this.actions.pipe(
        ofType(favoriteBooks.FAVORITEBOOKUPDATE),
        switchMap(() =>
            this.favoriteBooksService
            .getAll()
            .pipe(map(data => new favoriteBooks.FavoriteBooksUpdatedAction(data)))
        )
    );

    constructor(
        private favoriteBooksService: FavoriteBooksService,
        private actions: Actions
    ) {}
}