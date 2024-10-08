using System;

// Violating KISS
class TicketPriceCalculator
{
    public decimal CalculatePrice(
        Ticket ticket, DateTime purchaseDate,
        bool isVIP, bool isEarlyBird,
        int customerLoyaltyPoints)
    {
        decimal price = ticket.BasePrice;
        if (isVIP) price *= 1.5m;
        if (isEarlyBird) price *= 0.9m;
        price -= (customerLoyaltyPoints / 100);
        if (purchaseDate.DayOfWeek == DayOfWeek.Tuesday)
            price *= 0.95m;
        return Math.Max(price, 0);
    }
}

// Following KISS
abstract class Ticket
{
    public decimal BasePrice { get; set; }
    public abstract decimal CalculatePrice();
}

class RegularTicket : Ticket
{
    public override decimal CalculatePrice()
    {
        return BasePrice;
    }
}

class VIPTicket : Ticket
{
    public override decimal CalculatePrice()
    {
        return BasePrice * 1.5m;
    }
}

// Example usage
class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Violating KISS:");
        var calculator = new TicketPriceCalculator();
        var ticket = new Ticket() { BasePrice = 100 };
        var price = calculator.CalculatePrice(ticket, DateTime.Now, true, false, 50);
        Console.WriteLine($"Calculated Price: {price}");

        Console.WriteLine("\nFollowing KISS:");
        Ticket regularTicket = new RegularTicket() { BasePrice = 100 };
        Ticket vipTicket = new VIPTicket() { BasePrice = 100 };
        Console.WriteLine($"Regular Ticket Price: {regularTicket.CalculatePrice()}");
        Console.WriteLine($"VIP Ticket Price: {vipTicket.CalculatePrice()}");
    }
}

// Note: This Ticket class is only used for the "Violating KISS" example
class Ticket
{
    public decimal BasePrice { get; set; }
}
