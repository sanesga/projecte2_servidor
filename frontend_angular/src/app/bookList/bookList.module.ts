import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { BookListComponent } from './bookList.component';
import { BookListResolver } from './bookList-resolver.service';
import { MarkdownPipe } from './markdown.pipe';
import { SharedModule } from '../shared';
import { BookListRoutingModule } from './bookList-routing.module';

@NgModule({
  imports: [
    SharedModule,
    BookListRoutingModule
  ],
  declarations: [
 
    MarkdownPipe
  ],

  providers: [
    BookListResolver
  ]
})
export class BookListModule {}
