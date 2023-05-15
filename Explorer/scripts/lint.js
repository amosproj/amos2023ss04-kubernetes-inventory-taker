// Simple script to run eslint via precommit and fix the
// path of the js files to not have frontend/ in the path

const filesToLint = process.argv
  .slice(2)
  .map((path) => path.replace(/^Explorer\//, ""));

const { ESLint } = require("eslint");

(async function main() {
  // 1. Create an instance.
  const eslint = new ESLint({ fix: true });

  // 2. Lint files.
  const results = await eslint.lintFiles(filesToLint);

  // 3. Format the results.
  const formatter = await eslint.loadFormatter("stylish");
  const resultText = formatter.format(results);

  // 4. Output it.
  console.log(resultText);
})().catch((error) => {
  process.exitCode = 1;
  console.error(error);
});
