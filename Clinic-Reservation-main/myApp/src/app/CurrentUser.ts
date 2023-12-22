import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class CurrentUser {
  public uuid!: string;

  constructor() { }
}
