import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BookComponent } from './book.component';
import { BookResolver } from './book-resolver.service';

console.log("entra al routing de book");

const routes: Routes = [
  {
    path: ':slug',
    component: BookComponent,
    resolve: {
      article: BookResolver
    }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BookRoutingModule {}
