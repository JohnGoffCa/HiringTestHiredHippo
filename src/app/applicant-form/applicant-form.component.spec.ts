import { async, ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterOutlet, Router } from '@angular/router';
import { RouterTestingModule } from '@angular/router/testing';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { ApplicantFormComponent } from './applicant-form.component';

describe('ApplicantFormComponent', () => {
  let component: ApplicantFormComponent;
  let fixture: ComponentFixture<ApplicantFormComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        FormsModule,
        HttpClientModule
      ],
      declarations: [ ApplicantFormComponent ],
      providers: [ RouterOutlet ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(ApplicantFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
  it('should redirect to victory page on success', () => {
  });
  it('should redirect to loss page on failure', () => {
  });
});
