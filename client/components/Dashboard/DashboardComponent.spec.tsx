import * as React from "react";
import * as TestRenderer from "react-test-renderer";
import { ReactTestRenderer } from "react-test-renderer";

import DashboardComponent from "./DashboardComponent";

describe("DashboardComponent", () => {
  const testRenderer: ReactTestRenderer = TestRenderer.create(
    <DashboardComponent />
  );

  it.skip("", () => {});
});
