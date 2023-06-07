using AvioApp.Model;
using AvioApp.Model.DTO;
using AvioApp.Repository;
using AvioApp.SupportClasses.GEH.CustomExceptions;

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

        public void Buy(string flightId, int amount, string email)
        {
            var flight = _flightRepository.Get(flightId);
            if (flight == null)
            {
                throw new NotFoundException($"Flight with id {flightId} does not exist!");
            }
            var purchasedTickets = _ticketRepository.GetAll().Where(t => t.FlightId == flight.Id).Count();
            if ((flight.Tickets - purchasedTickets - amount) < 0)
            {
                throw new InsufficientResoucesException("There are not enough tickets in stock!");
            }
            for (var i = 0; i < amount; i++)
            {
                _ticketRepository.Create(
                    new Ticket
                    {
                        FlightId = flight.Id,
                        UserId = _userRepository.GetAll().First(u => u.Email == email).Id
                    });
            }
        }

        public IEnumerable<TicketPreviewDTO> GetAll(string email)
        {
            var user = _userRepository.GetAll().Where(u => u.Email == email).First();
            var tickets = _ticketRepository.GetAll().Where(t => t.UserId == user.Id);
            var response = new List<TicketPreviewDTO>();
            foreach (var ticket in tickets)
            {
                var flight = _flightRepository.Get(ticket.FlightId);
                response.Add(new TicketPreviewDTO
                {
                    TicketId = ticket.Id,
                    FlightId = flight.Id,
                    Date = flight.Date,
                    Start = flight.Start,
                    Destination = flight.Destination,
                    Price = flight.Price
                });
            }
            return response;
        }
    }
}
