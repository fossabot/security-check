import * as React from "react";
import * as TestRenderer from "react-test-renderer";
import { ReactTestRenderer } from "react-test-renderer";

import NotFoundComponent from "./NotFoundComponent";

describe("NotFoundComponent", () => {
  const testRenderer: ReactTestRenderer = TestRenderer.create(
    <NotFoundComponent />
  );

  it("contains 404", () => {
    expect(
      testRenderer
        .toJSON()
        // @ts-ignore
        .children[0].children.join("")
        .trim()
    ).toContain("404");
  });
});
