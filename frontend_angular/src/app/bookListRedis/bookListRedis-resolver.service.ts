import { Injectable, } from '@angular/core';
import { ActivatedRouteSnapshot, Resolve, RouterStateSnapshot } from '@angular/router';
import { Observable } from 'rxjs';
import { RedisService } from '../core';
import { Book } from '../core';

@Injectable()
export class BookListRedisResolver implements Resolve<Book> {
  constructor(
    private redisService: RedisService,
  ) {}


  resolve(
    route: ActivatedRouteSnapshot,
    state: RouterStateSnapshot

  ): Observable<any> { 
    //espera una respuesta de tipo no especificado
      return this.redisService.getAll()
    }
}
