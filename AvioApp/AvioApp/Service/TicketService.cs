using AvioApp.Model;
using AvioApp.Model.DTO;
using AvioApp.Repository;
using AvioApp.SupportClasses.GEH.CustomExceptions;
using Microsoft.EntityFrameworkCore;

namespace AvioApp.Service
{
    public class TicketService : ITicketService
    {
        private readonly IUnitOfWork _unitOfWork;

        public TicketService(IUnitOfWork unitOfWork)
        {
            _unitOfWork = unitOfWork;
        }

        public void Buy(long flightId, int amount, string email)
        {
            var flight = _unitOfWork.FlightRepository.GetAll().First(f => f.Id == flightId);
            if (flight == null)
            {
                throw new NotFoundException($"Flight with id {flightId} does not exist!");
            }
            var purchasedTickets = _unitOfWork.TicketRepository.GetAll().Include(t => t.Flight).Where(t => t.Flight.Id == flight.Id).Count();
            if ((flight.Tickets - purchasedTickets - amount) < 0)
            {
                throw new InsufficientResoucesException("There are not enough tickets in stock!");
            }
            for (var i = 0; i < amount; i++)
            {
                _unitOfWork.TicketRepository.Create(
                    new Ticket
                    {
                        Flight = flight,
                        User = _unitOfWork.UserRepository.GetAll().First(u => u.Email == email)
                    });
            }
            _unitOfWork.SaveChanges();
        }

        public IEnumerable<TicketPreviewDTO> GetAll(string email)
        {
            var tickets = _unitOfWork.TicketRepository.GetAll().Include(t => t.Flight).Where(t => t.User.Email == email);
            return tickets.Select(t =>
                new TicketPreviewDTO
                {
                    TicketId = t.Id,
                    FlightId = t.Flight.Id,
                    Date = t.Flight.Date,
                    Start = t.Flight.Start,
                    Destination = t.Flight.Destination,
                    Price = t.Flight.Price
                });
        }
    }
}
