import { blue, green, red } from "@suid/material/colors";
import { createTheme } from "@suid/material/styles";

export const theme = createTheme({
	// example
	palette: {
		primary: {
			main: green[300],
		},
		secondary: {
			main: red[300],
		},
		info: {
			main: blue[300],
		},
	},
});
