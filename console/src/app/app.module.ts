import { NgModule } from '@angular/core';
import { BrowserModule, Title } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

// environment
import { environment } from 'environments/environment';

import { Routings } from './app.routing';

// Third-party
import { BootstrapModalModule } from 'angular2-modal/plugins/bootstrap';
import { BsDropdownModule } from 'ngx-bootstrap/dropdown';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { AngularFireModule } from 'angularfire2';
import { AngularFireDatabaseModule } from 'angularfire2/database';
import { AngularFireAuthModule } from 'angularfire2/auth';

// Guards
import { AuthGuard } from './guards/auth.guard';

// Components
import { AppComponent } from './app.component';
import { ChannelComponent } from './components/channel/channel.component';
import { ProgramComponent } from './components/program/program.component';
import { HomeComponent } from './components/home/home.component';

// Services
import { AuthService } from './services/auth.service';
import { ChannelsService } from './services/channels.service';
import { ProgramsService } from './services/programs.service';

@NgModule({
    declarations: [
        AppComponent,
        ChannelComponent,
        ProgramComponent,
        HomeComponent
    ],
    imports: [
        AngularFireModule.initializeApp(environment.firebase, 'shucyan'),
        AngularFireDatabaseModule,
        AngularFireAuthModule,
        BsDropdownModule.forRoot(),
        BrowserModule,
		CollapseModule.forRoot(),
        Routings
    ],
    providers: [
        AuthService,
        AuthGuard,
        Title,
        ChannelsService,
        ProgramsService
    ],
    bootstrap: [AppComponent]
})
export class AppModule { }
