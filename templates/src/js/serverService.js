import { callbackServerServiceResponse } from "./callbacks.js";
export function serverService(sortType = "asc", path) {
    const xhr = new XMLHttpRequest();
    let url = new URL(window.location.href + 'dir');
    url.searchParams.set('root', path);
    url.searchParams.set('sort', sortType);
    xhr.open('GET', url.toString());
    xhr.responseType = 'json';
    xhr.send();
    xhr.onload = function () {
        callbackServerServiceResponse(path, xhr);
    };
    xhr.onerror = function () {
        alert("Запрос не удался");
    };
}
