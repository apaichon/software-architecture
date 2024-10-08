using System;
using System.Collections.Generic;

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

// Base Ticket class implementing common functionality
public abstract class Ticket : ITicket, IPriceable, IReservable
{
    public string EventName { get; set; }
    public DateTime EventDate { get; set; }
    public string Seat { get; set; }
    protected decimal BasePrice { get; set; }

    // DRY: Common implementation for GetTicketInfo
    public virtual string GetTicketInfo()
    {
        return $"Event: {EventName}, Date: {EventDate}, Seat: {Seat}";
    }

    // DRY: Common implementation for Reserve
    public virtual bool Reserve()
    {
        Console.WriteLine($"{GetType().Name} reserved.");
        return true;
    }

    // Abstract method to be implemented by derived classes
    public abstract decimal CalculatePrice();
}

public class ConcertTicket : Ticket
{
    public string Artist { get; set; }

    public ConcertTicket()
    {
        BasePrice = 50.00m;
    }

    public override decimal CalculatePrice()
    {
        return BasePrice;
    }

    public override string GetTicketInfo()
    {
        return base.GetTicketInfo() + $", Artist: {Artist}";
    }
}

// New ticket type
public class SportsTicket : Ticket
{
    public string Team { get; set; }

    public SportsTicket()
    {
        BasePrice = 40.00m;
    }

    public override decimal CalculatePrice()
    {
        return BasePrice;
    }

    public override string GetTicketInfo()
    {
        return base.GetTicketInfo() + $", Team: {Team}";
    }
}

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

    // New function to calculate total revenue
    public decimal CalculateTotalRevenue()
    {
        decimal totalRevenue = 0;
        foreach (var ticket in tickets)
        {
            if (ticket is IPriceable priceable)
            {
                totalRevenue += priceable.CalculatePrice();
            }
        }
        return totalRevenue;
    }
}

// Usage example
public class Program
{
    public static void Main()
    {
        var ticketManager = new TicketManager(new TicketPrinter());

        var concertTicket = new ConcertTicket
        {
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(30),
            Seat = "A1",
            Artist = "Rock Band"
        };

        var sportsTicket = new SportsTicket
        {
            EventName = "Football Match",
            EventDate = DateTime.Now.AddDays(15),
            Seat = "B5",
            Team = "Local Team"
        };

        ticketManager.AddTicket(concertTicket);
        ticketManager.AddTicket(sportsTicket);

        ticketManager.PrintAllTickets();

        Console.WriteLine($"Total Revenue: ${ticketManager.CalculateTotalRevenue()}");
    }
}
