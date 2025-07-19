import inputPlugin from "@macropygia/vite-plugin-glob-input";
import suidPlugin from "@suid/vite-plugin";
import { resolve } from "node:path";
import { type AliasOptions, defineConfig } from "vite";
// @ts-ignore
import handlebars from "vite-plugin-handlebars";
import solid from "vite-plugin-solid";

import customFullReload from "./plugins/vite-plugin-custom-full-reload";

const options = {
	vite: {
		root: "./",
		publicDir: "public",
		assetsInclude: "public/**/*.*",
		build: {
			outDir: "dist",
			sourcemap: true,
		} satisfies ViteBuildOptions,
	} satisfies ViteOptions,
	server: {} as ViteServerOptions satisfies ViteServerOptions,
	plugins: {
		inputPlugin: {
			patterns: ["./src/**/*.html"] as string[],
		},
		handlebars: {
			partialDirectory: resolve(__dirname, "./src/common/static/components"),
		},
	},
	alias: {
		"@dynamic": resolve(__dirname, "./src/common/dynamic"),
	} satisfies PathAlias,
} as const;

// https://ja.vite.dev/config/
export default defineConfig({
	...options.vite,
	server: {
		...options.server,
		port:
			Number(process.env.VITE_PORT) || options.server?.defaultPort || undefined,
		open: process.env.VITE_OPEN_BROWSER === "true",
		strictPort: true,
	},
	...{
		plugins: [
			solid(),
			suidPlugin(),
			inputPlugin(options.plugins.inputPlugin),
			handlebars(options.plugins.handlebars),
			customFullReload(),
		],
		css: {
			modules: {
				localsConvention: "dashes",
			},
		},
		resolve: {
			alias: options.alias as AliasOptions,
		},
	},
});

type ViteOptions = Required<
	Pick<
		import("vite").UserConfig,
		"root" | "publicDir" | "assetsInclude" | "build"
	>
>;

type ViteBuildOptions = Required<
	Pick<ViteOptions["build"], "outDir" | "sourcemap">
>;

type ViteServerOptions = Pick<
	import("vite").ServerOptions,
	"host" | "watch"
> & {
	defaultPort?: number;
};

type PathAlias = Record<`@${string}`, string>;
