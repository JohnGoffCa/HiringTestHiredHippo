import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { SuccessComponent } from './success/success.component';
import { FailureComponent } from './failure/failure.component';

@NgModule({
  declarations: [
    AppComponent,
    SuccessComponent,
    FailureComponent
  ],
  imports: [
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
