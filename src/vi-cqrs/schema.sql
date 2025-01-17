CREATE TABLE IF NOT EXISTS concerts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    date DATETIME NOT NULL,
    venue TEXT NOT NULL,
    available_seats INTEGER NOT NULL,
    ticket_price REAL NOT NULL
);

CREATE TABLE IF NOT EXISTS tickets (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    concert_id INTEGER NOT NULL,
    student_name TEXT NOT NULL,
    student_class TEXT NOT NULL,
    purchase_date DATETIME NOT NULL,
    FOREIGN KEY (concert_id) REFERENCES concerts(id)
); 

select * from concerts;

SELECT id, name, datetime(date) as date, venue, available_seats, ticket_price 
		FROM concerts 
		WHERE available_seats >= 10;