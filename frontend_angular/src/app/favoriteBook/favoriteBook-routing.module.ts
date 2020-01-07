import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { FavoriteBookComponent } from './favoriteBook.component';


const routes: Routes = [
  {
    path: '',
    component: FavoriteBookComponent,
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class favoriteBookRoutingModule {}