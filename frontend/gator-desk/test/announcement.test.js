import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { AnnouncemetService } from './announcement.service';
const axios = require('axios')

describe('AppInterceptor', () => {
    it ("using axios",  ()=>{
      fetch('http://localhost:8181/listAllAnnouncements', {
      method: 'GET',
      credentials: 'include',
      headers: {
        'Accept': '*/*'
      }
    })
      .then(response => {
        console.log(response)
      })
      .then(jsondata => {
        console.log(jsondata)
      })
    });
  });