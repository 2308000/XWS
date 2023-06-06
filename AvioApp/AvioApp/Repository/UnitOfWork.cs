using Microsoft.EntityFrameworkCore;

namespace AvioApp.Repository
{
    public class UnitOfWork : IUnitOfWork
    {
        private IUserRepository _userRepository;
        private IFlightRepository _flightRepository;
        private ITicketRepository _ticketRepository;
        private DbContext _dbContext;

        public UnitOfWork(DbContext dbContext)
        {
            _dbContext = dbContext;
        }

        public IUserRepository UserRepository
        {
            get
            {
                _userRepository ??= new UserRepository(_dbContext);
                return _userRepository;
            }
        }

        public IFlightRepository FlightRepository
        {
            get
            {
                _flightRepository ??= new FlightRepository(_dbContext);
                return _flightRepository;
            }
        }

        public ITicketRepository TicketRepository
        {
            get
            {
                _ticketRepository ??= new TicketRepository(_dbContext);
                return _ticketRepository;
            }
        }

        public void SaveChanges()
        {
            _dbContext.SaveChanges();
        }
    }
}
