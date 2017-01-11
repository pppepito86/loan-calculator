import { Component } from '@angular/core';
import { HttpService } from "./page/http.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  loanInfo = {
    month_payment: "",
    total_payment: "",
    interest_payment: ""
  }; 
  ammount: string = ""
  term: string = ""
  interest: string = ""
  month_payment: any = ""

  constructor(private httpService: HttpService) {
  }

  valueChanged(event: any) {
    if ( event.target.id === "loan_ammount") this.ammount=event.target.value;
    if ( event.target.id === "loan_term") this.term=event.target.value;
    if ( event.target.id === "loan_interest") this.interest=event.target.value;
   
    if (this.ammount !== "" && this.term !== "" && this.interest !== "") {
      this.httpService.getData(this.ammount, this.term, this.interest).subscribe(data => {this.loanInfo = data;});
    }
  }
}
