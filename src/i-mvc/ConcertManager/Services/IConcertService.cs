namespace ConcertManager.Services;

using System.Collections.Generic;
using ConcertManager.Models;

public interface IConcertService
{
    IEnumerable<Concert> GetAllConcerts();
    Concert GetConcertById(int id);
}

public interface ITicketService
{
    Concert GetConcertById(int id);
    void PurchaseTicket(Ticket ticket);
} 