// callback.ts
import { createTable } from "./fillingOutTheTable.js";
export function callbackServerServiceResponse(path, xhr) {
    if (xhr.status == 200 && xhr.response != null) {
        const responseObj = xhr.response;
        console.log(responseObj[0]);
        createTable(responseObj, path);
    }
    else {
        alert("Ошибка: " + xhr.status);
    }
}
