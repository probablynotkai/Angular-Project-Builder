import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root'
})
export class HttpService {
    apiUrl = "https://catfact.ninja/";

    constructor(private http: HttpClient) { }

    getCatFact() {
      return this.http.get(this.apiUrl + "fact");
    }
}