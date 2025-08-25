### Процесс создания заказа

| Этап                             | Тип события  |                  Название |
|:---------------------------------|:------------:|--------------------------:|
| заказ создан                     |    domain    |         OrderCreatedEvent |
| резервирование товара успешно    |    domain    |  ItemsReserveSuccessEvent |
| резервирование товара не успешно |   failure    |   ItemsReserveFailedEvent |
| резервирование товаров отменено  | compensation | ItemsReserveCanceledEvent |
| оплата не прошла                 |   failure    |            PayFailedEvent |
| оплата прошла успешно            |    domain    |           PaySuccessEvent |
| доставка оформлена               |    domain    |      DeliveryOrderCreated |

### Процесс отмены заказа

| Этап                       | Тип события  |                      Название |
|:---------------------------|:------------:|------------------------------:|
| заказ отменен              |    domain    |            OrderCanceledEvent |
| деньги возвращены          | compensation |            MoneyReturnedEvent |
| товары возвращены на склад | compensation | ItemsReturnedToWarehouseEvent |








