# CQRS

## Class diagram
```mermaid
classDiagram
    class Concert {
        +Guid Id
        +string Name
        +DateTime Date
        +string Venue
        +int AvailableTickets
        +decimal TicketPrice
    }

    class Ticket {
        +Guid Id
        +Guid ConcertId
        +string CustomerName
        +string CustomerEmail
        +DateTime PurchaseDate
    }

    class CreateConcertCommand {
        +string Name
        +DateTime Date
        +string Venue
        +int AvailableTickets
        +decimal TicketPrice
    }

    class PurchaseTicketCommand {
        +Guid ConcertId
        +string CustomerName
        +string CustomerEmail
    }

    class GetConcertQuery {
        +Guid Id
    }

    class GetAllConcertsQuery {
    }

    class ConcertDto {
        +Guid Id
        +string Name
        +DateTime Date
        +string Venue
        +int AvailableTickets
        +decimal TicketPrice
    }

    class CreateConcertCommandHandler {
        -ConcertDbContext _context
        +Handle(CreateConcertCommand, CancellationToken) Task~Guid~
    }

    class PurchaseTicketCommandHandler {
        -ConcertDbContext _context
        +Handle(PurchaseTicketCommand, CancellationToken) Task~Guid~
    }

    class GetConcertQueryHandler {
        -ConcertDbContext _context
        +Handle(GetConcertQuery, CancellationToken) Task~ConcertDto~
    }

    class GetAllConcertsQueryHandler {
        -ConcertDbContext _context
        +Handle(GetAllConcertsQuery, CancellationToken) Task~List~ConcertDto~~
    }

    class ConcertDbContext {
        +DbSet~Concert~ Concerts
        +DbSet~Ticket~ Tickets
    }

    class ConcertController {
        -IMediator _mediator
        +CreateConcert(CreateConcertCommand) Task~ActionResult~Guid~~
        +PurchaseTicket(PurchaseTicketCommand) Task~ActionResult~Guid~~
        +GetConcert(Guid) Task~ActionResult~ConcertDto~~
        +GetAllConcerts() Task~ActionResult~List~ConcertDto~~~
    }

    CreateConcertCommand <-- CreateConcertCommandHandler : handles
    PurchaseTicketCommand <-- PurchaseTicketCommandHandler : handles
    GetConcertQuery <-- GetConcertQueryHandler : handles
    GetAllConcertsQuery <-- GetAllConcertsQueryHandler : handles
    
    CreateConcertCommandHandler --> ConcertDbContext : uses
    PurchaseTicketCommandHandler --> ConcertDbContext : uses
    GetConcertQueryHandler --> ConcertDbContext : uses
    GetAllConcertsQueryHandler --> ConcertDbContext : uses
    
    ConcertDbContext --> Concert : manages
    ConcertDbContext --> Ticket : manages
    
    ConcertController --> CreateConcertCommand : sends
    ConcertController --> PurchaseTicketCommand : sends
    ConcertController --> GetConcertQuery : sends
    ConcertController --> GetAllConcertsQuery : sends
    
    GetConcertQueryHandler --> ConcertDto : returns
    GetAllConcertsQueryHandler --> ConcertDto : returns
```

## Flow
```mermaid
graph TD
    A[Start] --> B[ConcertController]
    B --> C{Request Type}
    
    C -->|Create Concert| D[CreateConcertCommand]
    C -->|Purchase Ticket| E[PurchaseTicketCommand]
    C -->|Get Concert| F[GetConcertQuery]
    C -->|Get All Concerts| G[GetAllConcertsQuery]
    
    D --> H[Mediator]
    E --> H
    F --> H
    G --> H
    
    H --> I{Handler Type}
    
    I -->|CreateConcertCommandHandler| J[Create Concert in DB]
    I -->|PurchaseTicketCommandHandler| K[Purchase Ticket in DB]
    I -->|GetConcertQueryHandler| L[Retrieve Concert from DB]
    I -->|GetAllConcertsQueryHandler| M[Retrieve All Concerts from DB]
    
    J --> N[Return Concert ID]
    K --> O[Return Ticket ID]
    L --> P[Return Concert DTO]
    M --> Q[Return List of Concert DTOs]
    
    N --> R[Controller Response]
    O --> R
    P --> R
    Q --> R
    
    R --> S[End]
    
    subgraph Database
    T[ConcertDbContext]
    end
    
    J --> T
    K --> T
    L --> T
    M --> T
```

