// callback.ts
import { createTable } from "./fillingOutTheTable.js";

export function callbackServerServiceResponse(path: string, xhr: XMLHttpRequest): void {
    if (xhr.status == 200 && xhr.response != null) {
        const responseObj = xhr.response;
        createTable(responseObj, path);
    } else {
        alert("Ошибка при получении запроса: " + xhr.status);
    }
}