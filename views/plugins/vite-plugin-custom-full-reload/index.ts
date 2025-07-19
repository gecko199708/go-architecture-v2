import type { HmrContext, Plugin } from "vite";

export default function customFullReload(): Plugin {
	return {
		name: "custom-full-reload",
		handleHotUpdate(ctx: HmrContext) {
			ctx.server.ws.send({
				type: "full-reload",
				path: ctx.file.replace(/\.html$/, ""),
			});
		},
	};
}
