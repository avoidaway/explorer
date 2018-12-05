import {AfterViewInit, Directive, ElementRef, forwardRef, Injector, Input, NgZone, OnDestroy, OnInit} from '@angular/core';
import {ControlValueAccessor, NG_VALUE_ACCESSOR, NgControl, Validators} from '@angular/forms';
import {AbstractControl} from '@angular/forms/src/model';
import {interval, Subscription} from 'rxjs';
import {startWith} from 'rxjs/operators';

/*export interface ReCaptchaConfig {
  theme?: 'dark' | 'light';
  type?: 'audio' | 'image';
  size?: 'compact' | 'normal';
  tabindex?: number;
}*/

declare const grecaptcha: any;

declare global {
  interface Window {
    grecaptcha: any;
    reCaptchaLoad: () => void;
  }
}

@Directive({
  selector: '[appRecaptcha]',
  providers: [
    {
      provide: NG_VALUE_ACCESSOR,
      useExisting: forwardRef(() => ReCaptchaDirective),
      multi: true
    }
  ],
})
export class ReCaptchaDirective implements OnInit, OnDestroy, AfterViewInit, ControlValueAccessor {
  @Input() key: string;
  @Input() actionName: string;

  private control: AbstractControl;

  private onChange: (value: string) => void;
  private onTouched: (value: string) => void;

  private _sub: Subscription;

  constructor(private _element: ElementRef, private _ngZone: NgZone, private _injector: Injector) {
  }

  ngOnInit() {
    this.registerReCaptchaCallback();
    this.addScript();
  }

  ngOnDestroy() {
    this._sub.unsubscribe();
  }

  ngAfterViewInit() {
    this.control = this._injector.get(NgControl).control;
    // this.setValidator();
  }

  writeValue(obj: any): void {
  }

  registerOnChange(fn: any): void {
    this.onChange = fn;
  }

  registerOnTouched(fn: any): void {
    this.onTouched = fn;
  }

  registerReCaptchaCallback() {
    window.reCaptchaLoad = () => {
      this._sub = interval(600000).pipe(
        startWith(0),
      ).subscribe(() => {
        this.getToken();
      });
    };
  }

  getToken() {
    grecaptcha.execute(this.key, {action: this.actionName}).then((token: string) => {
      this.onSuccess(token);
    });
  }

  addScript() {
    const script: HTMLScriptElement = <HTMLScriptElement>document.createElement('script');
    script.src = `https://www.google.com/recaptcha/api.js?onload=reCaptchaLoad&render=${this.key}`;
    script.async = true;
    script.defer = true;
    document.body.appendChild(script);
  }

  onSuccess(token: string) {
    this._ngZone.run(() => {
      this.onChange(token);
      this.onTouched(token);
    });
  }

  private setValidator() {
    this.control.setValidators(Validators.required);
    setTimeout(() => this.control.updateValueAndValidity(), 0);
  }
}
