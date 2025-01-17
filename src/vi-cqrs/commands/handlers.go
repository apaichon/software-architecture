package commands

import (
	"database/sql"
	"errors"
	"fmt"
)

type CommandHandler struct {
	db *sql.DB
}

func NewCommandHandler(db *sql.DB) *CommandHandler {
	return &CommandHandler{db: db}
}

func (h *CommandHandler) HandlePurchaseTicket(cmd PurchaseTicketCommand) error {
	tx, err := h.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Check concert availability
	var availableSeats int
	err = tx.QueryRow("SELECT available_seats FROM concerts WHERE id = ?", cmd.ConcertID).Scan(&availableSeats)
	if err != nil {
		return err
	}

	if availableSeats <= 0 {
		return errors.New("no available seats")
	}

	// Create ticket
	_, err = tx.Exec(`
		INSERT INTO tickets (concert_id, student_name, student_class, purchase_date)
		VALUES (?, ?, ?, datetime('now'))`,
		cmd.ConcertID, cmd.StudentName, cmd.StudentClass)
	if err != nil {
		return err
	}

	// Update available seats
	_, err = tx.Exec("UPDATE concerts SET available_seats = available_seats - 1 WHERE id = ?", cmd.ConcertID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (h *CommandHandler) HandleCreateConcert(cmd CreateConcertCommand) error {
	fmt.Println("CreateConcertCommand", cmd)
	_, err := h.db.Exec(`
		INSERT INTO concerts (name, date, venue, available_seats, ticket_price)
		VALUES (?, datetime(?), ?, ?, ?)`,
		cmd.Name, cmd.Date.Format("2006-01-02 15:04:05"),
		cmd.Venue, cmd.AvailableSeats, cmd.TicketPrice)
	return err
}
