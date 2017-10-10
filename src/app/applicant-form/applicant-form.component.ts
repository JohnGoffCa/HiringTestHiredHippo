import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Router } from '@angular/router';
import { Observable } from 'rxjs/Rx';
import { Applicant } from '../applicant';
import { ApplicantResponse } from '../applicant-response';

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
    this.http.post('api/entries', this.model)
      .subscribe(
        data => {
          this.data = data;
          this.id = JSON.parse(this.data).applicant_id;
// TODO remove this
console.log(this.data);
          this.redirectToVictoryPage();
        }, err => {
          this.err = true;
          console.log(err);
        }
      );
  }

  redirectToVictoryPage() {
    this.http.get('api/status/' + this.id)
      .subscribe(
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
}
