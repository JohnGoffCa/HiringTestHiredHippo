import { OnInit } from '@angular/core';

export class Applicant implements OnInit {
  id: string;
  applicant_name: string;
  applicant_email: string;

  constructor(applicant_name, applicant_email) {
    this.applicant_name = applicant_name;
    this.applicant_email = applicant_email;
  }

  ngOnInit() {
  }
}
