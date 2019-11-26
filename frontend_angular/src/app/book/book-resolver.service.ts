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

  ): Observable<any> { 
    
    //espera una respuesta de tipo no especificado
    if(route.params['slug']!=null){
      console.log(route.params['slug'])
      console.log("le pasamos slug");
       return this.booksService.get(route.params['slug'])
    }else{
      console.log("no le pasamos slug")
      return this.booksService.getAll();
    }
   
   
  }

}
