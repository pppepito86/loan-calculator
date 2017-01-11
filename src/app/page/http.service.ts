import { Injectable } from '@angular/core';
import { Http, Response } from '@angular/http';
import 'rxjs/Rx';

@Injectable()
export class HttpService {

  constructor(private http: Http) { }

  getData(ammount: string, term: string, interest: string) {
    return this.http.get("http://localhost:3005/loan/ammount/"+ammount+"/term/"+term+"/interest/"+interest)
      .map((response: Response) => response.json());
  }

}
