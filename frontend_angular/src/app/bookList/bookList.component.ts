import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { FormControl } from '@angular/forms';


import {
  Comment,
  CommentsService,
  User,
  UserService,
  Book
} from '../core';

@Component({
  selector: 'app-bookList-page',
  styleUrls: ['bookList.component.css'],
  templateUrl: './bookList.component.html'
})
export class BookListComponent implements OnInit {
  book: Book;
   commentControl = new FormControl();
   commentFormErrors = {};
   comments: Comment[];
   currentUser: User;
   canModify: boolean;
   isSubmitting = false;
   isDeleting = false;

  constructor(
    private route: ActivatedRoute,
    private commentsService: CommentsService,
  ) { }

  ngOnInit() {
    this.route.data.subscribe(
      (data: { book: Book }) => {
        this.book = data.book;

        this.populateComments();
      }
    );
  }
  populateComments() {
    this.commentsService.getAll(this.book.slug)
      .subscribe(comments =>{
        // console.log(comments); //recibimos el array de comentarios
        this.comments = comments;
      } );
  }
  onDeleteComment(comment) {
    this.commentsService.destroy(comment.id, this.book.slug)
      .subscribe(
        success => {
          this.comments = this.comments.filter((item) => item !== comment);
        }
      );
  }
   addComment() {
    this.isSubmitting = true;
   this.commentFormErrors = {};

    const commentBody = this.commentControl.value;
    this.commentsService
      .add(this.book.slug, commentBody)
      .subscribe(
        comment => {
          this.comments.unshift(comment);
          this.commentControl.reset('');
          this.isSubmitting = false;
        },
        errors => {
          this.isSubmitting = false;
          this.commentFormErrors = errors;
        }
      );
  }

}


