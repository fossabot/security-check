import * as React from "react";
import { Route, Router, Switch } from "react-router-dom";
import { createBrowserHistory } from "History";

import NotFoundComponent from "../components/NotFound/NotFoundComponent";

export default (
  <Router history={createBrowserHistory()}>
    <Switch>
      <Route component={NotFoundComponent} />
    </Switch>
  </Router>
);
