namespace ConcertManager.Data;

using ConcertManager.Models;

public static class DbSeeder
{
    public static void Initialize(ApplicationDbContext context)
    {
        if (context.Concerts.Any())
        {
            return;   // DB has been seeded
        }

        var concerts = new Concert[]
        {
            new Concert
            {
                Name = "Rock Festival 2024",
                Date = DateTime.Now.AddMonths(2),
                Venue = "Central Park",
                AvailableTickets = 1000,
                TicketPrice = 99.99m
            },
            new Concert
            {
                Name = "Jazz Night",
                Date = DateTime.Now.AddMonths(1),
                Venue = "Blue Note Club",
                AvailableTickets = 200,
                TicketPrice = 149.99m
            },
            new Concert
            {
                Name = "Classical Symphony",
                Date = DateTime.Now.AddMonths(3),
                Venue = "Concert Hall",
                AvailableTickets = 500,
                TicketPrice = 199.99m
            }
        };

        context.Concerts.AddRange(concerts);
        context.SaveChanges();
    }
} 