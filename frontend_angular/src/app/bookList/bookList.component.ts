import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import {
  Book
} from '../core';

@Component({
  selector: 'app-bookList-page',
  templateUrl: './bookList.component.html'
})
export class BookListComponent implements OnInit {
  book: Book;

  constructor(
    private route: ActivatedRoute,
  ) { }

  ngOnInit() {
    this.route.data.subscribe(
      (data: { book: Book }) => {
        this.book = data.book;
      }
    );
  }

}
