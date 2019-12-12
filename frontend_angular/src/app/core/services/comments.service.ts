import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';

import { ApiService } from './api.service';
import { Comment } from '../models';
import { map } from 'rxjs/operators';


@Injectable()
export class CommentsService {
  constructor (
    private apiService: ApiService
  ) {}

  //PARA OBTENER LOS COMENTARIOS LA RUTA ES BOOKS Y PARA CREAR Y BORRAR ES BOOK
  add(slug, payload): Observable<Comment> {
    return this.apiService
    .post(
      `/book/${slug}/comments`,
      { comment: { body: payload } }
    ).pipe(map(data => data.comment));
  }

  getAll(slug): Observable<Comment[]> {
    return this.apiService.get(`/books/${slug}/comments`)
      .pipe(map(data => data.comments));
  }

  destroy(commentId, bookSlug) {
    return this.apiService
           .delete(`/book/${bookSlug}/comments/${commentId}`);
  }

}
