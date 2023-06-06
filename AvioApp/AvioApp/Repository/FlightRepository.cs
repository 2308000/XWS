using AvioApp.Model;
using Microsoft.EntityFrameworkCore;

namespace AvioApp.Repository
{
    public class FlightRepository : IFlightRepository
    {
        private readonly DbContext _dbContext;

        public FlightRepository(DbContext dbContext)
        {
            _dbContext = dbContext;
        }
        public Flight? Create(Flight entity)
        {
            _dbContext.Set<Flight>().Add(entity);
            return entity;
        }

        public void Delete(Flight entity)
        {
            _dbContext.Set<Flight>().Remove(entity);
        }

        public IQueryable<Flight> GetAll()
        {
            return _dbContext.Set<Flight>();
        }

        public Flight Update(Flight entity)
        {
            _dbContext.Set<Flight>().Update(entity);
            return entity;
        }
    }
}
