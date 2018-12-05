/*CORE*/
import {Injectable} from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpParams, HttpRequest, HttpResponse} from '@angular/common/http';
import {Observable, throwError} from 'rxjs';
/*UTILS*/
import {environment} from '../../environments/environment';
import {catchError, map, retry} from 'rxjs/operators';
import {ToastrService} from '../modules/toastr/toastr.service';

@Injectable({
  providedIn: 'root'
})
export class ApiService {
  apiURL: string;

  constructor(private http: HttpClient, private toastrService: ToastrService) {
    this.apiURL = this.getApiURL();
  }

  getApiURL() {
    return environment.production
      ? window.location.origin + '/' + environment.API_PATH
      : environment.API_PROTOCOL + '://' + location.hostname + ':' + environment.API_PORT + '/' + environment.API_PATH;
  }

  get(url: string, params?: HttpParams): Observable<any> {
    return this.http.get<any>(this.apiURL + url, {
      params
    }).pipe(
      catchError(this._handleError)
    );
  }

  post(url: string, data?: any): Observable<any> {
    return this.http.post<any>(this.apiURL + url, data).pipe(
      catchError(this._handleError)
    );
  }

  request(method: string, url: string, data?: any) {
    const request = new HttpRequest(method, this.apiURL + url, data);
    return this.http.request(request).pipe(
      catchError(this._handleError),
      map((response: HttpResponse<any>) => response.body),
    );
  }

  private _handleError = (error: HttpErrorResponse) => {
    console.error(
      `Backend returned code ${error.status}, ` +
      `body was: ${error.error}`);
    this.toastrService.danger(error.error.error.message);
    // return an observable with a user-facing error message
    return throwError('Something bad happened; please try again later.');
  }
}

