import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService, User } from '../core';

@Component({
  selector: 'app-authSocial-page',
  templateUrl: './auth-social.component.html'
})
export class AuthSocialComponent implements OnInit {
  authType: String = '';
  isSubmitting = false;
  user: User;
  
  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private userService: UserService,
  ){
  }

  ngOnInit() {
    //recogemos el nombre del usuario que viene por el path
    this.route.url.subscribe(data => {
      this.authType = data[data.length - 1].path;
        this.loginSocial();
    });

  }

  loginSocial() {
    //recuperamos el usuario de la base de datos, a través del nombre
    this.userService.getUser(this.authType)
      .subscribe(data =>{
        this.user = data
        //redirigimos al home
        this.router.navigateByUrl('/')
      } );
  }
}
