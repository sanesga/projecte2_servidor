import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BookComponent } from './book.component';
import { BookResolver } from './book-resolver.service';

const routes: Routes = [
    {
      path: '',
      component: BookComponent,
      resolve: {
        books: BookResolver /*obtiene la información del resolver, que a su vez llama al service*/
      }
    }
  ];


@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BookRoutingModule {}
