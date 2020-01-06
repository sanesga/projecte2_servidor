import { Component, Input } from '@angular/core';
import { Book } from '../../core';
import { RedisService } from '../../core';

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

    //save to redis
    this.redisService.save({key: this.book.slug, value: this.book.favoritesCount}).subscribe(data=>{
      return data;
    })
 
  } else {
    this.book.favoritesCount--;

    //save to redis
    this.redisService.save({key: this.book.slug, value: this.book.favoritesCount}).subscribe(data=>{
      return data;
    })
  }
}

}
