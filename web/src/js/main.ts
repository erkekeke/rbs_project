import { serverService } from "./serverService.js";

var path = '/home/';

export const main = function main(sortType: string = 'asc'): void {
  serverService(sortType, path);
};

document.addEventListener("DOMContentLoaded", function() {
  main();
});
