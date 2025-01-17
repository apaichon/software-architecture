namespace ConcertManager.Services;

using System;
using System.Linq;
using ConcertManager.Models;
using ConcertManager.Data;

public class TicketService : ITicketService
{
    private readonly ApplicationDbContext _context;

    public TicketService(ApplicationDbContext context)
    {
        _context = context;
    }

    public Concert GetConcertById(int id)
    {
        return _context.Concerts.FirstOrDefault(c => c.Id == id);
    }

    public void PurchaseTicket(Ticket ticket)
    {
        var concert = _context.Concerts.FirstOrDefault(c => c.Id == ticket.ConcertId);
        if (concert != null && concert.AvailableTickets > 0)
        {
            ticket.PurchaseDate = DateTime.Now;
            _context.Tickets.Add(ticket);
            concert.AvailableTickets--;
            _context.SaveChanges();
        }
        else
        {
            throw new InvalidOperationException("No tickets available for this concert.");
        }
    }
} 