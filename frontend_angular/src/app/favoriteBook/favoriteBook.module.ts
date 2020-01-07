import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FavoriteBookComponent } from './favoriteBook.component';
import { favoriteBookRoutingModule } from './favoriteBook-routing.module';
import { FavoriteBookService } from '../core';


@NgModule({
  declarations: [FavoriteBookComponent],
  imports: [
    CommonModule,
    favoriteBookRoutingModule,
  ],
  providers: [FavoriteBookService]
})
export class favoriteBookModule { }
