using System;
using System.Collections.Generic;

// Abstractions (high-level modules)
public interface ITicketRepository
{
    void SaveTicket(Ticket ticket);
    Ticket GetTicket(int id);
}

public interface INotificationService
{
    void SendNotification(string message, string recipient);
}

// Low-level modules
public class SqlTicketRepository : ITicketRepository
{
    public void SaveTicket(Ticket ticket)
    {
        Console.WriteLine($"Saving ticket to SQL database: {ticket.Id}");
        // Actual database saving logic would go here
    }

    public Ticket GetTicket(int id)
    {
        Console.WriteLine($"Retrieving ticket from SQL database: {id}");
        // Actual database retrieval logic would go here
        return new Ticket { Id = id, EventName = "Sample Event" };
    }
}

public class EmailNotificationService : INotificationService
{
    public void SendNotification(string message, string recipient)
    {
        Console.WriteLine($"Sending email to {recipient}: {message}");
        // Actual email sending logic would go here
    }
}

// Domain entities
public class Ticket
{
    public int Id { get; set; }
    public string EventName { get; set; }
    public DateTime EventDate { get; set; }
    public decimal Price { get; set; }
}

// High-level module
public class TicketService
{
    private readonly ITicketRepository _ticketRepository;
    private readonly INotificationService _notificationService;

    public TicketService(ITicketRepository ticketRepository, INotificationService notificationService)
    {
        _ticketRepository = ticketRepository;
        _notificationService = notificationService;
    }

    public void PurchaseTicket(Ticket ticket, string customerEmail)
    {
        _ticketRepository.SaveTicket(ticket);
        _notificationService.SendNotification(
            $"Your ticket for {ticket.EventName} has been purchased.",
            customerEmail);
    }

    public Ticket GetTicketInfo(int ticketId)
    {
        return _ticketRepository.GetTicket(ticketId);
    }
}

// Example usage
class Program
{
    static void Main(string[] args)
    {
        // Setup
        ITicketRepository ticketRepository = new SqlTicketRepository();
        INotificationService notificationService = new EmailNotificationService();
        TicketService ticketService = new TicketService(ticketRepository, notificationService);

        // Usage
        Ticket newTicket = new Ticket
        {
            Id = 1,
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(30),
            Price = 99.99m
        };

        ticketService.PurchaseTicket(newTicket, "customer@example.com");

        Ticket retrievedTicket = ticketService.GetTicketInfo(1);
        Console.WriteLine($"Retrieved ticket for event: {retrievedTicket.EventName}");
    }
}
