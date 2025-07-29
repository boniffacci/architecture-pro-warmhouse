# Project_template

Это шаблон для решения проектной работы. Структура этого файла повторяет структуру заданий. Заполняйте его по мере работы над решением.

# Задание 1. Анализ и планирование

<aside>

Чтобы составить документ с описанием текущей архитектуры приложения, можно часть информации взять из описания компании и условия задания. Это нормально.

</aside

### 1. Описание функциональности монолитного приложения

**Управление отоплением:**

- Пользователи могут удалённо включать/выключать отопление в своих домах.

**Мониторинг температуры:**

- Система получает данные о температуре с датчиков, установленных в домах.
- Пользователи могут просматривать текущую температуру в своих домах через веб-интерфейс.

### 2. Анализ архитектуры монолитного приложения

Перечислите здесь основные особенности текущего приложения: какой язык программирования используется, какая база данных, как организовано взаимодействие между компонентами и так далее.

* Язык програмирования: Go
* База данных: PostgreSQL
* Архитектура: Монолитная, все компоненты системы (обработка запросов, бизнес-логика, работа с данными) находятся в рамках одного приложения.
* Взаимодействие: Синхронное, запросы обрабатываются последовательно.
* Масштабируемость: Ограничена, так как монолит сложно масштабировать по частям.
* Развертывание: Требует остановки всего приложения.

### 3. Определение доменов и границы контекстов

Опишите здесь домены, которые вы выделили.

**Домены**

- Управление устройствами
  - Изменение теммпературы в доме
  - Отключение/включение отопления

- Мониторинг темературы
  - Получение данных о текущей температуре

- Установка оборудования
  - Подключение датчиков и системы отопления в доме клиента с помощью специалиста

- Инфраструктура системы
  - Взаимодействие компонентов, протоколы связи

**Контексты**  

- Управление устройствами
  - Сервер отправляет команды на устройства
  - Все управление идет централизованно от сервера

- Мониторинг темературы
  - Данные о температуре запрашиваются синхронно от сервера к датчику
  - Нет собыйтийного или реактивного взаимодействия

- Установка и подключение оборудования
  - Подключение производится только специалистом
  - Пользователь не может добавить свой датчик

- Техническая арихтектура
  - Монолит на Go
  - СУБД: Postgres
  - Синхронное взаимодествие

### **4. Проблемы монолитного решения**

- Медленные релизы и развертывания
- Сложность поддержки и понимания
- Низкая отказоустойчивость
- Невозможность масштабировать части системы отдельно
- Невозможность быстрого добавления новых фитч

### 5. Визуализация контекста системы — диаграмма С4

```markdown
[Диаграмма PlantUML](https://www.plantuml.com/plantuml/png/PLBDQjn03BxxATGzRQ5UNthAeQI4qXIwi7H9Ju4m6Ilsm8mq7Sd9xjjNixPRkZt98xwVJv9lHLtHASNcNI0VfnvXLCroqtN5lUs6eEFqF0aMpwH8klESkXmTLKuhsiVmtDrzR0-rzR3_tYKdYiLQJtULST9ThkhD0OimRQlK1zWyMDbOhPZ2KM14fu46SDJ3Qy0t0SMKiJYT2aBlr6tUhng1B7IC5w5lyxUHcU2OqGSNW-YYJUZLb4RyeozdCU6BZI095u50Cu8jcvaqnmjaEoDDvvuJmlt2tpjo0oOaXHbNdVPC76rJmSDN1fTp37vEzC85_3nZNCTBJ46vs7DDzTVrwSj9jps-1ezYmiSbOgLzzZevjLTrF8ndCOzeuTknvbiuDebJA9Yvw3yRh0e59RDLvJKJI1Z8HQa1wcMWRJ-jAxg1w-Cyqi_hopJBBQwfordcv6Rs4WP32-2lAUIwlMQ16clDRZIRjvSrdDMKZ2-FzCDuAFh_Jz7S6cDAyJS0)
```

