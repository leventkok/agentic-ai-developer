# Emit comparison

## tsc (per-file emit)
Run: `npm run build`
Inspect: `dist/app.js`, `dist/cjs-import.js`
- Keeps import structure
- Does not bundle dependencies (lodash stays external)

## esbuild (bundled)
Run: `npm run bundle`
Inspect: `dist/bundle.js`
- Single file, lodash code inlined
- Smaller deploy unit, tree-shaking possible

## TODO
After both builds, note 2 differences you see in the output files.
