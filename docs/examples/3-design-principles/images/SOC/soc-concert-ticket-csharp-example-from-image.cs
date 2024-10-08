using System;

// Violating Separation of Concerns
class ConcertTicket
{
    public void BookTicket()
    {
        // Check availability
        Console.WriteLine("Checking ticket availability...");

        // Process payment
        Console.WriteLine("Processing payment...");

        // Update database
        Console.WriteLine("Updating database...");

        // Generate ticket
        Console.WriteLine("Generating ticket...");

        // Send confirmation email
        Console.WriteLine("Sending confirmation email...");
    }

    public void PrintTicket()
    {
        Console.WriteLine("Printing ticket...");
        // Printing logic here
    }

    public void CancelTicket()
    {
        Console.WriteLine("Cancelling ticket...");
        // Cancellation logic here
    }
}

// Following Separation of Concerns
class TicketBookingService
{
    public void BookTicket()
    {
        Console.WriteLine("Booking ticket...");
        // Ticket booking logic here
    }
}

class PaymentProcessor
{
    public void ProcessPayment()
    {
        Console.WriteLine("Processing payment...");
        // Payment processing logic here
    }
}

class TicketGenerator
{
    public Ticket GenerateTicket()
    {
        Console.WriteLine("Generating ticket...");
        // Ticket generation logic here
        return new Ticket();
    }
}

class EmailService
{
    public void SendConfirmation()
    {
        Console.WriteLine("Sending confirmation email...");
        // Email sending logic here
    }
}

// Helper class for the SoC example
class Ticket
{
    // Ticket properties and methods
}

class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Violating Separation of Concerns:");
        var badTicket = new ConcertTicket();
        badTicket.BookTicket();
        badTicket.PrintTicket();
        badTicket.CancelTicket();

        Console.WriteLine("\nFollowing Separation of Concerns:");
        var bookingService = new TicketBookingService();
        var paymentProcessor = new PaymentProcessor();
        var ticketGenerator = new TicketGenerator();
        var emailService = new EmailService();

        bookingService.BookTicket();
        paymentProcessor.ProcessPayment();
        var ticket = ticketGenerator.GenerateTicket();
        emailService.SendConfirmation();
    }
}
