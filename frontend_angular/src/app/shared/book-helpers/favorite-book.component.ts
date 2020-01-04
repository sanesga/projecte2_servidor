import { Component, Input } from '@angular/core';
import { RedisService } from '../../core/services/redis.service';
import { Book } from '../../core';

@Component({
  selector: 'app-favorite-book',
  styleUrls: ['favorite-book.component.css'],
  templateUrl: './favorite-book.component.html'
})
export class FavoriteBookComponent {
   @Input() book: Book;

   constructor(
     private redisService: RedisService
     ){}
    
// seeDetails(book){
// console.log(book);
// }
onToggleFavorite(favorited: boolean) {
  
  this.book.favorited = favorited;

  if (favorited) {
    this.book.favoritesCount++;

    //guardamos el libro y el número de favoritos en redis
    this.redisService.save({key: this.book.title, value: this.book.favoritesCount}).subscribe(data => {
      // this.books = books;
      console.log(data);
      return data;
    });

  } else {
    this.book.favoritesCount--;
  }
}

}
