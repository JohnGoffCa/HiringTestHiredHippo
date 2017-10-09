import { Directive } from '@angular/core';
import { AbstractControl, NG_VALIDATORS, Validator, ValidatorFn, Validators } from '@angular/forms';

export function phoneNumberValidator(): ValidatorFn {
  return (control: AbstractControl): {[key: string]: any} => {
    // regex to test if phone number is valid
    const forbidden = /^([0-9]( |-)?)?(\(?[0-9]{3}\)?|[0-9]{3})( |-)?([0-9]{3}( |-)?[0-9]{4}|[a-zA-Z0-9]{7})$/.test(control.value);
    return !forbidden ? {'appPhoneNumber': control.value} : null;
  };
}

@Directive({
  selector: '[appPhoneNumber]',
  providers: [{provide: NG_VALIDATORS, useExisting: PhoneNumberDirective, multi: true}]
})
export class PhoneNumberDirective implements Validator {
  validate(control: AbstractControl): {[key: string]: any} {
    return phoneNumberValidator()(control);
  }
}
