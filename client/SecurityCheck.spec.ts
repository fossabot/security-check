import SecurityCheck from "./SecurityCheck";

describe("SecurityCheck", () => {
  it("can be initialized", () => {
    expect(new SecurityCheck()).toBeInstanceOf(SecurityCheck);
  });
});
