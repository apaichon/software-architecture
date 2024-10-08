using System;

// Bad Design: Violating ISP
namespace BadDesign
{
    public interface ITicket
    {
        void Reserve();
        void Cancel();
        void Print();
        void SendEmail();
    }

    public class StandardTicket : ITicket
    {
        public void Reserve() => Console.WriteLine("Standard ticket reserved.");
        public void Cancel() => Console.WriteLine("Standard ticket cancelled.");
        public void Print() => Console.WriteLine("Standard ticket printed.");
        public void SendEmail() => throw new NotImplementedException("Standard tickets don't support email.");
    }

    public class ElectronicTicket : ITicket
    {
        public void Reserve() => Console.WriteLine("E-ticket reserved.");
        public void Cancel() => Console.WriteLine("E-ticket cancelled.");
        public void Print() => throw new NotImplementedException("E-tickets can't be printed.");
        public void SendEmail() => Console.WriteLine("E-ticket sent via email.");
    }
}

// Good Design: Following ISP
namespace GoodDesign
{
    public interface IReservable
    {
        void Reserve();
        void Cancel();
    }

    public interface IPrintable
    {
        void Print();
    }

    public interface IEmailable
    {
        void SendEmail();
    }

    public class StandardTicket : IReservable, IPrintable
    {
        public void Reserve() => Console.WriteLine("Standard ticket reserved.");
        public void Cancel() => Console.WriteLine("Standard ticket cancelled.");
        public void Print() => Console.WriteLine("Standard ticket printed.");
    }

    public class ElectronicTicket : IReservable, IEmailable
    {
        public void Reserve() => Console.WriteLine("E-ticket reserved.");
        public void Cancel() => Console.WriteLine("E-ticket cancelled.");
        public void SendEmail() => Console.WriteLine("E-ticket sent via email.");
    }
}

// Usage example
class Program
{
    static void Main(string[] args)
    {
        Console.WriteLine("Bad Design:");
        var badStandardTicket = new BadDesign.StandardTicket();
        badStandardTicket.Reserve();
        badStandardTicket.Print();
        try
        {
            badStandardTicket.SendEmail(); // This will throw an exception
        }
        catch (NotImplementedException e)
        {
            Console.WriteLine($"Error: {e.Message}");
        }

        var badETicket = new BadDesign.ElectronicTicket();
        badETicket.Reserve();
        badETicket.SendEmail();
        try
        {
            badETicket.Print(); // This will throw an exception
        }
        catch (NotImplementedException e)
        {
            Console.WriteLine($"Error: {e.Message}");
        }

        Console.WriteLine("\nGood Design:");
        var goodStandardTicket = new GoodDesign.StandardTicket();
        goodStandardTicket.Reserve();
        goodStandardTicket.Print();
        // goodStandardTicket.SendEmail(); // This is not even an option now

        var goodETicket = new GoodDesign.ElectronicTicket();
        goodETicket.Reserve();
        goodETicket.SendEmail();
        // goodETicket.Print(); // This is not even an option now
    }
}
