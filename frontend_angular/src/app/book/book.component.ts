import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';

import {
  Book
} from '../core';

@Component({
  selector: 'app-book-page',
  templateUrl: './book.component.html'
})
export class BookComponent implements OnInit {
  books: Book[];



  constructor(
    private route: ActivatedRoute,
  ) { }

  ngOnInit() {
    this.route.data.subscribe(
      (data: { books: Book[] }) => {
        this.books = data.books;
      }
    );
  }
}