# Задание 2. Проектирование микросервисной архитектуры

В этом задании вам нужно предоставить только диаграммы в модели C4. Мы не просим вас отдельно описывать получившиеся микросервисы и то, как вы определили взаимодействия между компонентами To-Be системы. Если вы правильно подготовите диаграммы C4, они и так это покажут.

**Диаграмма контейнеров (Containers)**

```markdown
[Containers](https://www.plantuml.com/plantuml/png/RLFHJjjA47ttLypDomqIYL_k4n9NK4W12Qg0YVg8nlRWR6NlkhjZGbHLQllQXrRzW_u1gbgLLXNVuFnHP-tWG0GlPcRFdBDpzcJ7CLgkYZpuH-aahrA2ANFfjiBGucoOAPvMSUN89aOpQHucfWZB7BMV6JXESnM7U_yDnhurEJaE2tHCLbeNUpA1IfCTbfu_-1UsjxUX_bp_NdoexnRl5Xzz8nYJTKRtlSOcz2RowSdN3rkJEt05MeQfAMWJ7B63geIKhbH23b2dW1MR0baPxNeRGH3DHRwu65rpVuaIigZXs6zhppwMMfFjY2gdT0Qc90sbDMmIamjRT-MlzUtYlLpwH_sdll6jUvbdfj8fsddV8Qw43ar14256C4gCQv2z3NWJ121_dIlz6SLObZBrakBmnCGg9zWjIt-_CyA48OJdUSNS6Z8vWajZ8MbzoUM-ZsnmnEnxuixA4LbXJHDB9LoxuoCuGAOPpZtrWH6f5oQbuIlNg9kAXTRIwueSEs03IZDPZRdesDR5TUgbtOTLB0ojrI15oz--Q55KMYNekMUIAoYdSoTrtYrofH0SexSqxHxY2IcNa4Qhp4XdKdlVbmreEolHqJKbpJ8-NeEKBkNqOMR0LlaJo-p7_RJP0nbZTDHanZZEB4MdnxwAsDWbBQI2seGkRgYVL7ZRnUbB_RE-gMybJ7VrTmdLhoPStoHadpmWQ4W7W_-XJGXimSIH2vQLFv1db6v4EdLmDehEmmEBvVJqk7kpm2CyidiLmNU_gRQv0ZtsKv0dg34Zzs2DO0sppdrkLPQHlUS6JFoK2rQOrOJhxfv89j3_4uAbXDILd37bL13R-Me7Bt16cBfmPXMJdzENAckIt_eRx0XVLUH_0G00)
```

**Диаграмма компонентов (Components)**

```markdown
[API Gateway](https://www.plantuml.com/plantuml/png/PP51pn8n48Rl-oj6Jf08lNZoX0426524LOzcw8vijJkjxLIoCVvtM_m21FQowRjzdtdQdKR18AcprIl3sgQ6e1NnyQrI0S-Jav4s7LEae1qBiKoqwvItoAKpZj9OSrIpD-DTYJvlDwh3A1Hoz6tcEk-uT7pXLxDCGCCKXkZDQnYyswrXYK9dx0TvkNIWu8DhQF8ZblM29VJWdM61xov0aG3a1cAV9tGGw5UYA74mWZyLvESwRnXoNgO5Pos5pDel3hLQ1lJjfmtSlfGnguoq5C5mFfdX4wpgUdS7Vq1ZadPhcbmvOw1oYXp0BId8BsXDWqBnlxEudyGHXk-_rkfZsJvweWQNv2AwvlncjB2_98KsDz6Zw9RYrIeZOKw_ZQPLEbxkvA29CHYtu9DXolo_LRKd-tGJJ_etLzIQOZbmwJtkkmhUsPWNLzldKZKbRifl_WS0)
```

