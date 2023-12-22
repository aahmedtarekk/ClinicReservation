import { Component } from '@angular/core';
import { Router } from '@angular/router';
import { HttpClient } from '@angular/common/http';
import { CurrentUser } from 'src/app/CurrentUser';
import { MyResponse } from 'src/MyResponse';
@Component({
  selector: 'app-sign-in',
  templateUrl: './sign-in.component.html',
  styleUrls: ['./sign-in.component.css']
})
export class SignInComponent {
  userName!: string; 
  password!: string; 
  userType: string = 'Patient'

  constructor(private user: CurrentUser,private router: Router, private http : HttpClient) { }

  signIn() {
    var object = {
      "Name" : this.userName,
      "Password" : this.password,
      "Type" : this.userType
    }
    this.http.post<MyResponse>('http://127.0.0.1:3001/ClinicReservation/SignIn', object).subscribe((response) => {
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
