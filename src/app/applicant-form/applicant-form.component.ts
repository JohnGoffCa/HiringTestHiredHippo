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
    this.http.post('http://ec2-13-59-135-248.us-east-2.compute.amazonaws.com:3000/entries', this.model)
      .subscribe(
        data => {
          this.data = data;
          this.id = JSON.parse(this.data).applicant_id;
// TODO remove this
console.log(this.id);
          this.redirectToVictoryPage();
        }, err => {
          this.err = true;
        }
      );
  }

  redirectToVictoryPage() {
    this.http.get('http://ec2-13-59-135-248.us-east-2.compute.amazonaws.com:3000/status/' + this.id)
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
