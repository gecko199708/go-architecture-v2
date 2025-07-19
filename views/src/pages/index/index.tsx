import { type Applications, Mount } from "@dynamic/application";
import App from "./App";

const apps = [
	{
		mountPoint: "app-main",
		app: () => <App />,
	},
] as Applications;

Mount(apps);
