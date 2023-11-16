import { serverService } from "./serverService.js";
export {};
export function dirClick(fileName: string, path: string): void {
  path += fileName + '/';
  serverService('asc', path);
}

export function createTable(data: any[], path: string): void {
  const tableBody = document.getElementById('table-body') as HTMLTableElement;
  const backButton: HTMLElement | null = document.getElementById('back-button');
  const ascButton: HTMLElement | null = document.getElementById('ascButton');
  const descButton: HTMLElement | null = document.getElementById('descButton');

  if (!tableBody || !backButton || !ascButton || !descButton) {
    return;
  }

  backButton.onclick = function() {
    if (path !== "") {
      path = path.substring(0, path.slice(0, -1).lastIndexOf('/') + 1);
      serverService('asc', path);
    } else {
      alert("Ошибка. Задана неверная директория");
    }
  };

  ascButton.onclick = function() {
    serverService('asc', path);
  };

  descButton.onclick = function() {
    serverService('desc', path);
  };

  tableBody.innerHTML = "";

  let i = 1;
  data.forEach((file) => {
    if (file.deepIndex == 0) {
      let row = tableBody.insertRow();
      let cell1 = row.insertCell(0);
      let cell2 = row.insertCell(1);
      let cell3 = row.insertCell(2);
      let cell4 = row.insertCell(3);

      i++
      cell1.innerHTML = i.toString();
      cell2.innerHTML = file.fileName;
      cell3.innerHTML = file.fileSize;
      cell4.innerHTML = file.fileType;

      if (file.fileType == "folder") {
        cell2.addEventListener('click', function() {
          dirClick(file.fileName, path);
        });
      }
    }
  });
}