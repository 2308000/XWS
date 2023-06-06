using AvioApp.Model;
using AvioApp.Model.DTO;
using AvioApp.Repository;
using AvioApp.SupportClasses.GEH.CustomExceptions;
using Microsoft.EntityFrameworkCore;

namespace AvioApp.Service
{
    public class FlightService : IFlightService
    {
        private readonly IUnitOfWork _unitOfWork;

        public FlightService(IUnitOfWork unitOfWork)
        {
            _unitOfWork = unitOfWork;
        }

        public long Create(NewFlightDTO flight)
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
            newFlight = _unitOfWork.FlightRepository.Create(newFlight);
            _unitOfWork.SaveChanges();
            return newFlight.Id;
        }

        public void Delete(long id)
        {
            var flight = _unitOfWork.FlightRepository.GetAll().FirstOrDefault(f => f.Id == id);
            if (flight == null)
            {
                throw new NotFoundException($"Flight with id {id} does not exist!");
            }
            _unitOfWork.FlightRepository.Delete(flight);
            _unitOfWork.SaveChanges();
        }

        public IEnumerable<FlightAdminPreviewDTO> GetAll()
        {
            var flights = _unitOfWork.FlightRepository.GetAll();
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
                        RemainingTickets = f.Tickets - _unitOfWork.TicketRepository.GetAll().Include(t => t.Flight).Where(t => t.Flight.Id == f.Id).Count()
                    }
                );
        }

        public IEnumerable<FlightUserPreviewDTO> Search(FlightSearchDTO query)
        {
            var flights = _unitOfWork.FlightRepository.GetAll().Where(f => f.Date.Date == query.Date)
                                                               .Where(f => query.Start != "" ? f.Start.Contains(query.Start) : true)
                                                               .Where(f => query.Destination != "" ? f.Destination.Contains(query.Destination) : true)
                                                               .Where(f => query.RequiredTickets != 0 ? (f.Tickets - _unitOfWork.TicketRepository.GetAll().Include(t => t.Flight).Where(t => t.Flight.Id == f.Id).Count()) >= query.RequiredTickets : true);
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
