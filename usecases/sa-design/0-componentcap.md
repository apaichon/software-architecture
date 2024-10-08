| Component       | Payment Gateway | Concert Ticketing | Gov Bond Reservation | Online Lottery | Mixed Election |
|-----------------|-----------------|-------------------|----------------------|----------------|----------------|
| API             | 1,000           | 5,000             | 2,000                | 3,000          | 1,500          |
| Load Balancer   | 1,200           | 6,000             | 2,400                | 3,600          | 1,800          |
| Message Queue   | 800             | 4,000             | 1,600                | 2,500          | 1,200          |
| Cache Read      | 5,000           | 20,000            | 10,000               | 15,000         | 7,500          |
| Cache Write     | 1,000           | 5,000             | 2,000                | 3,000          | 1,500          |
| Database Read   | 2,000           | 8,000             | 4,000                | 6,000          | 3,000          |
| Database Write  | 500             | 2,000             | 1,000                | 1,500          | 750            |
| Core Process    | 700             | 3,500             | 1,400                | 2,100          | 1,050          |



| Use Case                  | Highest Traffic Load    | Average Traffic Load   |
|---------------------------|-------------------------|------------------------|
| Payment Gateway           | 5,000 TPS               | 500 TPS                |
| Concert Ticketing         | 20,000 TPS              | 1,000 TPS              |
| Gov Bond Reservation      | 10,000 TPS              | 800 TPS                |
| Online Lottery            | 15,000 TPS              | 2,000 TPS              |
| Mixed Election            | 8,000 TPS               | 500 TPS                |


se.sPj,