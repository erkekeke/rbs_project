import { getDir } from "./getDir";

var path = '/home/';

export const main = function main(sortType: string = 'asc'): void {
  getDir(sortType, path);
};

document.addEventListener("DOMContentLoaded", function() {
  main();
});
