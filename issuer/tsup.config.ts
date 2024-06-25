import { defineConfig } from "tsup";
import type { Options, Format } from "tsup";

// Ensure that these option fields are not undefined
type MandatoryOptions = Options & {
  outDir: string;
  platform: string;
  format: Format | Format[];
};

// Default config, used as a base template
const DEFAULT_CONFIG: Options = {
  bundle: true,
  clean: true,
  dts: true,
  minify: true,
  entry: ["src/server.ts"],
  sourcemap: true,
  splitting: true,
  target: "es2022",
  // bundle everything for runtime compatibility.
  skipNodeModulesBundle: false,
  noExternal: [],
};

// Common.js config
const COMMON_CONFIG: MandatoryOptions = {
  ...DEFAULT_CONFIG,
  format: "cjs",
  outDir: "dist/common",
  platform: "node",
};

export default defineConfig([
  COMMON_CONFIG,
]);
