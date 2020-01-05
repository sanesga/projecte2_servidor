import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { BookListRedisComponent } from './bookListRedis.component';
import { BookListRedisResolver } from './bookListRedis-resolver.service';
import { MarkdownPipe } from './markdown.pipe';
import { SharedModule } from '../shared';
import { BookListRedisRoutingModule } from './bookListRedis-routing.module';

@NgModule({
  imports: [
    SharedModule,
    BookListRedisRoutingModule
  ],
  declarations: [
 
    MarkdownPipe
  ],

  providers: [
    BookListRedisResolver
  ]
})
export class BookListRedisModule {}
