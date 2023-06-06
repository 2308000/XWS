using AvioApp.Model;
using Microsoft.EntityFrameworkCore;

namespace AvioApp.Repository
{
    public class TicketRepository : ITicketRepository
    {
        private readonly DbContext _dbContext;
        public TicketRepository(DbContext dbContext)
        {
            _dbContext = dbContext;
        }

        public Ticket? Create(Ticket entity)
        {
            _dbContext.Set<Ticket>().Add(entity);
            return entity;
        }

        public void Delete(Ticket entity)
        {
            _dbContext.Set<Ticket>().Remove(entity);
        }

        public IQueryable<Ticket> GetAll()
        {
            return _dbContext.Set<Ticket>();
        }

        public Ticket Update(Ticket entity)
        {
            _dbContext.Set<Ticket>().Update(entity);
            return entity;
        }
    }
}
