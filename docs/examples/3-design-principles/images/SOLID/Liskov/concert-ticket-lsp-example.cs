using System;
using System.Collections.Generic;

public class ConcertTicket
{
    public virtual string GetTicketInfo() => "Standard Concert Ticket";
    public virtual decimal CalculatePrice() => 50m;
    public virtual void PrintTicket() => Console.WriteLine($"Printing: {GetTicketInfo()}");
}

public class VIPTicket : ConcertTicket
{
    public override string GetTicketInfo() => "VIP Concert Ticket";
    public override decimal CalculatePrice() => 100m;
    // VIPTicket adheres to LSP by not changing the behavior of PrintTicket
}

public class EarlyBirdTicket : ConcertTicket
{
    public override string GetTicketInfo() => "Early Bird Concert Ticket";
    public override decimal CalculatePrice() => 40m;
    // EarlyBirdTicket adheres to LSP by not changing the behavior of PrintTicket
}

// This class violates LSP
public class FreePassTicket : ConcertTicket
{
    public override string GetTicketInfo() => "Free Pass Ticket";
    public override decimal CalculatePrice() => 0m;
    public override void PrintTicket()
    {
        throw new InvalidOperationException("Free passes cannot be printed.");
    }
}

public class TicketManager
{
    public void ProcessTickets(List<ConcertTicket> tickets)
    {
        foreach (var ticket in tickets)
        {
            Console.WriteLine($"Ticket Info: {ticket.GetTicketInfo()}");
            Console.WriteLine($"Price: ${ticket.CalculatePrice()}");
            try
            {
                ticket.PrintTicket();
            }
            catch (InvalidOperationException ex)
            {
                Console.WriteLine($"Error: {ex.Message}");
            }
            Console.WriteLine();
        }
    }
}

class Program
{
    static void Main(string[] args)
    {
        var tickets = new List<ConcertTicket>
        {
            new ConcertTicket(),
            new VIPTicket(),
            new EarlyBirdTicket(),
            new FreePassTicket()
        };

        var manager = new TicketManager();
        manager.ProcessTickets(tickets);
    }
}
