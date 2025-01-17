using Microsoft.AspNetCore.Mvc;
using ConcertManager.Services;
using ConcertManager.Models;

public class TicketController : Controller
{
    private readonly ITicketService _ticketService;

    public TicketController(ITicketService ticketService)
    {
        _ticketService = ticketService;
    }

    [HttpGet]
    public IActionResult Purchase(int concertId)
    {
        var concert = _ticketService.GetConcertById(concertId);
        if (concert == null || concert.AvailableTickets <= 0)
        {
            return NotFound();
        }
        return View(new Ticket { ConcertId = concertId });
    }

    [HttpPost]
    public IActionResult Purchase(Ticket ticket)
    {
        if (ModelState.IsValid)
        {
            _ticketService.PurchaseTicket(ticket);
            return RedirectToAction("Confirmation", new { id = ticket.Id });
        }
        return View(ticket);
    }
} 