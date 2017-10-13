import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Rx';
import { Applicant } from '../applicant';
import { ApplicantResponse } from '../applicant-response';
import { environment } from '../../environments/environment';

@Component({
  selector: 'app-applicant-form',
  templateUrl: './applicant-form.component.html',
  styleUrls: ['./applicant-form.component.css']
})
export class ApplicantFormComponent {
  model = new Applicant('', '');
  id: number;
  data: any;
  err = false;

  constructor(private http: HttpClient, private router: Router) {}

  onSubmit() {
    this.http.post(environment.API_URL + '/entries', this.model)
      .subscribe(
        data => {
          this.data = data;
          if (environment.API_URL === 'api') {
            this.id = this.data.applicant_id;
          } else {
            this.id = JSON.parse(this.data).applicant_id; // funky workaround, can't seem to access the data any other way
          }
          this.redirectToVictoryPage(this.id);
        }, err => {
          this.err = true;
          console.log(err);
        }
      );
  }

  redirectToVictoryPage(id: number) {
    this.callApi(environment.API_URL + '/status', id).subscribe(
      data => {
        switch (data['status']) {
          case 'Won': {
            this.router.navigateByUrl('success');
            break;
          }
          case 'Lost': {
            this.router.navigateByUrl('failure');
            break;
          }
        }
      }
    );
  }

  callApi(url: string, id: number) {
    return this.http.get(url + '/' + id);
  }
}
