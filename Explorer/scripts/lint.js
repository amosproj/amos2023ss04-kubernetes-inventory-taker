// Simple script to run eslint via precommit and fix the
// path of the js files to not have frontend/ in the path
// Inspired by https://eslint.org/docs/latest/integrate/nodejs-api#eslint-class
const filesToLint = process.argv
    .slice(2)
    .map(path => path.replace(/^Explorer\//, ""));

const { ESLint } = require("eslint");

(async function main() {
    // 1. Create an instance with the `fix` option.
    const eslint = new ESLint({ fix: true });

    // 2. Lint files. This doesn't modify target files.
    const results = await eslint.lintFiles(filesToLint);

    // 3. Modify the files with the fixed code.
    await ESLint.outputFixes(results);

    // 4. Format the results.
    const formatter = await eslint.loadFormatter("stylish");
    const resultText = formatter.format(results);

    // 5. Output it.
    console.log(resultText);

    if (results.map(e => e.messages.length).reduce((acc, curr) => acc + curr, 0) != 0) {
        process.exitCode = 1; // If there were changes the pre-commit check should fail
    }
})().catch((error) => {
    process.exitCode = 1;
    console.error(error);
});
