import { getData } from "./getData";

var path = '/home/';

export const main = function main(sortType: string = 'asc'): void {
  getData(sortType, path);
};

document.addEventListener("DOMContentLoaded", function() {
  main();
});
