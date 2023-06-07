using AvioApp.Model.DTO;
using AvioApp.Repository;

namespace AvioApp.Service
{
    public class TicketService : ITicketService
    {
        private readonly IFlightRepository _flightRepository;
        private readonly ITicketRepository _ticketRepository;
        private readonly IUserRepository _userRepository;

        public TicketService(IFlightRepository flightRepository, ITicketRepository ticketRepository, IUserRepository userRepository)
        {
            _flightRepository = flightRepository;
            _ticketRepository = ticketRepository;
            _userRepository = userRepository;
        }

        public void Buy(long flightId, int amount, string email)
        {
            /*var flight = _flightRepository.GetAll().First(f => f.Id == flightId);
            if (flight == null)
            {
                throw new NotFoundException($"Flight with id {flightId} does not exist!");
            }
            var purchasedTickets = _ticketRepository.GetAll().Include(t => t.Flight).Where(t => t.Flight.Id == flight.Id).Count();
            if ((flight.Tickets - purchasedTickets - amount) < 0)
            {
                throw new InsufficientResoucesException("There are not enough tickets in stock!");
            }
            for (var i = 0; i < amount; i++)
            {
                _ticketRepository.Create(
                    new Ticket
                    {
                        Flight = flight,
                        User = _userRepository.GetAll().First(u => u.Email == email)
                    });
            }*/
            throw new NotImplementedException();
        }

        public IEnumerable<TicketPreviewDTO> GetAll(string email)
        {
            throw new NotImplementedException();

            /*var tickets = _ticketRepository.GetAll().Include(t => t.Flight).Where(t => t.User.Email == email);
            return tickets.Select(t =>
                new TicketPreviewDTO
                {
                    TicketId = t.Id,
                    FlightId = t.Flight.Id,
                    Date = t.Flight.Date,
                    Start = t.Flight.Start,
                    Destination = t.Flight.Destination,
                    Price = t.Flight.Price
                });*/
        }
    }
}
