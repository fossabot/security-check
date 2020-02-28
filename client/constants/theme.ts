import { createMuiTheme } from "@material-ui/core/styles";
import { blue, pink } from "@material-ui/core/colors";

// @see https://material-ui.com/customization/theming/
export const theme = createMuiTheme({
  palette: {
    primary: blue,
    secondary: pink
  }
});
