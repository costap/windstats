import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from "@angular/forms";

import { AppComponent } from './app.component';
import { OverviewComponent } from './overview/overview.component';
import { RealtimeComponent } from './realtime/realtime.component';
import { FiveminavgComponent } from './fiveminavg/fiveminavg.component';
import { WindStatsComponent } from './wind-stats/wind-stats.component';

@NgModule({
  declarations: [
    AppComponent,
    OverviewComponent,
    RealtimeComponent,
    FiveminavgComponent,
    WindStatsComponent
  ],
  imports: [
    BrowserModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
