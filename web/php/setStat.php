<?php
    ini_set('display_errors', 1);
    ini_set('display_startup_errors', 1);
    error_reporting(E_ALL);

    $jsonData = file_get_contents('php://input');
    if ($jsonData) {
        echo "Данные получены ";
    } else {
        echo "Данные не получены ";
    }

    $data = json_decode($jsonData, true);

    // Подключение к базе данных
    $servername = "localhost";
    $username = "root";
    $password = "password";
    $dbname = "statistics_db";

    if ($data !== null) {
        $fileName = $data['fileName'];
        $fileSize = $data['fileSize'];
        $elapsedTime = $data['elapsedTime']; // Добавлено для получения elapsedTime из JSON
        $currentDate = $data['currentDate'];

        $conn = new mysqli($servername, $username, $password, $dbname);

        // Проверка соединения
        if ($conn->connect_error) {
            echo "Ошибка подключения: " . $conn->connect_error;
        }

        // Используйте подготовленные запросы, чтобы избежать SQL-инъекций
        $sql = "INSERT INTO statistics (root, size, elapsedTime, currentDate) VALUES (?, ?, ?, ?)";

        // Подготовка запроса
        $stmt = $conn->prepare($sql);

        // Привязка параметров
        $stmt->bind_param("sdds", $fileName, $fileSize, $elapsedTime, $currentDate);

        // Выполнение запроса
        if ($stmt->execute()) {
            echo "Данные успешно добавлены";
        } else {
            echo "Ошибка: " . $stmt->error;
        }

        // Закрытие запроса и соединения
        $stmt->close();
        $conn->close();
    }
?>
