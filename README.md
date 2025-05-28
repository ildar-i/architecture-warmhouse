# Анализ и планирование

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут управлять отоплением в доме
- Система поддерживает только синхронные запросы
- Данные о температуре также получаются через запрос от сервера к датчику

**Мониторинг температуры:**

- Пользователи могут проверять температуру
- Система поддерживает только синхронные запросы
- Данные о температуре также получаются через запрос от сервера к датчику

### 2. Анализ архитектуры монолитного приложения

Архитектура приложения представляет из себя монолит на Go с СУБД Postgres. 
Всё синхронно. Никаких асинхронных вызовов, микросервисов и реактивного взаимодействия в системе нет. 
Всё управление идёт от сервера к датчику. 
Данные о температуре также получаются через запрос от сервера к датчику.
Самостоятельно подключить свой датчик к системе пользователь не может.

### 3. Определение доменов и границы контекстов

Опишите здесь домены, которые вы выделили.
Auth Service - Домен авторизации
Device Control Service - Домен по управлению девайсами
Monitoring Service - Домен системы мониторинга
Database - Домен для базы данных

### **4. Проблемы монолитного решения**

- Сложность масштабирования
- Деплой всего приложения при изменение каких-то частей
- Медленная работа, так как доступны только синхронные запросы
- Пользователи должны иметь возможность управлять не только отоплением, но и умными устройствами в доме.

### 5. Визуализация контекста системы — диаграмма С4

[Диаграмма контекста](/schemas%2Fcontext.puml)

![context-Warm_House_Context_Diagram.png](/schemas%2Fcontext-Warm_House_Context_Diagram.png)

# Проектирование микросервисной архитектуры

**Диаграмма контейнеров (Containers)**

[Containers](schemas%2Fcontainer.puml)

![container-Warm_House_Context_Diagram.png](/schemas%2Fcontainer-Warm_House_Context_Diagram.png)

**Диаграмма компонентов (Components)**

[Auth Service](schemas%2Fauth_component.puml)

![auth_component-Monitoring_Service_Component_Diagram.png](/schemas%2Fauth_component-Monitoring_Service_Component_Diagram.png)

[Device Control Service](schemas%2Fdevice_component.puml)

![device_component-Device_Control_Service_Component_Diagram.png](/schemas%2Fdevice_component-Device_Control_Service_Component_Diagram.png)

[Monitoring Service](schemas%2Fmonitoring_component.puml)

![monitoring_component-Monitoring_Service_Component_Diagram.png](/schemas%2Fmonitoring_component-Monitoring_Service_Component_Diagram.png)


**Диаграмма кода (Code)**

![auth_code-Auth_Service_Code_Diagram.png](/schemas%2Fauth_code-Auth_Service_Code_Diagram.png)Добавьте одну диаграмму или несколько.

![monitoring_code-Monitoring_Service_Code_Diagram.png](/schemas%2Fmonitoring_code-Monitoring_Service_Code_Diagram.png)

# Разработка ER-диаграммы

![auth_code_er-Auth_Service_ER_Diagram.png](/schemas%2Fauth_code_er-Auth_Service_ER_Diagram.png)
![device_code_er-Device_Service_ER_Diagram.png](/schemas%2Fdevice_code_er-Device_Service_ER_Diagram.png)
![monitoring_code_er-Monitoring_Service_ER_Diagram.png](/schemas%2Fmonitoring_code_er-Monitoring_Service_ER_Diagram.png)

# Создание и документирование API

### 1. Тип API

Укажите, какой тип API вы будете использовать для взаимодействия микросервисов. Объясните своё решение.
- Гибкость и лёгкость использования. Echo подходит для проектов, требующих простоты и удобства работы.
- Поддержка масштабирования. Фреймворк предназначен для работы с большим количеством пользователей и может расширяться по мере роста приложения.
- Поддержка HTTP/2. Это повышает скорость и эффективность. 
- Множество встроенных промежуточных программ. Разработчики могут определять свои собственные, которые можно устанавливать на уровне root, группы или маршрута.

### Документация API
http://127.0.0.1:8081/swagger/index.html#/

# Работа с docker и docker-compose

Запустить приложение можно через команду **docker-compose up --build**

