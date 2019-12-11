import { Component, EventEmitter, Input, Output, OnInit, OnDestroy } from '@angular/core';

import { Comment, User, UserService } from '../core';
import { Subscription } from 'rxjs';

@Component({
  selector: 'app-book-comment',
  templateUrl: './book-comment.component.html',
  styleUrls: ['book-comment.component.css']
})
export class BookCommentComponent implements OnInit, OnDestroy {

  constructor(
    private userService: UserService
  ) {}

  private subscription: Subscription;

  @Input() comment: Comment;
 
  @Output() deleteComment = new EventEmitter<boolean>();

  canModify: boolean;

  ngOnInit() {
    console.log("entra al controlador de book comment ");
   
   //Load the current user's data
    // this.subscription = this.userService.currentUser.subscribe(
    //   (userData: User) => {
    //     this.canModify = (userData.username === this.comment.author.username);
    //   }
    // );

  }

  ngOnDestroy() {
    this.subscription.unsubscribe();
  }

  deleteClicked() {
    this.deleteComment.emit(true);
  }


}

