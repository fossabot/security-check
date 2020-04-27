//------------------------------------------------------------------------------
//
// Husky Configuration, https://github.com/typicode/husky
//
//------------------------------------------------------------------------------

module.exports = {
  hooks: {
    "pre-commit": "yarn format && gofmt -w ./go && gofmt -w ./main.go"
  }
};
