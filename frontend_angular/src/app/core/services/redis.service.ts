import { Injectable } from '@angular/core';
import { ApiService } from './api.service';

@Injectable()
export class RedisService {

  constructor (
    private api: ApiService
  ) {}

  save(body: Object = {}) {
      body["value"] = JSON.stringify(body["value"])
      body["key"] = JSON.stringify(body["key"])
      
    return this.api.post("/redis/", body).subscribe(
        data => {
          return data
        },
        err => console.log(err)
      );
  }

  getOne(key: string) {
    return this.api.get("/redis/"+key).subscribe(
        data => {
          return data
        },
        err => console.log(err)
    )
  }

  getAll() {
    return this.api.get("/redis/").subscribe(
        data => {
          console.log(data)
          return data
        },
        err => console.log(err)
    )
  }
}
