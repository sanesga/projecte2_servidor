import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

import {
  Book,
  BooksService,
  User,
  UserService
} from '../core';

@Component({
  selector: 'app-book-page',
  templateUrl: './book.component.html'
})
export class BookComponent implements OnInit {
  book: Book;
  currentUser: User;
  canModify: boolean;
  isSubmitting = false;
  isDeleting = false;

  constructor(
    private route: ActivatedRoute,
    private booksService: BooksService,
    private router: Router,
    private userService: UserService,
  ) { }

  ngOnInit() {
    // Retreive the prefetched book
    this.route.data.subscribe(
      (data: { book: Book }) => {
        console.log(data.book)
        this.book = data.book;
      }
    );
    }
  }
