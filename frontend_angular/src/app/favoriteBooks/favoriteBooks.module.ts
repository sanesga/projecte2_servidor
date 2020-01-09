import { NgModule } from '@angular/core';
import { SharedModule } from '../shared';
import { FavoriteBooksRoutingModule } from './favoriteBooks-routing.module';
import { FavoriteBooksComponent } from './favoriteBooks.component';

@NgModule({
  declarations: [FavoriteBooksComponent],
  imports: [
    SharedModule,
    FavoriteBooksRoutingModule
  ],
})
export class FavoriteBooksModule { }
