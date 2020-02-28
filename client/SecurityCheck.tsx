import * as React from "react";
import * as ReactDOM from "react-dom";
import { Provider } from "react-redux";
import { ThemeProvider } from '@material-ui/core/styles';

import store from "./store";
import routes from "./routes";
import { theme } from "./constants/theme";

export default class SecurityCheck {
  render() {
    ReactDOM.render(
      <Provider store={store}>
        <ThemeProvider theme={theme}>{routes}</ThemeProvider>
      </Provider>,
      document.querySelector(".root")
    );
  }
}