## Code
```csharp
using System;
using System.Collections.Generic;
using System.Threading.Tasks;
using MediatR;
using Microsoft.EntityFrameworkCore;

// Domain Models
public class Concert
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    public DateTime Date { get; set; }
    public string Venue { get; set; }
    public int AvailableTickets { get; set; }
    public decimal TicketPrice { get; set; }
}

public class Ticket
{
    public Guid Id { get; set; }
    public Guid ConcertId { get; set; }
    public string CustomerName { get; set; }
    public string CustomerEmail { get; set; }
    public DateTime PurchaseDate { get; set; }
}

// Command Models
public class CreateConcertCommand : IRequest<Guid>
{
    public string Name { get; set; }
    public DateTime Date { get; set; }
    public string Venue { get; set; }
    public int AvailableTickets { get; set; }
    public decimal TicketPrice { get; set; }
}

public class PurchaseTicketCommand : IRequest<Guid>
{
    public Guid ConcertId { get; set; }
    public string CustomerName { get; set; }
    public string CustomerEmail { get; set; }
}

// Query Models
public class GetConcertQuery : IRequest<ConcertDto>
{
    public Guid Id { get; set; }
}

public class GetAllConcertsQuery : IRequest<List<ConcertDto>>
{
}

public class ConcertDto
{
    public Guid Id { get; set; }
    public string Name { get; set; }
    public DateTime Date { get; set; }
    public string Venue { get; set; }
    public int AvailableTickets { get; set; }
    public decimal TicketPrice { get; set; }
}

// Command Handlers
public class CreateConcertCommandHandler : IRequestHandler<CreateConcertCommand, Guid>
{
    private readonly ConcertDbContext _context;

    public CreateConcertCommandHandler(ConcertDbContext context)
    {
        _context = context;
    }

    public async Task<Guid> Handle(CreateConcertCommand request, CancellationToken cancellationToken)
    {
        var concert = new Concert
        {
            Id = Guid.NewGuid(),
            Name = request.Name,
            Date = request.Date,
            Venue = request.Venue,
            AvailableTickets = request.AvailableTickets,
            TicketPrice = request.TicketPrice
        };

        _context.Concerts.Add(concert);
        await _context.SaveChangesAsync(cancellationToken);

        return concert.Id;
    }
}

public class PurchaseTicketCommandHandler : IRequestHandler<PurchaseTicketCommand, Guid>
{
    private readonly ConcertDbContext _context;

    public PurchaseTicketCommandHandler(ConcertDbContext context)
    {
        _context = context;
    }

    public async Task<Guid> Handle(PurchaseTicketCommand request, CancellationToken cancellationToken)
    {
        var concert = await _context.Concerts.FindAsync(new object[] { request.ConcertId }, cancellationToken);

        if (concert == null || concert.AvailableTickets <= 0)
            throw new Exception("Concert not found or no tickets available");

        var ticket = new Ticket
        {
            Id = Guid.NewGuid(),
            ConcertId = request.ConcertId,
            CustomerName = request.CustomerName,
            CustomerEmail = request.CustomerEmail,
            PurchaseDate = DateTime.UtcNow
        };

        concert.AvailableTickets--;

        _context.Tickets.Add(ticket);
        await _context.SaveChangesAsync(cancellationToken);

        return ticket.Id;
    }
}

// Query Handlers
public class GetConcertQueryHandler : IRequestHandler<GetConcertQuery, ConcertDto>
{
    private readonly ConcertDbContext _context;

    public GetConcertQueryHandler(ConcertDbContext context)
    {
        _context = context;
    }

    public async Task<ConcertDto> Handle(GetConcertQuery request, CancellationToken cancellationToken)
    {
        var concert = await _context.Concerts.FindAsync(new object[] { request.Id }, cancellationToken);

        if (concert == null)
            return null;

        return new ConcertDto
        {
            Id = concert.Id,
            Name = concert.Name,
            Date = concert.Date,
            Venue = concert.Venue,
            AvailableTickets = concert.AvailableTickets,
            TicketPrice = concert.TicketPrice
        };
    }
}

public class GetAllConcertsQueryHandler : IRequestHandler<GetAllConcertsQuery, List<ConcertDto>>
{
    private readonly ConcertDbContext _context;

    public GetAllConcertsQueryHandler(ConcertDbContext context)
    {
        _context = context;
    }

    public async Task<List<ConcertDto>> Handle(GetAllConcertsQuery request, CancellationToken cancellationToken)
    {
        return await _context.Concerts
            .Select(c => new ConcertDto
            {
                Id = c.Id,
                Name = c.Name,
                Date = c.Date,
                Venue = c.Venue,
                AvailableTickets = c.AvailableTickets,
                TicketPrice = c.TicketPrice
            })
            .ToListAsync(cancellationToken);
    }
}

// Database Context
public class ConcertDbContext : DbContext
{
    public ConcertDbContext(DbContextOptions<ConcertDbContext> options) : base(options) { }

    public DbSet<Concert> Concerts { get; set; }
    public DbSet<Ticket> Tickets { get; set; }
}

// Example usage in a controller
public class ConcertController : ControllerBase
{
    private readonly IMediator _mediator;

    public ConcertController(IMediator mediator)
    {
        _mediator = mediator;
    }

    [HttpPost]
    public async Task<ActionResult<Guid>> CreateConcert([FromBody] CreateConcertCommand command)
    {
        var concertId = await _mediator.Send(command);
        return Ok(concertId);
    }

    [HttpPost("purchase")]
    public async Task<ActionResult<Guid>> PurchaseTicket([FromBody] PurchaseTicketCommand command)
    {
        var ticketId = await _mediator.Send(command);
        return Ok(ticketId);
    }

    [HttpGet("{id}")]
    public async Task<ActionResult<ConcertDto>> GetConcert(Guid id)
    {
        var concert = await _mediator.Send(new GetConcertQuery { Id = id });
        if (concert == null)
            return NotFound();
        return Ok(concert);
    }

    [HttpGet]
    public async Task<ActionResult<List<ConcertDto>>> GetAllConcerts()
    {
        var concerts = await _mediator.Send(new GetAllConcertsQuery());
        return Ok(concerts);
    }
}

```