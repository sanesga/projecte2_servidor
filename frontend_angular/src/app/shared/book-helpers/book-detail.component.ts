import { Component, Input } from '@angular/core';

import { Book } from '../../core';

@Component({
  selector: 'app-book-detail',
  templateUrl: './book-detail.component.html'
})
export class BookDetailComponent {
   @Input() book: Book;

seeDetails(book){
console.log(book);
}

 
}
