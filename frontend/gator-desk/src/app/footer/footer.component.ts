import { Component, ElementRef, ViewChild, AfterViewInit } from '@angular/core';
import { MatToolbar } from '@angular/material';
@Component({
selector: 'my-app',
templateUrl: './footer.component.html',
styleUrls: [ './footer.component.css' ] })
export class FooterComponent implements AfterViewInit {
@ViewChild("toolbar") toolbar!: MatToolbar;
@ViewChild("main") main!: ElementRef;
@ViewChild("footer") footer!: ElementRef;
constructor() { }
ngOnInit(){
}
ngAfterViewInit(){
}
}