# FoodOrdering
Сервис заказа еды
![image](https://github.com/MikhailMishutkin/FoodOrdering/assets/74837722/202479bd-3332-4f2b-adcb-215520daffe1)

Требования:
1. Брокеры сообщений: NATS или Kafka
2. База данных: PostgreSQL, но на каждом сервере разные драйверы: pgx, gorm и чистая
3. Логирование
4. Docker-контейнеризация всех серверов
5. API внешняя, скачана в данном проекте - папка pkg. Gateways, сгенерированые на основе этих данных, в папке proto
6. Пока не прикручивал базы данных, складываю пока в мапу
7. Данные генерирую, пакет gen в папке microservices
8. Установил NATS, тестовое прокидываю сообщения между office(customer) и restaurant
