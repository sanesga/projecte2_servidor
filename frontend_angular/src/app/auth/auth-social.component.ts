import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormControl, Validators } from '@angular/forms';
import { ActivatedRoute, Router } from '@angular/router';

import { Errors, UserService, User } from '../core';
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
  user: User;
  

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
     //console.log(this.authType)
        this.loginSocial();
    });

  }

  loginSocial() {
   // console.log("entra en login social en auth-social.cmponent")
  
    this.isSubmitting = true;
    this.errors = { errors: {} };

     
    //RECUPERAMOS EL USUARIO DE LA BASE DE DATOS A TRAVÉS DEL NOMBRE
    this.userService.getUser(this.authType)
      .subscribe(data =>{
        console.log(data)
        this.user = data
      } );
  }
}
