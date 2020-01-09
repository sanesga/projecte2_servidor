import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';

import {
  Book
} from '../core';

@Component({
  selector: 'app-book-page',
  styleUrls: ['book.component.css'],
  templateUrl: './book.component.html'
})
export class BookComponent implements OnInit {
  books: Book[];
  book: Book;


  constructor(
    private route: ActivatedRoute,
  ) { }

  ngOnInit() {
    //este método devuelve todos los libros
    this.route.data.subscribe(
      (data: { books: Book[] }) => {
        this.books = data.books;
      }
    );
  }
}


