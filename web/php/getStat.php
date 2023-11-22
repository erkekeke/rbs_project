<?php
    ini_set('display_errors', 1);
    ini_set('display_startup_errors', 1);
    error_reporting(E_ALL);

    // Подключение к базе данных
    $servername = "localhost";
    $username = "root";
    $password = "password";
    $dbname = "statistics_db";

    $conn = new mysqli($servername, $username, $password, $dbname);

    // Проверка соединения
    if ($conn->connect_error) {
        die("Ошибка подключения: " . $conn->connect_error);
    }

    // Запрос к таблице
    $sql = "SELECT root, size, elapsedTime, currentDate FROM statistics ORDER BY elapsedTime ASC";
    $result = $conn->query($sql);

    // Создание массивов для данных Chart.js
    $elapsedTimeData = [];
    $currentDateData = [];
    $sizeData = [];
    $rootData = [];

    // Обработка результатов запроса
    if ($result->num_rows > 0) {
        while($row = $result->fetch_assoc()) {
            $rootData[] = $row["root"];
            $elapsedTimeData[] = $row["elapsedTime"];
            $sizeData[] = $row["size"];
            $currentDateData[] = $row["currentDate"];
        }
    }

    // Закрытие соединения
    $conn->close();

    // Формирование массива данных для передачи в JavaScript
    $data = [
        'elapsedTime' => $elapsedTimeData,
        'size' => $sizeData,
    ];
?>



<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Пример таблицы в PHP с Chart.js</title>
    <!-- Подключение библиотеки Chart.js -->
    <script src="https://cdn.jsdelivr.net/npm/chart.js"></script>
    <style>
        html {
            display: grid;            
            justify-content: center;
            place-self: center;
        }

        #elapsedTime-size {
            max-width: 800px;
            max-height: 500px;
        }

        .table-container {
            max-height: 300px;
            overflow: auto;
        }

        table {
            border-spacing: 0;
            border-collapse: collapse;
            border: 3px solid #ffffff;
            text-align:center;
        }

        th {
            background-color: #4682B4;
            color: white;
        }

        td, th {
            width: 200px;
            position: relative; /* Необходимо для позиционирования псевдоэлемента */
            padding: 8px; /* Пример отступов внутри ячейки */
        }
        td::after {
            content: "";
            position: absolute;
            bottom: 0;
            left: 0;
            width: 100%;
            height: 2px; /* Высота границы */
            background-color: white; /* Цвет границы */
        }
    </style>
</head>
<body>
    <!-- Создание элемента canvas для графика -->
    <canvas id="elapsedTime-size" width="800" height="400"></canvas>

    <script>
        // Получение данных из PHP
        var data = <?php echo json_encode($data); ?>;
        var size = data.size;
        var elapsedTime = data.elapsedTime;
        
        // Получение элемента canvas
        var ctx = document.getElementById('elapsedTime-size').getContext('2d');

        // Создание графика
        var myChart = new Chart(ctx, {
            type: 'line', // Используем scatter plot для показа зависимости
            data: {
                labels: [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20],
                datasets: [{
                    label: 'Зависимость затраченного времени от размера директории',
                    data: data.size,
                }]
            },
            options: {
                scales: {
                    x: {
                        type: 'linear',
                        position: 'bottom',
                        title: {
                            display: true,
                            text: 'Затраченное время (секунды)'
                        }
                    },
                    y: {
                        type: 'linear',
                        position: 'left',
                        title: {
                            display: true,
                            text: 'Размер директории (байты)'
                        }
                    }
                }
            }
        });
    </script>

    <!-- Вывод таблицы -->
    <div class="table-container">
        <table border=1>
            <?php
                // Вывод данных в таблицу
                echo "<thead>";
                echo "<tr>";
                echo "<th>Имя директории</th>";
                echo "<th>Затраченное время (с.)</th>";
                echo "<th>Размер файла(байты)</th>";
                echo "<th>Дата сканирования</th>";
                echo "</tr>";
                echo "</thead>";

                echo "<tbody>";
                foreach ($rootData as $index => $root) {
                    echo "<tr>";
                    echo "<td>$root</td>";
                    echo "<td>{$elapsedTimeData[$index]}</td>";
                    echo "<td>{$sizeData[$index]}</td>";
                    echo "<td>{$currentDateData[$index]}</td>";
                    echo "</tr>";
                }
                echo "</tbody>";
            ?>
        </table>
    </div>
</body>
</html>