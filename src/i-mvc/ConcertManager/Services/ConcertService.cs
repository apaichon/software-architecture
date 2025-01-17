namespace ConcertManager.Services;

using System.Collections.Generic;
using System.Linq;
using ConcertManager.Models;
using ConcertManager.Data;

public class ConcertService : IConcertService
{
    private readonly ApplicationDbContext _context;

    public ConcertService(ApplicationDbContext context)
    {
        _context = context;
    }

    public IEnumerable<Concert> GetAllConcerts()
    {
        return _context.Concerts.ToList();
    }

    public Concert GetConcertById(int id)
    {
        return _context.Concerts.FirstOrDefault(c => c.Id == id);
    }
} 