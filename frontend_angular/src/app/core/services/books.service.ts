import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { Book } from '../models';
import { map } from 'rxjs/operators';


@Injectable()
export class BooksService {
  constructor(
    private apiService: ApiService
  ) { }

  getAll(): Observable<[string]> {
    return this.apiService.get('/books/')
      .pipe(map(data =>{
        return data.book
      } ));
  }

  get(slug): Observable<Book> {
    return this.apiService.get('/books/' + slug)
      .pipe(map(data =>{
        //console.log(data.book)
        return data.book
      } ));
  }

  destroy(slug) {
    return this.apiService.delete('/books/' + slug);
  }

  save(book): Observable<Book> {
    // If we're updating an existing article
    if (book.slug) {
      return this.apiService.put('/books/' + book.slug, { book: book })
        .pipe(map(data => data.book));

      // Otherwise, create a new article
    } else {
      return this.apiService.post('/books/', { book: book })
        .pipe(map(data => data.book));
    }
  }
  favorite(slug): Observable<Book> {
    return this.apiService.post('/book/' + slug + '/favorite');
  }

  unfavorite(slug): Observable<Book> {
    return this.apiService.delete('/book/' + slug + '/favorite');
  }
}
