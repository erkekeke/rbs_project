import { serverService } from "./serverService.js";
var path = '/home/erke/Desktop/MyFolder/';
export const main = function main(sortType = 'asc') {
    serverService(sortType, path);
};
document.addEventListener("DOMContentLoaded", function () {
    main();
});
