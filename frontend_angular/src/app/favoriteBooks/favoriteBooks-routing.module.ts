import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FavoriteBooksComponent } from './favoriteBooks.component';


const routes: Routes = [
  {
    path: '',
    component: FavoriteBooksComponent
  }
]

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class FavoriteBooksRoutingModule { }
