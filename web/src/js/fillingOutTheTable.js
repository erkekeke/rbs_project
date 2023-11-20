import { serverService } from "./serverService.js";
export function dirClick(fileName, path) {
    path += fileName + '/';
    serverService('asc', path);
}
function convertBytes(sizeInBytes) {
    const kilobyte = 1024;
    const megabyte = kilobyte * kilobyte;
    const gigabyte = megabyte * kilobyte;
    if (sizeInBytes < kilobyte) {
        return `${sizeInBytes} B`;
    }
    else if (sizeInBytes < megabyte) {
        return `${(sizeInBytes / kilobyte).toFixed(2)} KB`;
    }
    else if (sizeInBytes < gigabyte) {
        return `${(sizeInBytes / megabyte).toFixed(2)} MB`;
    }
    else {
        return `${(sizeInBytes / gigabyte).toFixed(2)} GB`;
    }
}
export function createTable(data, path) {
    const tableBody = document.getElementById('table-body');
    const backButton = document.getElementById('back-button');
    const ascButton = document.getElementById('ascButton');
    const descButton = document.getElementById('descButton');
    if (!tableBody || !backButton || !ascButton || !descButton) {
        return;
    }
    backButton.onclick = function () {
        if (path !== "") {
            path = path.substring(0, path.slice(0, -1).lastIndexOf('/') + 1);
            serverService('asc', path);
        }
        else {
            alert("Ошибка. Задана неверная директория");
        }
    };
    ascButton.onclick = function () {
        serverService('asc', path);
    };
    descButton.onclick = function () {
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
            cell1.innerHTML = i.toString();
            cell2.innerHTML = file.fileName;
            cell3.innerHTML = convertBytes(file.fileSize);
            cell4.innerHTML = file.fileType;
            i += 1;
            if (file.fileType == "folder") {
                cell2.addEventListener('click', function () {
                    dirClick(file.fileName, path);
                });
            }
        }
    });
}
