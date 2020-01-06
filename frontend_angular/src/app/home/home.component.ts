import { Component, OnInit } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { ArticleListConfig, TagsService, UserService, BooksService } from '../core';
import { RedisService } from '../core';

@Component({
  selector: 'app-home-page',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {
  constructor(
    private router: Router,
    // private tagsService: TagsService,
    private userService: UserService,
    private route: ActivatedRoute,
    // private booksService: BooksService,
    private redisService: RedisService,
  ) { }

  isAuthenticated: boolean;
  listConfig: ArticleListConfig = {
    type: 'all',
    filters: {}
  };
  //tags: Array<string> = [];
  //tagsLoaded = false;

  books: Array<Object> = [];



  ngOnInit() {
    this.userService.isAuthenticated.subscribe(
      (authenticated) => {
        this.isAuthenticated = authenticated;

        // set the article list accordingly
        // if (authenticated) {
        //   this.setListTo('feed');
        // } else {
        //   this.setListTo('all');
        // }
      }
    );

    // this.tagsService.getAll()
    // .subscribe(tags => {
    //   this.tags = tags;
    //   this.tagsLoaded = true;
    // });


    /*obtenemos de redis los libros más votados */
    this.redisService.getAll().subscribe(data => {
     
      //data es un array de objetos
      for (const key in data.keys) {
        //imprime la clave
        //console.log(key)
        //imprime el valor
        //console.log(data.keys[key]);

        //pasamos el número de favoritos, de string a number
        var likes = parseInt(data.keys[key]);

        //si el libro tiene más de 5 likes, se mostrará en el home, si no, no.
        if (likes >= 5) { //SI EN EL HOME NO SE IMPRIMIE NADA ES QUE NO HAY LIBROS CON MÁS DE 5 LIKES

          //creamos la propiedad título, sustituyendo los guiones por espacios para imprimir en el home
          var title = new String();

          for (var j = 0; j < key.length; j++) {
            title = key.split("-").join(" ");
          }
          //creamos un objeto con el título y con el slug (el título lo imprimimos y el slug lo pasamos al componente detalle)
          var favBook = {
            title: title,
            slug: key
          }
          //guardamos los objetos en el array de libros con más favoritos
          this.books.push(favBook);
        }
      }
    });

    //si quisiéramos obtener uno
    // this.redisService.getOne(this.book.slug).subscribe(data=>{
    //   console.log(data);
    //   return data;
    // });

    //antes obteníamos la lista de libros desde la base de datos
    // this.booksService.getAll()
    // .subscribe(books => {
    //   this.books = books;
    // });

  }

  // setListTo(type: string = '', filters: Object = {}) {
  //   // If feed is requested but user is not authenticated, redirect to login
  //   if (type === 'feed' && !this.isAuthenticated) {
  //     this.router.navigateByUrl('/login');
  //     return;
  //   }

  //   // Otherwise, set the list object
  //   this.listConfig = {type: type, filters: filters};
  // }
}
