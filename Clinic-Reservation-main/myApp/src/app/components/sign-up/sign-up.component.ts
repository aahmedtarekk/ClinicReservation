// sign-up.component.ts

import { HttpClient } from '@angular/common/http';
import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { CurrentUser } from 'src/app/CurrentUser';
import { MyResponse } from 'src/MyResponse';

@Component({
  selector: 'app-sign-up',
  templateUrl: './sign-up.component.html',
  styleUrls: ['./sign-up.component.css']
})
export class SignUpComponent {
  email!: string;
  password!: string;
  userType: string = 'Patient';
  userName!: string;

  constructor(private user: CurrentUser, private router: Router, private http : HttpClient) { }

  registerUser() {
    var object = {
      "Name" : this.userName,
      "Mail" : this.email,
      "Password" : this.password,
      "Type" : this.userType
    }
    this.http.post<MyResponse>('http://127.0.0.1:3001/ClinicReservation/SignUp', object).subscribe((response) => {
      if(!response.ResponseStatus){
        alert(response.ResponseMessage)
      }
      else{
        this.user.uuid = response.UserUUID
        if (this.userType === 'Patient') {
          this.router.navigate(['/patient']);
        } else if (this.userType === 'Doctor') {
          this.router.navigate(['/doctor']);
        }
      }
    })
  }
}
