import { type Applications, Mount } from "@dynamic/application";
import App from "./App";

const apps = [
	{
		mountPoint: "app-main",
		app: () => <App />,
	},
] satisfies Applications;

Mount(apps);
