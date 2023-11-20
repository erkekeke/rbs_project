import { callbackServerServiceResponse } from "./callbacks.js";

export function serverService(sortType: string = "asc", path: string): void {
    const xhr: XMLHttpRequest = new XMLHttpRequest();
    let url: URL = new URL(window.location.href + 'dir');
    url.searchParams.set('root', path);
    url.searchParams.set('sort', sortType);
    
    xhr.open('GET', url.toString());
  
    xhr.responseType = 'json';
  
    xhr.send();
    
    xhr.onload = function() {
      callbackServerServiceResponse(path, xhr);
    };
}