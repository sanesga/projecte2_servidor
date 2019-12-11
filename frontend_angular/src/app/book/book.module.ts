import { NgModule } from '@angular/core';
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
    MarkdownPipe
  ],
  providers: [
    BookResolver
  ]
})
export class BookModule {}
