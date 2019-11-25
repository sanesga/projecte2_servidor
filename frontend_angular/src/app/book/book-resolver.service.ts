import { Injectable, } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, Router, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';

import { Book, BooksService, UserService } from '../core';
import { catchError } from 'rxjs/operators';

@Injectable()
export class BookResolver implements Resolve<Book> {
  constructor(
    private booksService: BooksService,
  ) {}


  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot
  ): Observable<any> { //espera una respuesta de tipo no especificado
    return this.booksService.getAll();
  }
}
