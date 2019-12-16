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

  //ESTO NO SE PUEDE HACER, TENGO QUE HACER UN COMPONENTE, QUE HACE UN FOR Y PASA POR EL INPUT EL BOOK AL BOTON
  // onToggleFavorite(favorited: boolean) {
  //   for (var book of this.books){
  //     this.book=book;
  //     this.book.favorited = favorited;

  //     if (favorited) {
  //       this.book.favoritesCount++;
  //     } else {
  //       this.book.favoritesCount--;
  //     }
  //   }
  // }

  //podemos hacer un for y hacer una petición al backend
  
}


