import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';

import { AppComponent } from './app.component';
import { SuccessComponent } from './success/success.component';
import { FailureComponent } from './failure/failure.component';
import { ApplicantFormComponent } from './applicant-form/applicant-form.component';

const appRoutes = [
  {path: '', component: ApplicantFormComponent },
  {path: 'success', component: SuccessComponent },
  {path: 'failure', component: FailureComponent },
  {path: '**', component: PageNotFoundComponent },
];

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
