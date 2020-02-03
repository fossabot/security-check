import * as React from "react";
import * as ReactDOM from "react-dom";
import { Provider } from "react-redux";

import store from "./store";
import routes from "./routes";

export default class SecurityCheck {
  render() {
    ReactDOM.render(
      <Provider store={store}>{routes}</Provider>,
      document.querySelector(".root")
    );
  }
}
