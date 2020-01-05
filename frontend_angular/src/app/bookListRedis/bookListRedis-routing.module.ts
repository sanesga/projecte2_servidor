import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';
import { BookListRedisComponent } from './bookListRedis.component';
import { BookListRedisResolver } from './bookListRedis-resolver.service';


const routes: Routes = [
    
    {
      path: ':slug',
      component: BookListRedisComponent,
      resolve: {
        book: BookListRedisResolver /*obtiene la información del resolver, que a su vez llama al service*/
      }
    }
  ];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BookListRedisRoutingModule {}
