import { Component, OnInit } from '@angular/core';
import { Applicant } from '../applicant';

@Component({
  selector: 'app-applicant-form',
  templateUrl: './applicant-form.component.html',
  styleUrls: ['./applicant-form.component.css']
})
export class ApplicantFormComponent {
  model = new Applicant('', '');

  onSubmit() {
  }
  // TODO: Remove this when we're done
  get diagnostic() { return JSON.stringify(this.model); }
}
