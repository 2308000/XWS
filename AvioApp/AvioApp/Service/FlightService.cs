using AvioApp.Model;
using AvioApp.Model.DTO;
using AvioApp.Repository;
using AvioApp.SupportClasses.GEH.CustomExceptions;

namespace AvioApp.Service
{
    public class FlightService : IFlightService
    {
        private readonly IFlightRepository _flightRepository;
        private readonly ITicketRepository _ticketRepository;

        public FlightService(IFlightRepository flightRepository, ITicketRepository ticketRepository)
        {
            _flightRepository = flightRepository;
            _ticketRepository = ticketRepository;
        }

        public string Create(NewFlightDTO flight)
        {
            var newFlight = new Flight
            {
                Date = flight.Date,
                Duration = flight.Duration,
                Start = flight.Start,
                Destination = flight.Destination,
                Price = flight.Price,
                Tickets = flight.Tickets,
            };
            newFlight = _flightRepository.Create(newFlight);
            return newFlight.Id;
        }

        public void Delete(string id)
        {
            var flight = _flightRepository.GetAll().FirstOrDefault(f => f.Id == id);
            if (flight == null)
            {
                throw new NotFoundException($"Flight with id {id} does not exist!");
            }
            _flightRepository.Delete(flight.Id);
        }

        public IEnumerable<FlightAdminPreviewDTO> GetAll()
        {
            var flights = _flightRepository.GetAll();
            return flights.Select(
                    f => new FlightAdminPreviewDTO
                    {
                        Id = f.Id,
                        Date = f.Date,
                        Duration = f.Duration,
                        Start = f.Start,
                        Destination = f.Destination,
                        Price = f.Price,
                        Tickets = f.Tickets,
                        RemainingTickets = f.Tickets - _ticketRepository.GetAll().Where(t => t.FlightId == f.Id).Count()
                    }
                );
        }

        public IEnumerable<FlightUserPreviewDTO> Search(FlightSearchDTO query)
        {
            var flights = _flightRepository.GetAll().Where(f => f.Date.Date == query.Date)
                                                               .Where(f => query.Start != "" ? f.Start.Contains(query.Start) : true)
                                                               .Where(f => query.Destination != "" ? f.Destination.Contains(query.Destination) : true)
                                                               .Where(f => query.RequiredTickets != 0 ? (f.Tickets - _ticketRepository.GetAll().Where(t => t.FlightId == f.Id).Count()) >= query.RequiredTickets : true);
            return flights.Select(
                    f => new FlightUserPreviewDTO
                    {
                        Id = f.Id,
                        Date = f.Date,
                        Duration = f.Duration,
                        Start = f.Start,
                        Destination = f.Destination,
                        Price = f.Price,
                        TotalPrice = f.Price * query.RequiredTickets
                    }
                );
        }
    }
}
