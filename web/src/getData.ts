import { callbackServerServiceResponse } from "./callbacks";
import { Loader } from "./loader";

// getData() Отравить запрос на сервер, обработать его.
export function getData(sortType: string = "asc", path: string): void {
    Loader.show()
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