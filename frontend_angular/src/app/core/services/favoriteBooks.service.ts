import { Injectable } from '@angular/core';
import { ApiService } from './api.service';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { FavoriteBook } from '../models/favoriteBook.model';

@Injectable()
export class FavoriteBooksService {
	favoriteBooks: Array<FavoriteBook> = [];
	favoriteBook: FavoriteBook;

	constructor(
		private apiService: ApiService,
	) { }

	getAll(): Observable<FavoriteBook[]> {
		return this.apiService.get('/redis/')
			.pipe(map(data => {
				console.log(data.keys);
				this.favoriteBooks = [];

				for (const key in data.keys) {
					//imprime la clave
					//console.log(key)
					//imprime el valor
					//console.log(data.keys[key]);

					//pasamos el número de favoritos, de string a number
					var likes = parseInt(data.keys[key]);

					//si el libro tiene más de 5 likes, se mostrará en el home, si no, no.
					if (likes >= 5) { //SI EN EL HOME NO SE IMPRIMIE NADA ES QUE NO HAY LIBROS CON MÁS DE 5 LIKES

						//creamos la propiedad título, sustituyendo los guiones por espacios para imprimir en el home
						var title = new String();

						for (var j = 0; j < key.length; j++) {
							title = key.split("-").join(" ");
						}
						//creamos un objeto de tipo favoriteBook con la key y el value recuperados
						this.favoriteBook = {
							key: title,
							value: data.keys[key]
						}
						//guardamos los objetos en el array de libros favoritos
						this.favoriteBooks.push(this.favoriteBook);
					}
			
				}
				return this.favoriteBooks;
			}));
	}
}