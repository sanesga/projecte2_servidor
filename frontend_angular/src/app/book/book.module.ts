import { ModuleWithProviders, NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { BookComponent } from './book.component';
import { BookResolver } from './book-resolver.service';
import { MarkdownPipe } from './markdown.pipe';
import { SharedModule } from '../shared';
import { BookRoutingModule } from './book-routing.module';

@NgModule({
  imports: [
    SharedModule,
    BookRoutingModule
  ],
  declarations: [
    BookComponent,
    MarkdownPipe
  ],

  providers: [
    BookResolver
  ]
})
export class BookModule {}
