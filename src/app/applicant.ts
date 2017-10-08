import { OnInit } from '@angular/core';

export class Applicant implements OnInit {
  id: string;

  constructor(
    public applicant_name,
    public applicant_email,
    public phone_number,
  ) { }

  ngOnInit() {
  }
}
