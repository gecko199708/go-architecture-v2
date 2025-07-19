import type { JSX } from "solid-js";
import { render } from "solid-js/web";

type MountPoint = `app-${string}`;

type Applications = {
	mountPoint: MountPoint;
	app: () => JSX.Element;
}[];

function Mount(apps: Applications) {
	for (const app of apps) {
		const mountPoint = document.getElementById(app.mountPoint);
		if (mountPoint) {
			render(app.app, mountPoint);
		} else {
			console.error(`Mount point ${app.mountPoint} not found`);
		}
	}
}

export { Mount };
export type { Applications };