```markdown
[Device Hub](https://www.plantuml.com/plantuml/png/PP7FIZD14CJl-nJzvPGFJFRYoLCa4HKKGZQUfJFJIGRdpp3TQn3ntUt9BWfc3yDEBVMhedR1WaNwu9j_BXhVMuAZIEQRjYrucX-S7FjTprHCYa9HvYQ5DdkCrJDZiTxjskNrR5sbb-UdDY0B5PLUboda5AidLtwpL0AwI6LgwTqPUkXtLp1PdTz1BnEztQTwgZFqqHaKOZXf1U20HM3myUG_V3QWpq_2D9SaoIH_Qp5hl3BMemAZLBcF8NiA58LXSzTjmKNziCSpyW_Gf10msevAZHmxgG0hnnd570TYHz4or6sAHWqc5IIDNI_1Wzu9YXwfqEy46peubl9HkTk2ver720GIj2W8jG2hjQ_ShwRPaBzeUZc5mYBxyvWLi4_bXCLMHBEWQElF-GO0)
```

```markdown
[Scenario Engine](https://www.plantuml.com/plantuml/png/PP71JiCm38RlVGgpqv3OUk74QMWg0Z4G6aCSaPTuNQGqgH9d827UdRXT9rXxgZ_5dtydYyWOE3Msk31EsQG9zinjl2dBW9_psl0-RLEaeBnZSZnNlYbRYqvwPf6rDTjoUJrx4VJsj2eRZ4mXeu-bRrhlfASLVx7C1ZIEmZGgSXYChroTwokOl1u1z6IIqQDdQvpyLL-a4bC4YJ7JjCidD61Yto0Rxo0aIt5o2Ty5vEyqTih1r3M5bObCUMnMRNe20nBzEngEiFC1D1sC8g13Riw-CvjCgGve4x8NrphNS09YwWj9-d07oZjj9DxOX4fulrQdkkq03EJVqflxwXbG3QQVebYJ7QzsdkwKPRYThktir2Z8yNxYdu75WfpktkyN)
```

**Диаграмма кода (Code)**
Не работает диаграмма кода в онлайн плантюмл, вставил кодом
```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Code.puml

Container(scenarioEngine, "Scenario Engine", "Go")

Component(ruleEvaluator, "Rule Evaluator", "Go", "Evaluates user automation conditions")

Class "RuleEvaluator" {
    +evaluate(rule: Rule, event: Event): Boolean
}

Class "Rule" {
    +condition: Expression
}

Class "Expression" {
    +type: String
    +field: String
    +operator: String
    +value: String
}

Class "Event" {
    +deviceId: String
    +timestamp: Instant
    +payload: Map<String, Any>
}

RuleEvaluator --> Rule : uses
Rule --> Expression : has
RuleEvaluator --> Event : compares with

@enduml
```

```plantuml
@startuml
!include https://raw.githubusercontent.com/plantuml-stdlib/C4-PlantUML/master/C4_Code.puml

Container(deviceHub, "Device Hub", "Go")

Component(protocolAdapter, "Protocol Adapter", "Go", "Handles communication with external devices")

Interface "DeviceProtocol" {
    +SendCommand(cmd Command): error
    +ReceiveTelemetry(): Telemetry
}

Class "RESTAdapter" {
    +endpoint: string
    +SendCommand(cmd Command): error
    +ReceiveTelemetry(): Telemetry
}

protocolAdapter ..> DeviceProtocol
RESTAdapter ..|> DeviceProtocol

@enduml
```

# Задание 3. Разработка ER-диаграммы

