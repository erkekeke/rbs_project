import { Loader } from "./loader";
import { createTable } from "./fillingOutTheTable";

// getData() Отравить запрос на сервер, обработать его.
export function getDir(sortType: string = "asc", path: string): void {
    Loader.show()
    const xhr: XMLHttpRequest = new XMLHttpRequest();
    let url: URL = new URL(window.location.href + 'dir');
    url.searchParams.set('root', path);
    url.searchParams.set('sort', sortType);
    
    xhr.open('GET', url.toString());
  
    xhr.responseType = 'json';
  
    xhr.send();
    
    xhr.onload = function() {
      if (xhr.status == 200) {
          const responseObj = xhr.response;
          createTable(responseObj, path);
      } else {
          alert("Ошибка при получении запроса: " + xhr.status);
          Loader.hide()
      }
    };
}