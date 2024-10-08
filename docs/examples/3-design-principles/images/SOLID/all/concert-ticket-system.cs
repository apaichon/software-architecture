using System;
using System.Collections.Generic;

// Single Responsibility Principle: Each class has a single responsibility
public interface ITicket
{
    string GetTicketInfo();
}

public interface IPriceable
{
    decimal CalculatePrice();
}

public interface IReservable
{
    bool Reserve();
}

// Open/Closed Principle: Open for extension, closed for modification
public abstract class Ticket : ITicket, IPriceable, IReservable
{
    public string EventName { get; set; }
    public DateTime EventDate { get; set; }
    public string Seat { get; set; }

    public abstract decimal CalculatePrice();
    public abstract bool Reserve();

    public virtual string GetTicketInfo()
    {
        return $"Event: {EventName}, Date: {EventDate}, Seat: {Seat}";
    }
}

// Liskov Substitution Principle: Subtypes must be substitutable for their base types
public class ConcertTicket : Ticket
{
    public string Artist { get; set; }

    public override decimal CalculatePrice()
    {
        // Implementation for concert ticket pricing
        return 50.00m; // Base price for simplicity
    }

    public override bool Reserve()
    {
        // Implementation for concert ticket reservation
        Console.WriteLine("Concert ticket reserved.");
        return true;
    }

    public override string GetTicketInfo()
    {
        return base.GetTicketInfo() + $", Artist: {Artist}";
    }
}

// Interface Segregation Principle: Clients should not be forced to depend on interfaces they do not use
public interface ITicketPrinter
{
    void PrintTicket(ITicket ticket);
}

public class TicketPrinter : ITicketPrinter
{
    public void PrintTicket(ITicket ticket)
    {
        Console.WriteLine(ticket.GetTicketInfo());
    }
}

// Dependency Inversion Principle: High-level modules should not depend on low-level modules. Both should depend on abstractions.
public class TicketManager
{
    private readonly List<ITicket> tickets;
    private readonly ITicketPrinter printer;

    public TicketManager(ITicketPrinter printer)
    {
        this.tickets = new List<ITicket>();
        this.printer = printer;
    }

    public void AddTicket(ITicket ticket)
    {
        tickets.Add(ticket);
    }

    public void PrintAllTickets()
    {
        foreach (var ticket in tickets)
        {
            printer.PrintTicket(ticket);
        }
    }
}

// Usage example
public class Program
{
    public static void Main()
    {
        var concertTicket = new ConcertTicket
        {
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(30),
            Seat = "A1",
            Artist = "Rock Band"
        };

        var ticketManager = new TicketManager(new TicketPrinter());
        ticketManager.AddTicket(concertTicket);
        ticketManager.PrintAllTickets();
    }
}
