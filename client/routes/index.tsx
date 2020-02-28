import * as React from "react";
import { Route, Router, Switch } from "react-router-dom";
import { createBrowserHistory } from "history";

import NotFoundComponent from "../components/NotFound/NotFoundComponent";
import DashboardComponent from "../components/Dashboard/DashboardComponent";

export default (
  <Router history={createBrowserHistory()}>
    <Switch>
      <Route exact path="/" component={DashboardComponent}/>
      <Route component={NotFoundComponent}/>
    </Switch>
  </Router>
);
