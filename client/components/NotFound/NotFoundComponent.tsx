import * as React from "react";

import { StyledNotFoundComponent } from "./NotFoundComponent.styles";

export default class NotFoundComponent extends React.Component {
  render() {
    return (
      <StyledNotFoundComponent>
        <h1>404 Not Found</h1>
        <p>Sorry, we seem to have misplaced what you are looking for.</p>
      </StyledNotFoundComponent>
    );
  }
}
