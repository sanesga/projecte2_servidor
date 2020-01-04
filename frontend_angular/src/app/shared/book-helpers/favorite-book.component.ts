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

  
    this.redisService.getAll().subscribe(data => {
      // this.books = books;
      console.log(data);
      return data;
    });

    //getOne
    this.redisService.getOne("sandra").subscribe(data=>{
      console.log(data);
      return data;
    });

    //save
    this.redisService.save({key: this.book.title, value: this.book.favoritesCount}).subscribe(data=>{
      console.log(data);
      return data;
    })
 
  } else {
    this.book.favoritesCount--;
  }
}

}
