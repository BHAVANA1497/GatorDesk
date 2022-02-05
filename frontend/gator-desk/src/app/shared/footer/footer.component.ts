import { Component, OnInit } from '@angular/core';
import { faCoffee } from '@fortawesome/free-solid-svg-icons';

@Component({
  selector: 'app-footer',
  templateUrl: './footer.component.html',
  styleUrls: ['./footer.component.css'],
})
export class FooterComponent implements OnInit {
  faCoffee = faCoffee;
  textColor: string = '#4FBDBA';
  constructor() {}

  ngOnInit(): void {}
}
