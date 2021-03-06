import { CommonModule } from '@angular/common';
import { NgModule } from '@angular/core';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';
import { RouterModule } from '@angular/router';
import { ArticleListComponent, ArticleMetaComponent, ArticlePreviewComponent } from './article-helpers';
import { FavoriteButtonComponent, FollowButtonComponent } from './buttons';
import { ListErrorsComponent } from './list-errors.component';
import { ShowAuthedDirective } from './show-authed.directive';
import { BookComponent } from '../book/book.component';
import { BookDetailComponent } from './book-helpers';
import { BookListComponent } from '../bookList/bookList.component';
import { BookCommentComponent } from '../book/book-comment.component';
import { FavBookComponent } from './book-helpers/fav-book.component';
import { FavoriteBookButtonComponent } from './buttons/favorite-book-button.component';
import { FavoriteBooksComponent } from '../favoriteBooks/favoriteBooks.component';

@NgModule({
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    RouterModule
  ],
  declarations: [
    ArticleListComponent,
    ArticleMetaComponent,
    ArticlePreviewComponent,
    FavoriteButtonComponent,
    FollowButtonComponent,
    ListErrorsComponent,
    ShowAuthedDirective,
    BookComponent,
    BookDetailComponent,
    BookListComponent,
    BookCommentComponent,
    FavBookComponent,
    FavoriteBookButtonComponent,
    FavoriteBooksComponent
  ],
  exports: [
    ArticleListComponent,
    ArticleMetaComponent,
    ArticlePreviewComponent,
    CommonModule,
    FavoriteButtonComponent,
    FollowButtonComponent,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    ListErrorsComponent,
    RouterModule,
    ShowAuthedDirective,
    BookComponent,
    BookDetailComponent,
    BookListComponent,
    BookCommentComponent,
    FavBookComponent,
    FavoriteBookButtonComponent,
    FavoriteBooksComponent
  ]
})
export class SharedModule {}
