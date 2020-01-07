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

  constructor(private redisService: RedisService,
    private http: HttpClient,
    private api: ApiService) { }
  favoriteBooks: FavoriteBook[]
  
  // getAll(): Observable<string[]> {
  //   return this.api.get("/redis/")
  //   .pipe((data=>{
  //     return data
  //   }));   
  // }
  getAll(): Observable<FavoriteBookHttp> {
        let params = new HttpParams();
        return this.http.get<FavoriteBookHttp>(`${environment.api_url}/redis/`, {params});
      
      }
}
