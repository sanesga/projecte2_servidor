import { Injectable } from '@angular/core';
import { HttpClient, HttpParams } from '@angular/common/http';
import { RedisService } from './redis.service';
import { ApiService } from './api.service';
import { FavoriteBook } from '../models/favoriteBook.model';
import { Observable } from 'rxjs';
import { FavoriteBookHttp } from '../models/favoriteBook-http.model';
import { environment } from '../../../environments/environment';

@Injectable()
export class FavoriteBookService {

  favoriteBooksUrl = `${environment.api_url}/redis/`;

  constructor(
    private redisService: RedisService,
    private _http: HttpClient,
    private api: ApiService) { }

  
  // getAll(): Observable<FavoriteBookHttp> {
  //   return this.api.get("/redis/")
  //   .pipe((data=>{
  //     return data
  //   }));   
  // }
  getFavoriteBooks(): Observable<FavoriteBookHttp> {
      return this._http.get<FavoriteBookHttp>(this.favoriteBooksUrl);
      }
}
