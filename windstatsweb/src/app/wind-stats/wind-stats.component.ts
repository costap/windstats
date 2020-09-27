import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-wind-stats',
  templateUrl: './wind-stats.component.html',
  styleUrls: ['./wind-stats.component.css']
})
export class WindStatsComponent implements OnInit {

  allowSomething: boolean = false;
  status = "Off";
  username: string = "default";
  servers = ['Testserver', 'Testserver 2']

  constructor() { 
    this.status = Math.random() > 0.5 ? "Off" : "On";
    setTimeout(()=>{
      this.allowSomething = true;
    }, 2000);
  }

  ngOnInit(): void {
  }

  onDoSomething(){
    this.status = "On";
    this.servers.push("New Server");
  }

  getColor() {
    return this.status === 'On' ? 'red' : 'greed';
  }
}
