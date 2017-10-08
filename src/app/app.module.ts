import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { SuccessComponent } from './success/success.component';
import { FailureComponent } from './failure/failure.component';
import { ApplicantFormComponent } from './applicant-form/applicant-form.component';

@NgModule({
  declarations: [
    AppComponent,
    SuccessComponent,
    FailureComponent,
    ApplicantFormComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
