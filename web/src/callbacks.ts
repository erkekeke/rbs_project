// callback.ts
import { createTable } from "./fillingOutTheTable";
import { Loader } from "./loader";

export function callbackServerServiceResponse(path: string, xhr: XMLHttpRequest): void {
    if (xhr.status == 200) {
        const responseObj = xhr.response;
        createTable(responseObj, path);
    } else {
        alert("Ошибка при получении запроса: " + xhr.status);
        Loader.hide()
    }
}