import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

@Injectable()
export class RedisService {

  constructor (
    private api: ApiService
  ) {}

  save(body: Object = {}):Observable<[string]> {
      body["value"] = JSON.stringify(body["value"])
      body["key"] = JSON.stringify(body["key"])
      
    return this.api.post("/redis/", body)
    .pipe(data => {
          return data
        },);}

  getOne(key: string): Observable<[string]> {
    return this.api.get("/redis/"+key)
    .pipe(data => {
          return data
        },)}

  getAll(): Observable<[string]> {
    return this.api.get("/redis/")
    .pipe(map(data=>{
      return data
    }));   
  }
}
