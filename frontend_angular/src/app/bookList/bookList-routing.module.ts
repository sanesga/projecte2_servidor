import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BookListComponent } from './bookList.component';
import { BookListResolver } from './bookList-resolver.service';


const routes: Routes = [
    
    {
      path: ':slug',
      component: BookListComponent,
      resolve: {
        book: BookListResolver /*obtiene la información del resolver, que a su vez llama al service*/
      }
    }
  ];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BookListRoutingModule {}
