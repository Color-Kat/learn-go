const readline = require('readline');

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

let lines = [];
let index = 0;

rl.on('line', (line) => {
    lines.push(line);
}).on('close', () => {
    main();
});

function input() {
    return lines[index++];
}

function main() {
    const n = parseInt(input());
    console.log(input().split(' ').reduce((acc, year) => acc + +year, 0) / n)
}