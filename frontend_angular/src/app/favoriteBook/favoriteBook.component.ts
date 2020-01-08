import { Component, OnInit } from '@angular/core';
import { Store } from '@ngrx/store';
import { FavoriteBookService, RedisService } from "../core";
import { FavoriteBook } from '../core/models';
import { AppState } from '../store/state/app.state';
import { selectfavoriteBookList } from '../store/selectors/favoriteBook.selector';
import { GetfavoriteBooks } from '../store/actions/favoriteBook.actions';
import { Observable } from 'rxjs';

@Component({
  selector: 'app-favoriteBook',
  templateUrl: './favoriteBook.component.html',
  styleUrls: ['./favoriteBook.component.css']
})
export class FavoriteBookComponent implements OnInit {

  favoriteBooks$ = this._store.select(selectfavoriteBookList);

  constructor(
    private favoriteBookService: FavoriteBookService,
    private _store: Store<AppState>,
    private redisService: RedisService,
  ) { }

  ngOnInit() {
    this._store.dispatch(new GetfavoriteBooks());
    // this.favoriteBooks$ = this.store.select(selectfavoriteBookList);
    console.log("estamos en favorite Book component");
    console.log(this.favoriteBooks$);

    //  this.favoriteBookService.getAll().subscribe((data) => {
    //   console.log(data);
    //  });
  }

}
