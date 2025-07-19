import { theme } from "@design/theme";
import { ThemeProvider } from "@suid/material";
import type { JSX } from "solid-js";
import { render } from "solid-js/web";

type MountPoint = `app-${string}`;

type Applications = {
	mountPoint: MountPoint;
	app: () => JSX.Element;
}[];

function App(props: { children: JSX.Element }) {
	return <ThemeProvider theme={theme}>{props.children}</ThemeProvider>;
}

function Mount(apps: Applications) {
	for (const app of apps) {
		const mountPoint = document.getElementById(app.mountPoint);
		if (mountPoint) {
			render(() => <App>{app.app()}</App>, mountPoint);
		} else {
			console.error(`Mount point ${app.mountPoint} not found`);
		}
	}
}

export { Mount };
export type { Applications };
