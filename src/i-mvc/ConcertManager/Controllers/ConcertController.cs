namespace ConcertManager.Controllers;

using Microsoft.AspNetCore.Mvc;
using ConcertManager.Services;
using ConcertManager.Models;

public class ConcertController : Controller
{
    private readonly IConcertService _concertService;

    public ConcertController(IConcertService concertService)
    {
        _concertService = concertService;
    }

    [HttpGet]
    public IActionResult Index()
    {
        var concerts = _concertService.GetAllConcerts();
        return View(concerts);
    }

    [HttpGet]
    [Route("[controller]/{id}")]
    public IActionResult Details(int id)
    {
        var concert = _concertService.GetConcertById(id);
        if (concert == null)
        {
            return NotFound();
        }
        return View(concert);
    }
} 