import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

import { Errors, UserService } from '../core';
import { BehaviorSubject } from 'rxjs';

@Component({
  selector: 'app-authSocial-page',
  templateUrl: './auth-social.component.html'
})
export class AuthSocialComponent implements OnInit {
  authType: String = '';
  title: String = '';
  errors: Errors = { errors: {} };
  isSubmitting = false;
  authForm: FormGroup;
  

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService,
    private fb: FormBuilder

  ){
  }

  ngOnInit() {
    this.route.url.subscribe(data => {
      // Get the last piece of the URL
      this.authType = data[data.length - 1].path;
     console.log(this.authType)
        this.loginSocial();
    });

  }

  loginSocial() {
   // console.log("entra en login social en auth-social.cmponent")
  
    this.isSubmitting = true;
    this.errors = { errors: {} };

      // const credentials = this.authForm.value;
     
      var user = {
        email: this.authType,
        password: "12345678"
    };

    console.log(user) 
    console.log(typeof(user))
      
    this.userService
      .attemptAuth('login', user)
      .subscribe(
        data => this.router.navigateByUrl('/'),
        err => {
          this.errors = err;
          this.isSubmitting = false;
        }
      );
  }
}
