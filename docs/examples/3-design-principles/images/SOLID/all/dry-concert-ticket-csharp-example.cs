using System;
using System.Collections.Generic;

// DRY Violation Example
namespace DRYViolation
{
    public class RegularTicket
    {
        public string EventName { get; set; }
        public DateTime EventDate { get; set; }
        public decimal BasePrice { get; set; }

        public decimal CalculatePrice()
        {
            decimal price = BasePrice;
            
            // Apply early bird discount
            if (DateTime.Now < EventDate.AddDays(-30))
            {
                price *= 0.9m; // 10% discount
            }

            // Apply tax
            price *= 1.1m; // 10% tax

            return Math.Round(price, 2);
        }

        public string GenerateTicketInfo()
        {
            return $"Regular Ticket - {EventName} on {EventDate.ToShortDateString()}\nPrice: ${CalculatePrice()}";
        }
    }

    public class VIPTicket
    {
        public string EventName { get; set; }
        public DateTime EventDate { get; set; }
        public decimal BasePrice { get; set; }

        public decimal CalculatePrice()
        {
            decimal price = BasePrice * 1.5m; // VIP premium
            
            // Apply early bird discount
            if (DateTime.Now < EventDate.AddDays(-30))
            {
                price *= 0.9m; // 10% discount
            }

            // Apply tax
            price *= 1.1m; // 10% tax

            return Math.Round(price, 2);
        }

        public string GenerateTicketInfo()
        {
            return $"VIP Ticket - {EventName} on {EventDate.ToShortDateString()}\nPrice: ${CalculatePrice()}";
        }
    }
}

// DRY Adherence Example
namespace DRYAdherence
{
    public abstract class Ticket
    {
        public string EventName { get; set; }
        public DateTime EventDate { get; set; }
        public decimal BasePrice { get; set; }

        protected abstract decimal ApplyPremium(decimal price);

        public decimal CalculatePrice()
        {
            decimal price = ApplyPremium(BasePrice);
            
            // Apply early bird discount
            if (DateTime.Now < EventDate.AddDays(-30))
            {
                price *= 0.9m; // 10% discount
            }

            // Apply tax
            price *= 1.1m; // 10% tax

            return Math.Round(price, 2);
        }

        public string GenerateTicketInfo()
        {
            return $"{GetType().Name} - {EventName} on {EventDate.ToShortDateString()}\nPrice: ${CalculatePrice()}";
        }
    }

    public class RegularTicket : Ticket
    {
        protected override decimal ApplyPremium(decimal price)
        {
            return price; // No premium for regular tickets
        }
    }

    public class VIPTicket : Ticket
    {
        protected override decimal ApplyPremium(decimal price)
        {
            return price * 1.5m; // 50% premium for VIP tickets
        }
    }
}

// Example usage
class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("DRY Violation Example:");
        var regularTicketViolation = new DRYViolation.RegularTicket
        {
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(45),
            BasePrice = 100
        };
        var vipTicketViolation = new DRYViolation.VIPTicket
        {
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(45),
            BasePrice = 100
        };
        Console.WriteLine(regularTicketViolation.GenerateTicketInfo());
        Console.WriteLine(vipTicketViolation.GenerateTicketInfo());

        Console.WriteLine("\nDRY Adherence Example:");
        var regularTicketAdherence = new DRYAdherence.RegularTicket
        {
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(45),
            BasePrice = 100
        };
        var vipTicketAdherence = new DRYAdherence.VIPTicket
        {
            EventName = "Rock Concert",
            EventDate = DateTime.Now.AddDays(45),
            BasePrice = 100
        };
        Console.WriteLine(regularTicketAdherence.GenerateTicketInfo());
        Console.WriteLine(vipTicketAdherence.GenerateTicketInfo());
    }
}
