namespace ConcertManager.Data;

using Microsoft.EntityFrameworkCore;
using ConcertManager.Models;

public class ApplicationDbContext : DbContext
{
    public ApplicationDbContext(DbContextOptions<ApplicationDbContext> options)
        : base(options)
    {
    }

    public DbSet<Concert> Concerts { get; set; }
    public DbSet<Ticket> Tickets { get; set; }
} 