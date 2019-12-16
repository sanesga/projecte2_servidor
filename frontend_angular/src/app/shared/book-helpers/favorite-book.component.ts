import { Component, Input } from '@angular/core';

import { Book } from '../../core';

@Component({
  selector: 'app-favorite-book',
  styleUrls: ['favorite-book.component.css'],
  templateUrl: './favorite-book.component.html'
})
export class FavoriteBookComponent {
   @Input() book: Book;

// seeDetails(book){
// console.log(book);
// }
onToggleFavorite(favorited: boolean) {
  this.book.favorited = favorited;

  if (favorited) {
    this.book.favoritesCount++;
  } else {
    this.book.favoritesCount--;
  }
}
}
