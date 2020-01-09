import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { FavoriteBook } from '../core/models/favoriteBook.model';
import { Store } from '@ngrx/store';
import * as fromRoot from "../reducers";
import { FavoriteBooksUpdateAction } from '../actions/favoriteBooks';

@Component({
	selector: 'app-favoriteBooks',
	templateUrl: './favoriteBooks.component.html',
	styleUrls: ['./favoriteBooks.component.css']
})

export class FavoriteBooksComponent implements OnInit {
	public favoriteBooksList: Observable<FavoriteBook[]>;

	constructor(public store: Store<fromRoot.State>) {
		this.favoriteBooksList = store.select(fromRoot.getFavoriteBooksList);
	}

	ngOnInit() {
		this.store.dispatch(new FavoriteBooksUpdateAction());
	}
}