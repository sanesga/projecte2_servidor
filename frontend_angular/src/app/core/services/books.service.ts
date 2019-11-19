import { Injectable } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { Book, BookListConfig } from '../models';
import { map } from 'rxjs/operators';

@Injectable()
export class BooksService {
  constructor (
    private apiService: ApiService
  ) {}

  query(config: BookListConfig): Observable<{books: Book[], booksCount: number}> {
    // Convert any filters over to Angular's URLSearchParams
    const params = {};

    Object.keys(config.filters)
    .forEach((key) => {
      params[key] = config.filters[key];
    });

    return this.apiService
    .get(
      '/books/' + ((config.type === 'feed') ? 'feed' : ''),
      new HttpParams({ fromObject: params })
    );
  }

  get(slug): Observable<Book> {
    return this.apiService.get('/books/' + slug)
      .pipe(map(data => data.book));
  }

  destroy(slug) {
    return this.apiService.delete('/books/' + slug);
  }

  save(book): Observable<Book> {
    // If we're updating an existing article
    if (book.slug) {
      return this.apiService.put('/books/' + book.slug, {book: book})
        .pipe(map(data => data.book));

    // Otherwise, create a new article
    } else {
      return this.apiService.post('/books/', {article: book})
        .pipe(map(data => data.book));
    }
  }
}