```markdown
[ER - диаграмма](https://www.plantuml.com/plantuml/png/dPD1YlCm38NtEONNF_G2MPMVCZ2BMRLT1zKMhS6sWgMqbAPtdpWuuqum56OflOU-9yNPia2KmRi6WrYvgoDZL8z6gNzgO8oTDQfLn-FdRcu5y3ZLUuasdESEUh3kTwi7vXj5qrs0BqLwDWjZXrUhCLFCNBpboBt75IQIaAOLd0LauDAhB5A2V94Pt1B4f-9z48Eie-t5KgZjbkvNnYxmokpgVQi-onreiTSKv3-HGmWlmBt60D5Ind4kwwm_h9OJO82JGrEDS4277YNUSmPPwXMn-gd5UcG1tq_D7GWUfdeMqWIhkrNC8bp13KdvS0IIqinFTnmt6teiJwnLT0lygfGFb5GUJcdb4t8drBn8pyayZkh76D6155SULexVAXqnEPfCpePocbQTCM2SFSqMWqa_tJS0)
```

# Задание 4. Создание и документирование API

### 1. Тип API

Укажите, какой тип API вы будете использовать для взаимодействия микросервисов. Объясните своё решение.

## Используемые типы API

### REST API — для команд и запросов

REST используется там, где важна предсказуемость и быстрый ответ. Подходит для UI и операций управления:

- **User Service** — регистрация, аутентификация, управление аккаунтом
- **Device Management Service** — добавление, удаление и настройка устройств
- **API Gateway** — точка входа для фронтенда, маршрутизация REST-запросов

REST — это классический запрос-ответ, хорошо подходит для синхронного взаимодействия между фронтом и микросервисами.

### Event-driven — для событий и фоновой логики

Событийная архитектура применяется там, где важна масштабируемость, слабая связанность компонентов и реактивность:

- **Scenario Engine** — реагирует на события от устройств и запускает сценарии

Event-driven архитектура позволяет обрабатывать множество событий в реальном времени, не блокируя взаимодействие между сервисами.

### Итог

- **REST API** — для управления и UI: синхронно, понятно, быстро
- **Event-driven** — для телеметрии, автоматизации и масштабируемой фоновой логики


### 2. Документация API
Здесь приложите ссылки на документацию API для микросервисов, которые вы спроектировали в первой части проектной работы. Для документирования используйте Swagger/OpenAPI или AsyncAPI.

- [User Service API](./user-service.yaml)
- [Device Management Service API](./device-service.yaml)
- [Telemetry Service API](./telemetry-service.yaml)
- [Scenario Engine API](./scenario-service.yaml)

# Задание 5. Работа с docker и docker-compose

Перейдите в apps.

Там находится приложение-монолит для работы с датчиками температуры. В README.md описано как запустить решение.

Вам нужно:

1) сделать простое приложение temperature-api на любом удобном для вас языке программирования, которое при запросе /temperature?location= будет отдавать рандомное значение температуры.

Locations - название комнаты, sensorId - идентификатор названия комнаты

```
	// If no location is provided, use a default based on sensor ID
	if location == "" {
		switch sensorID {
		case "1":
			location = "Living Room"
		case "2":
			location = "Bedroom"
		case "3":
			location = "Kitchen"
		default:
			location = "Unknown"
		}
	}

	// If no sensor ID is provided, generate one based on location
	if sensorID == "" {
		switch location {
		case "Living Room":
			sensorID = "1"
		case "Bedroom":
			sensorID = "2"
		case "Kitchen":
			sensorID = "3"
		default:
			sensorID = "0"
		}
	}
```

2) Приложение следует упаковать в Docker и добавить в docker-compose. Порт по умолчанию должен быть 8081

3) Кроме того для smart_home приложения требуется база данных - добавьте в docker-compose файл настройки для запуска postgres с указанием скрипта инициализации ./smart_home/init.sql

Для проверки можно использовать Postman коллекцию smarthome-api.postman_collection.json и вызвать:

- Create Sensor
- Get All Sensors

Должно при каждом вызове отображаться разное значение температуры

Ревьюер будет проверять точно так же.


# **Задание 6. Разработка MVP**

Монолит будет взаимодествовать с сервисами через Rest API
- [MVP](./mvp/)