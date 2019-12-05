import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

import { Errors, UserService } from '../core';

@Component({
  selector: 'app-authSocial-page',
  templateUrl: './auth-social.component.html'
})
export class AuthSocialComponent implements OnInit {
  authType: String = '';
  title: String = '';
  errors: Errors = {errors: {}};
  isSubmitting = false;
  authForm: FormGroup;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService,
    private fb: FormBuilder
  ) {
    // use FormBuilder to create a form group
    // this.authForm = this.fb.group({
    //   'email': ['', Validators.required],
    //   'password': ['', Validators.required]
    // });
  }

  ngOnInit() {
    this.route.url.subscribe(data => {
      // Get the last piece of the URL (it's either 'login', 'register' or 'socialLogin')
      this.authType = data[data.length - 1].path;
      
      // Set a title for the page accordingly
      if (this.authType === 'socialLogin') {
        //vamos al social login
        console.log("entramos a socialLogin")
      }

      
    });
  }
}
